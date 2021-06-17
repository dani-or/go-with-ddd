package credits

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"nequi.com/poc-services/internal/services/credits"
)


type createRequest struct {
	DocumentType string `json:"documentType" binding:"required"`
	DocumentNumber string `json:"documentNumber" binding:"required"`
}

func GetCreditsByClientHandler(getCreditsByClientService services.GetCreditsByClientService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println(req.DocumentType, req.DocumentNumber)
		list, err := getCreditsByClientService.GetCreditsByClient(ctx, req.DocumentType, req.DocumentNumber)

		
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		fmt.Println(list);
		n := map[string][]string{"list": list}

		ctx.JSON(http.StatusOK, n)
	}
}