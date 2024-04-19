package officeservice


type Office interface {
	TokenIsValid() bool
	RefreshToken()
	RequestPass(user User)
}

type User interface {
	
}