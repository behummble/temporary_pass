package officeservice


type TemproraryPassOffice interface {
	RequestPass()
}

type CookieOffice interface {
	DefineCurrentToken()
	TokenIsValid() bool
	RefreshToken()
	GetCurrentToken() string
}