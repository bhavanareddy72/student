package dto

import (
	"github.com/bhavanareddy72/student/errs"
	"github.com/bhavanareddy72/student/logger"
	"github.com/jmoiron/sqlx"
)

type studentsRepositoryDb struct {
	client *sqlx.DB
}

func (d studentsRepositoryDb) Save(s Student) (*Student, *errs.AppError) {
	sqlInsert := "INSERT INTO students(student_id, studentlastName,studentFirstName,studentphone,studentcity) values(?,?,?,?)"
	result, err := d.client.Exec(sqlInsert, s.StudentId, s.StudentlastName, s.StudentfirstName, s.Studentphone, s.Studentcity)
	if err != nil {
		logger.Error("Error while creating new account:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new student:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	s.StudentId = id
	return &s, nil
}
func NewstudentRepositoryDb(dbclient *sqlx.DB) studentsRepositoryDb {
	return studentsRepositoryDb{dbclient}
}
