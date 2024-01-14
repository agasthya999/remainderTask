package dto

type AddTaskReq struct {
	UserID      int
	ContactId   int
	Title       string
	Description string
	Priority    int
	DueDate     int //epoch time
}

type GetTaskReq struct {
	UserID int
}
