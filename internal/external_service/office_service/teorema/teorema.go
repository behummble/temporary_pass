package teorema

import (
	"os"

	"github.com/behummble/temporary_pass/internal/external_service/db/redis"
	"github.com/behummble/temporary_pass/internal/external_service/office_service"
)

type Teorema struct {
	currentToken string
	cookie Cookie
	credentials credentials
	OfficeName string
}

type TemporaryPass struct {

}

type Cookie struct {
	Token string
}

type credentials struct {
	url string
	login string
	password string
}

func (office Teorema) RequestPass(user officeservice.User) {
	
}

func (office Teorema) TokenIsValid() bool {
	
	/*if cookie.Token != currentToken {
		cookie.refreshToken()
	} */
	return true
}

func (office Teorema) RefreshToken() {
	
}

func getActuallToken() string {
	//currentCookie := http.Post("")
	return ""
}

func GetOffice() officeservice.Office {
	return Teorema{
		currentToken: getToken(),
		credentials: setCredentials(),
	}
}

func getToken() string {
	redis := redis.Connect()
	return redis.GetCookie().String()
}

func setCredentials() credentials {
	
	return credentials{
		url: os.Getenv("TEOREMA_WINTER_URL"),
		login: os.Getenv("TEOREMA_WINTER_LOGIN"),
		password: os.Getenv("TEOREMA_WINTER_PASSWORD"),
	}
}