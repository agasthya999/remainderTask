package service

import (
	"log"
	"remainderTask/database"
	dto "remainderTask/dto"
	"time"
)

func AddTask(request dto.AddTaskReq) error {
	er := database.AddTask(request)
	if er != nil {
		log.Println("AddUser: cannot add user in database", er)
		return er
	}
	return nil
}

func GetTask(request dto.GetTaskReq) (dto.GetTaskRespDto, error) {
	task, er := database.GetTask(request, (int)(time.Now().UTC().UnixMilli()))
	taskRespDto := dto.ToTask(task)
	if er != nil {
		log.Println("AddUser: cannot add user in database", er)
		return *taskRespDto, er
	}
	return *taskRespDto, nil
}
