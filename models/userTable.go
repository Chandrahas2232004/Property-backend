package models

import (
	"time"

	"gorm.io/gorm"
)

/* =========================
   user
========================= */

type User struct {
	UserID         uint      `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	FirstName      string    `gorm:"column:first_name;type:varchar(100);not null" json:"first_name"`
	LastName       string    `gorm:"column:last_name;type:varchar(100);not null" json:"last_name"`
	Email          string    `gorm:"column:email;type:varchar(150);unique;not null" json:"email"`
	HashedPassword string    `gorm:"column:hashed_password;type:varchar(255);not null" json:"hashed_password"`
	PhoneNumber    string    `gorm:"column:phone_number;type:varchar(15)" json:"phone_number"`
	RoleID         uint      `gorm:"column:role_id;not null" json:"role_id"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`

	Role RolesMaster `gorm:"foreignKey:RoleID;references:RolesMasterID" json:"role"`
}

func (User) TableName() string {
	return "user"
}

/* =========================
   AUTO INSERT INTO user_roles
========================= */

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	userRole := UserRoles{
		UserID: u.UserID,
		RoleID: u.RoleID,
	}
	return tx.Create(&userRole).Error
}

/* =========================
   roles_master
========================= */

type RolesMaster struct {
	RolesMasterID uint   `gorm:"column:roles_master_id;primaryKey;autoIncrement" json:"roles_master_id"`
	Role          string `gorm:"column:role;type:varchar(50);unique;not null" json:"role"`
}

func (RolesMaster) TableName() string {
	return "roles_master"
}

/* =========================
   user_roles
========================= */

type UserRoles struct {
	UserRolesID uint `gorm:"column:user_roles_id;primaryKey;autoIncrement" json:"user_roles_id"`
	UserID      uint `gorm:"column:user_id;not null" json:"user_id"`
	RoleID      uint `gorm:"column:role_id;not null" json:"role_id"`
}

func (UserRoles) TableName() string {
	return "user_roles"
}
