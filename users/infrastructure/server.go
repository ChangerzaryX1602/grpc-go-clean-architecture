package infrastructure

import (
	"fmt"
	"log"
	"net"
	"users/api/pb"
	"users/api/server"
	"users/internal/repository"
	"users/internal/usecase"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Resources struct {
	*gorm.DB
}

func NewServer(version, buildTag, runEnv string) (servers *Resources, err error) {
	mainDbConn, err := ConnectDb(DbConfig{
		DbDriver: "postgres",
		DbName:   viper.GetString("db.postgres.db_name"),
		Host:     viper.GetString("db.postgres.host"),
		Username: viper.GetString("db.postgres.username"),
		Password: viper.GetString("db.postgres.password"),
		Port:     viper.GetInt("db.postgres.port"),
		Timezone: "Asia/Bangkok",
	})
	if err != nil {
		return nil, err
	}
	return &Resources{mainDbConn}, nil
}
func (s *Resources) Run() {
	AutoMigrate(s.DB)
	// Start GRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("grpc.port")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	userRepository := repository.NewUserRepository(s.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userServer := server.NewUserServer(userUsecase)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(LogResponsesInterceptor))
	pb.RegisterUserServiceServer(grpcServer, userServer)
	log.Println("Server started on port :", viper.GetInt("grpc.port"))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
