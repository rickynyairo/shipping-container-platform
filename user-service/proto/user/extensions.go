package shippy_auth

import (
	"github.com/jinzhu/gorm"
	uid "github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
