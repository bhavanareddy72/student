package app

import (
	"encoding/json"
	"net/http"

	"github.com/bhavanareddy72/student/Service"
	"github.com/bhavanareddy72/student/dto"
)

type StudentHandler struct {
	service Service.StudentService
}

func (h StudentHandler) Newstudent(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// studentId := vars["student_id"]
	var request dto.NewStudentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.StudentId = 1 //strconv.Atoi(studentId);
		student, appError := h.service.Newstudent(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, student)
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
