package credits

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"nequi.com/poc-services/internal/services/credits"
)


type createRequest struct {
	Debenture string `json:"debenture" binding:"required"`
	CustomerId string `json:"customerId" binding:"required"`
}

func GetCreditHandler(getCreditService services.GetCreditService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println(req.CustomerId, req.Debenture)
		credit, err := getCreditService.GetCredit(req.CustomerId, req.Debenture)

		if err != nil {
			errormio := map[string]string{"errorDesc": err.Error()}
			ctx.JSON(http.StatusInternalServerError, errormio)
			return
		}		
		ctx.JSON(http.StatusOK, credit)
	}
}