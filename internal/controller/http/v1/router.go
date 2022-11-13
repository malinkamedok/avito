package v1

import (
	"avito/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine,
	u usecase.UserContract) {

	h := handler.Group("/v1")

	{
		newUserRoutes(h, u)
	}
}
