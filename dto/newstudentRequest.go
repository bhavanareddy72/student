package dto

import (
	"github.com/bhavanareddy72/student/errs"
)

// type NewStudentRequest struct {
// 	studentId   int    `json:"student_id"`
// 	studentType string `json:"student_type"`
// }

type NewStudentRequest struct {
	StudentId        int    `json:"student_id"`
	StudentFirstName string `json:"student_first_name"`
	StudentLastName  string `json:"student_last_name"`
	StudentPhone     string `json:"student_phone"`
	StudentCity      string `json:"student_city"`
}

func (s NewStudentRequest) Validate() *errs.AppError {
	if s.StudentId != 25 {
		return errs.NewValidationError("To open a new Id we need new Id")
	}
	return nil

}
