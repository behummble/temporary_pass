package officeservice


type TemproraryPassOffice interface {
	RequestPass()
}

type CookieOffice interface {
	ValidateToken()
}