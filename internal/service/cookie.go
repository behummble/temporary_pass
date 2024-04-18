package service

import(
	"github.com/behummble/temporary_pass/internal/external_service/office_service"
)

func StartTokenValidationProccess(cookieOffice []officeservice.CookieOffice) {
	for _, cookie := range cookieOffice {
		cookie.DefineCurrentToken()
		if !cookie.TokenIsValid() {
			cookie.RefreshToken()
		}
	}
}