package db


type DB interface {
	GetCookie() (Cookie, error)
	GetMessages(chan<- UserMessage, string)
	SetCookie(string)
}

type UserMessage interface {
	JsonString() string
	GetUserName() string
	GetPhoneNumber() string
	GetOfficeName() string
}

type Cookie interface {
	String() string
}