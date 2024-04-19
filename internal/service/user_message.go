package service

import (
	"github.com/behummble/temporary_pass/internal/external_service/db"
)

type User struct{
	fullName string
	phoneNumber string
}

func ListenUserMessagesFromDB(dbase db.DB) {
	
	messages := make(chan db.UserMessage)
	go dbase.GetMessages(messages)

	for {
		select {
		case message := <-messages :
			message.GetOffice().RequestPass(getUser(message))
		}
	}
}

func getUser(msg db.UserMessage) User {
	return User{
		fullName: msg.GetUserName(),
		phoneNumber: msg.GetPhoneNumber(),
	}
}

func (user User) GetUserName() string {
	return user.fullName
}

func (user User) GetPhoneNumber() string {
	return user.phoneNumber
} 