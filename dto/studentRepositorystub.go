package dto

type studentRepositorystub struct {
	students []Student
}

func (s studentRepositorystub) FindAll() ([]Student, error) {
	return s.students, nil
}
func NewstudentRepositorystub() studentRepositorystub {
	students := []Student{
		{StudentId: 1, StudentfirstName: "gangidi", StudentlastName: "praneethReddy", Studentphone: "984567831", Studentcity: "Boston"},
		{StudentId: 2, StudentfirstName: "Adla", StudentlastName: "BhavanaReddy", Studentphone: "9876543210", Studentcity: "MA Boston"},
	}
	return studentRepositorystub{students: students}
}
