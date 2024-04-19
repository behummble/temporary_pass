package service

import (
	"time"

	"github.com/behummble/temporary_pass/internal/external_service/office_service"
)

func StartTokenValidationProccess(offices []officeservice.Office) {
	for _, office := range offices {
		if !office.TokenIsValid() {
			office.RefreshToken()
		}
	}

	time.Sleep(time.Hour)
	//StartTokenValidationProccess(cookieOffice)
}