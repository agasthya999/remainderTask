package dto

import "remainderTask/models"

type GetTaskRespDto struct {
	UserID           int
	ContactId        int
	Title            string
	Description      string
	Priority         int
	DueDate          int    //epoch time
	Status           string `json:"status"`
	ErrorDescription string `json:"errorDescription"`
}

func ToTask(task models.Task) *GetTaskRespDto {
	return &GetTaskRespDto{
		UserID:      task.UserID,
		ContactId:   task.ContactId,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		DueDate:     task.DueDate, //epoch time
	}
}
