package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bhavanareddy72/student/Service"
	"github.com/bhavanareddy72/student/dto"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environemnt variables are not defined")
	}

}

func Start() {
	sanityCheck()
	router := mux.NewRouter()

	dbclient := getDbclient()
	studentRepositoryDb := dto.NewstudentRepositoryDb(dbclient)

	ah := StudentHandler{Service.NewstudentService(studentRepositoryDb)}

	// router.HandleFunc("/students", ch).Methods(http.MethodGet)
	// router.HandleFunc("/students/{student_id:[0-9]+}", ch.getstudent).Methods(http.MethodGet)
	router.HandleFunc("/students/{student_id:[0-9]+}/student", ah.Newstudent).Methods(http.MethodPost)
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
func getDbclient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbpasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbport := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbpasswd, dbAddr, dbport, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
