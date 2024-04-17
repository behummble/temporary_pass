package teorema

import (
	"net/http"

)

type TemproraryPassTeorema struct {

}

type CookieTeorema struct {
	Token string
}

func (pass TemproraryPassTeorema) RequestPass() {

}

func (cookie CookieTeorema) ValidateToken() {
	currentToken := getCurrentToken()

	if cookie.Token != currentToken {
		cookie.refreshToken()
	}
}

func getCurrentToken() string{

}

func (cookie CookieTeorema) refreshToken() {
	
}