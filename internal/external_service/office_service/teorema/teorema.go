package teorema

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"io"
	"github.com/behummble/temporary_pass/internal/external_service/db/redis"
	"github.com/behummble/temporary_pass/internal/external_service/office_service"
)

type Teorema struct {
	currentToken string
	credentials credentials
	OfficeName string
}

type TemporaryPass struct {
	IssueType string `json:"issue_type_id"`
	LocationID string `json:"location_id"`
	Attributes struct{
		Value string `json:"3pavd71wx3nb"`
	}				`json:"attributes"`
}

type credentials struct {
	urlAuth string
	urlPass string
	login string
	password string

}

type authorization struct {
	PhoneNumber string
	Password string
}

func (office *Teorema) RequestPass(user officeservice.User) {
	jsonString, err := json.Marshal(newTemporaryPass(user))
	if err != nil {
		log.Println(err)
	} else {
		resp, err := http.Post(office.credentials.urlPass, "application/json", bytes.NewBuffer(jsonString))
		if err != nil {
			log.Println(err)
		}
		if resp.StatusCode != 201 {
			answer, _ := io.ReadAll(resp.Body)
			log.Println(string(answer))
		}
		resp.Body.Close()
	}
}

func (office *Teorema) TokenIsValid() bool {
	if office.currentToken == "" {
		return false
	}

	actuallToken := getActuallToken(
		office.credentials.urlAuth, 
		office.credentials.login,
		office.credentials.password)

	if actuallToken != "" && office.currentToken != actuallToken {
		office.currentToken = actuallToken
		return false
	}

	return true
}

func (office *Teorema) RefreshToken() {
	redisClient := redis.NewClient()
	redisClient.SetCookie(office.currentToken)
}

func getActuallToken(host, login, password string) string {
	jsonAuth, err := json.Marshal(newAuthorizationRequest(login, password))
	if err != nil {
		log.Println(err)
		return ""
	}
	resp, err := http.Post(host, "application/json", bytes.NewBuffer(jsonAuth))
	if err != nil {
		log.Println(err)
		return ""
	}
	if resp.StatusCode != 200 {
		answer, _ := io.ReadAll(resp.Body)
		log.Println(string(answer))
		return ""
	}
	token := resp.Header.Get("Authorization")
	
	return token
}

func GetOffice() officeservice.Office {
	return &Teorema{
		currentToken: getToken(),
		credentials: setCredentials(),
	}
}

func getToken() string {
	redisClient := redis.NewClient()
	cookie, err := redisClient.GetCookie()
	
	if err != nil {
		log.Println(err)
		return ""
	}
	return cookie.String()
}

func setCredentials() credentials {
	return credentials{
		urlAuth: os.Getenv("TEOREMA_WINTER_AUTH_URL"),
		urlPass: os.Getenv("TEOREMA_WINTER_TMP_PASS_URL"),
		login: os.Getenv("TEOREMA_WINTER_LOGIN"),
		password: os.Getenv("TEOREMA_WINTER_PASSWORD"),
	}
}

func newTemporaryPass(user officeservice.User) TemporaryPass {
	return TemporaryPass{
		IssueType: "ns2a0",
		LocationID: "13",
		Attributes: struct{
			Value string `json:"3pavd71wx3nb"`}{
			Value: user.GetPhoneNumber()},
	}
}

func newAuthorizationRequest(phoneNumber, password string) authorization {
	return authorization{
		PhoneNumber: phoneNumber,
		Password: password,
	}
}