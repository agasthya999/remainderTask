package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId   string `gorm:"primary_key"`
	UserName string
	Password string
	Contacts []User `gorm:"many2many:userContacts;"`
}

type UserContacts struct {
	UserId    int
	ContactId int
}

func ToUserContacts(userId int, contactId int) *UserContacts {
	return &UserContacts{
		UserId:    userId,
		ContactId: contactId,
	}
}
