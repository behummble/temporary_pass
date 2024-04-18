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

func (cookie CookieTeorema) TokenIsValid() bool {
	
	/*if cookie.Token != currentToken {
		cookie.refreshToken()
	} */
}

func (cookie CookieTeorema) DefineCurrentToken(){

}

func (cookie CookieTeorema) RefreshToken() {
	
}

func (cookie CookieTeorema) GetCurrentToken() string {
	currentCookie := http.Post("")
}