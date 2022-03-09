package dto

import (
	"github.com/bhavanareddy72/student/errs"
)

type Student struct {
	StudentfirstName string
	StudentlastName  string
	StudentId        int64
	Studentphone     string
	Studentcity      string
}

func (s Student) ToNewstudentResponseDto() NewstudentResponse {
	return NewstudentResponse{s.StudentId}

}

type StudentRepository interface {
	Save(Student) (*Student, *errs.AppError)
}
