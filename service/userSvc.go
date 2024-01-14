package service

import (
	"log"
	"remainderTask/database"
	dto "remainderTask/dto"
	"remainderTask/models"
)

func AddUser(request dto.AddUserReq) error {
	er := database.AddUser(request)
	if er != nil {
		log.Println("AddUser: cannot add user in database", er)
		return er
	}
	return nil

}

func GetAllUsers() ([]models.User, error) {
	users, err := database.GetAllUsers()
	if err != nil {
		log.Println("GetAllUsers: unable to get users, err -", err)
		return users, err
	}
	return users, nil
}

func MasterPlayers(userContacts dto.AddUserContactsReqDto) error {
	contact := models.ToUserContacts(userContacts.UserId, userContacts.ContactId)
	error := database.CreateUserContacts(*contact)
	if error != nil {
		log.Println("MasterPlayers: unable to create master players", error)
		return error
	}
	return nil
}
