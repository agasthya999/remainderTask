package handler

import (
	"log"

	constants "remainderTask/constants"
	dto "remainderTask/dto"
	svc "remainderTask/service"

	"github.com/gofiber/fiber/v2"
)

func AddUser(c *fiber.Ctx) error {
	c.Accepts("json", "text")
	c.Accepts("application/json")

	// 1. Create default response dto
	respDto := dto.DefaultRespDto{}
	respDto.Status = constants.ERROR_STATUS

	// 2. Parse request body
	reqStr := string(c.Body())
	log.Println("CreateUser: Request Body is - ", reqStr)
	reqDto := dto.AddUserReq{}
	err := c.BodyParser(&reqDto)
	if err != nil {
		respDto.ErrorDescription = constants.INVALID_REQ_ERROR_DESC
		log.Println("CreateUser: Body parsing failed with error - ", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(respDto)
	}

	// 3. Event creation
	er := svc.AddUser(reqDto)
	if er != nil {
		respDto.ErrorDescription = er.Error()
		return c.Status(fiber.StatusOK).JSON(respDto)
	}

	respDto.Status = constants.OK_STATUS
	respDto.ErrorDescription = ""
	return c.Status(fiber.StatusOK).JSON(respDto)
}

func GetAllUsers(c *fiber.Ctx) error {
	c.Accepts("json", "text")
	c.Accepts("application/json")

	// 1. Create response dto
	respDto := dto.UserDetailsRespDto{}
	respDto.ErrorDescription = constants.GENERAL_ERROR_DESC
	respDto.Status = constants.ERROR_STATUS

	// 4. Get all users
	users, err := svc.GetAllUsers()
	if err != nil {
		log.Println("GetAllUsers: unable to get users")
		respDto.ErrorDescription = constants.SOMETHING_WENT_WRONG
		return c.Status(fiber.StatusInternalServerError).JSON(respDto)
	}

	// 5. Send success response
	respDto.Users = users
	respDto.Status = constants.OK_STATUS
	respDto.ErrorDescription = ""
	return c.Status(fiber.StatusOK).JSON(respDto)
}

func AddContactToUser(c *fiber.Ctx) error {
	c.Accepts("json", "text")
	c.Accepts("application/json")

	// 1. Create default response dto
	respDto := dto.DefaultRespDto{}
	respDto.Status = constants.ERROR_STATUS
	respDto.ErrorDescription = constants.GENERAL_ERROR_DESC

	// 3. Parse request body
	reqStr := string(c.Body())
	log.Println("CreateUser: Request Body is - ", reqStr)
	reqDto := dto.AddUserContactsReqDto{}
	err := c.BodyParser(&reqDto)

	if err != nil {
		respDto.ErrorDescription = constants.INVALID_REQ_ERROR_DESC
		log.Println("CreateUser: Body parsing failed with error - ", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(respDto)
	}

	// 4. Event creation
	er := svc.MasterPlayers(reqDto)
	if er != nil {
		respDto.ErrorDescription = er.Error()
		log.Println("CreateMasterPlayers: Failed to create Players - ", er.Error())
		respDto.ErrorDescription = "Failed to add contact to user "
		return c.Status(fiber.StatusOK).JSON(respDto)
	}
	respDto.Status = constants.OK_STATUS
	return c.Status(fiber.StatusOK).JSON(respDto)
}
