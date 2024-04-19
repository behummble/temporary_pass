package db

import(
	"github.com/behummble/temporary_pass/internal/external_service/office_service"
)


type DB interface {
	GetCookie() Cookie
	GetMessages(chan<- UserMessage)
}

type UserMessage interface {
	JsonString() string
	GetOffice() officeservice.Office
	GetUserName() string
	GetPhoneNumber() string
}

type Cookie interface {
	String() string
}