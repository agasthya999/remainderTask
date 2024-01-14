package database

import (
	"fmt"
	"log"
	dto "remainderTask/dto"
	"remainderTask/models"
)

func AddUser(user dto.AddUserReq) error {
	resp := DBWrite.Create(&user)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func GetUserDetailsById(userId string) (models.User, error) {
	var userDet models.User
	resp := DB.Table("users").Select("*").Where("user_id =?", userId).Scan(&userDet)
	if resp.Error != nil || userDet.UserId == "" {
		log.Println("GetUserDetailsById: user doesnot exist")
		return userDet, err
	}
	return userDet, nil
}

func GetAllUsers() (users []models.User, err error) {
	result := DBRead.Table("users").Select("*").Scan(&users)
	if result.Error != nil {
		log.Println("GetAllUsers: unable to get users from database, err - ", result.Error)
		return users, result.Error
	}
	return users, nil
}

func CreateUserContacts(userContacts models.UserContacts) error {
	resp := DBWrite.Create(&userContacts)
	if resp.Error != nil {
		err := fmt.Errorf("unable to add contacts")
		return err
	}
	return nil
}
