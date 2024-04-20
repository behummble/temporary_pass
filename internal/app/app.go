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
	go startTokenValidationProccess()
	startListenQueue()
}


func startTokenValidationProccess() {
	service.StartTokenValidationProccess(getOffices())
}

func initLog() {
//	log = log.New(os.Stdout, "TEMP_PASS", 1)
}

func startListenQueue() {
	redis := redis.NewClient()
	queues := getMessagesQueues()
	service.ListenUserMessagesFromDB(redis, queues)
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

func getMessagesQueues() []string {
	queues := make([]string,0)
	queues = append(queues, "WINTER_OFFICE")
	return queues
}