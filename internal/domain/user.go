package domain

import (
	"time"

	"github.com/fatih/structs"
)

type User struct {
	ID        string    `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Age       int       `json:"age" gorm:"column:age"`
	IDCompany string    `json:"id_company" gorm:"column:id_company"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (*User) TableName() string {
	return "user"
}

func (u *User) ToMap() map[string]interface{} {
	return structs.Map(u)
}
