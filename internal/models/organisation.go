package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organisation struct{
	OrgId string `json:"orgId" gorm:"primaryKey;type:varchar(255)"`
	Name    string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description;required;not null" json:"description"`

}

func (u *Organisation) BeforeCreate(tx *gorm.DB) error {
	u.OrgId = uuid.NewString()

	return nil
}