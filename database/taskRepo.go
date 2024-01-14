package database

import (
	dto "remainderTask/dto"
	"remainderTask/models"
)

func AddTask(task dto.AddTaskReq) error {
	resp := DBWrite.Create(&task)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func GetTask(user dto.GetTaskReq, time int) (models.Task, error) {
	var task models.Task
	resp := DBRead.Table("task").
		Where("user_id = ? AND due_date > time ", user.UserID, time).
		Find(&task)
	if resp.Error != nil {
		return task, resp.Error
	}
	return task, nil
}
