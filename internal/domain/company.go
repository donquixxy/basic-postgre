package domain

import (
	"encoding/json"
	"time"
)

type Company struct {
	ID        string    `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Phone     string    `json:"phone" gorm:"column:phone"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (*Company) TableName() string {
	return "company"
}

func (c *Company) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(c)

	return data, err
}

func (c *Company) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
