package entity

type User struct {
	ID    uint32 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
