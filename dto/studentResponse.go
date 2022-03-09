package dto

type StudentResponse struct {
	Id        int64  `json:"student_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
	Phone     string `json:"zipcode"`
}
