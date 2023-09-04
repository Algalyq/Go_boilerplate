package handler
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Algalyq/Go_boilerplate/data")

func (h *Handler) signup(ctx *gin.Context) {
	var input data.User

	if err := ctx.BindJSON(&input); err!= nil {
			newErrorResponse(ctx,http.StatusBadRequest, err.Error())
			return
	}
	
	if err := h.validate(ctx,input); err == 1{
	msg, err := h.services.Authorization.CreateUser(input)
	if err!= nil {
		newErrorResponse(ctx,http.StatusBadRequest, err.Error())
		return 
	}
	ctx.JSON(http.StatusOK,map[string]interface{}{
		"msg":msg,
		"status":http.StatusOK,
	})
}	
}