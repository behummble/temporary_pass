package service

import (
	"fmt"
	"log"

	"github.com/behummble/temporary_pass/internal/external_service/db"
	"github.com/behummble/temporary_pass/internal/external_service/office_service"
	"github.com/behummble/temporary_pass/internal/external_service/office_service/teorema"
)

type User struct{
	fullName string
	phoneNumber string
}

func ListenUserMessagesFromDB(dbase db.DB, queues []string) {
	
	messages := make(chan db.UserMessage)

	for _, queue := range queues {
		go dbase.GetMessages(messages, queue)
	}

	for {
		select {
		case message := <-messages :
			office, err := defineOffice(message.GetOfficeName())
			if err != nil {
				log.Println(err)
				continue
			}
			office.RequestPass(getUser(message))
		}
	}
}

func getUser(msg db.UserMessage) User {
	return User{
		fullName: msg.GetUserName(),
		phoneNumber: msg.GetPhoneNumber(),
	}
}

func (user User) GetFullName() string {
	return user.fullName
}

func (user User) GetPhoneNumber() string {
	return user.phoneNumber
}

func defineOffice(name string) (officeservice.Office, error) {
	switch name {
	case "pool:teorema_request" :
		return teorema.GetOffice(), nil
	}
	return nil, fmt.Errorf("CantDefineOfficeByName:%s", name)
}
