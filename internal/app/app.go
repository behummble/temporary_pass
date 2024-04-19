package app

import (
	"log"
	"github.com/behummble/temporary_pass/internal/service"
	"github.com/behummble/temporary_pass/internal/external_service/office_service/teorema"
	"github.com/behummble/temporary_pass/internal/external_service/office_service"
	"github.com/behummble/temporary_pass/internal/external_service/db/redis"
	"github.com/joho/godotenv"
)


func Run() {
	initLog()
	initEnv()
	startListenQueue()
	go startTokenValidationProccess()
}


func startTokenValidationProccess() {
	service.StartTokenValidationProccess(getOffices())
}

func initLog() {
//	log = log.New(os.Stdout, "TEMP_PASS", 1)
}

func startListenQueue() {
	redis := redis.Connect()
	service.ListenUserMessagesFromDB(redis)
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func getOffices() []officeservice.Office {
	offices := make([]officeservice.Office, 0)
	offices = append(offices, teorema.GetOffice())
	return offices
}