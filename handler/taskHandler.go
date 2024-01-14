package handler

import (
	"log"
	constants "remainderTask/constants"
	dto "remainderTask/dto"
	svc "remainderTask/service"

	"github.com/gofiber/fiber/v2"
)

func AddTask(c *fiber.Ctx) error {
	c.Accepts("json", "text")
	c.Accepts("application/json")

	// 1. Create default response dto
	respDto := dto.DefaultRespDto{}
	respDto.Status = constants.ERROR_STATUS

	// 2. Parse request body
	reqStr := string(c.Body())
	log.Println("AddTask: Request Body is - ", reqStr)
	reqDto := dto.AddTaskReq{}
	err := c.BodyParser(&reqDto)
	if err != nil {
		respDto.ErrorDescription = constants.INVALID_REQ_ERROR_DESC
		log.Println("CreateUser: Body parsing failed with error - ", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(respDto)
	}

	// 3. Event creation
	er := svc.AddTask(reqDto)
	if er != nil {
		respDto.ErrorDescription = er.Error()
		return c.Status(fiber.StatusOK).JSON(respDto)
	}

	respDto.Status = constants.OK_STATUS
	respDto.ErrorDescription = ""
	return c.Status(fiber.StatusOK).JSON(respDto)
}

func GetNextTaskForUser(c *fiber.Ctx) error {
	c.Accepts("json", "text")
	c.Accepts("application/json")

	// 1. Create default response dto
	respDto := dto.GetTaskRespDto{}
	respDto.Status = constants.ERROR_STATUS

	// 2. Parse request body
	reqStr := string(c.Body())
	log.Println("GetTask: Request Body is - ", reqStr)
	reqDto := dto.GetTaskReq{}
	err := c.BodyParser(&reqDto)
	if err != nil {
		respDto.ErrorDescription = constants.INVALID_REQ_ERROR_DESC
		log.Println("GetTask: Body parsing failed with error - ", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(respDto)
	}

	// 3. Event creation
	task, er := svc.GetTask(reqDto)
	if er != nil {
		respDto.ErrorDescription = er.Error()
		return c.Status(fiber.StatusOK).JSON(respDto)
	}

	respDto.Status = constants.OK_STATUS
	respDto.ErrorDescription = ""
	return c.Status(fiber.StatusOK).JSON(task)
}
