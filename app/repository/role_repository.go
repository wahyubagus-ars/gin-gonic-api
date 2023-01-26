package repository

import (
	"gin-gonic-api/app/domain/dao"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAllRole()
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func (r RoleRepositoryImpl) FindAllRole() {
	panic("implement me")
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepositoryImpl {
	db.AutoMigrate(&dao.Role{}, &dao.User{})
	return &RoleRepositoryImpl{
		db: db,
	}
}
