package Service

import (
	"github.com/bhavanareddy72/student/dto"
	"github.com/bhavanareddy72/student/errs"
)

type StudentService interface {
	Newstudent(dto.NewStudentRequest) (*dto.NewstudentResponse, *errs.AppError)
}
type DefaultstudentService struct {
	repo dto.StudentRepository
}

func (s DefaultstudentService) Newstudent(req dto.NewStudentRequest) (*dto.NewstudentResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := dto.Student{
		StudentId:        int64(req.StudentId),
		StudentfirstName: req.StudentFirstName,
		StudentlastName:  req.StudentLastName,
		Studentphone:     req.StudentPhone,
		Studentcity:      req.StudentCity,
	}
	Newstudent, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := Newstudent.ToNewstudentResponseDto()
	return &response, nil
}
func NewstudentService(repo dto.StudentRepository) DefaultstudentService {
	return DefaultstudentService{repo}
}
