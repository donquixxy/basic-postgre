package domain

import (
	"encoding/json"
	"fmt"
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
	Company   Company   `json:"company" gorm:"foreignKey:ID;references:IDCompany"`
}

func (*User) TableName() string {
	return "user"
}

func (u *User) ToMap() map[string]interface{} {
	return structs.Map(u)
}

func (u *User) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err.Error())
	}
	return data, err
}

func (u *User) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, u)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON:", err.Error())
	}
	return err
}
