package db


type DB interface {
	Connect()
	GetCookie()
}