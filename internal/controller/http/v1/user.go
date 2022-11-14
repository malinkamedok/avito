package v1

import (
	"avito/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type userRoutes struct {
	t usecase.UserContract
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.UserContract) {
	us := userRoutes{t: t}

	handler.POST("/append", us.append)
	handler.GET("/get-balance/:id", us.getBalance)
	handler.POST("/reserve-money", us.reserveMoney)
	handler.GET("/get-reserve/:id", us.getReserve)
	handler.POST("/accept-income", us.acceptIncome)
}

type appendRequest struct {
	User uuid.UUID `json:"user"`
	Sum  uint64    `json:"sum"`
}

type reserveRequest struct {
	UserUUID    uuid.UUID `json:"userUUID"`
	ServiceUUID uuid.UUID `json:"serviceUUID"`
	OrderUUID   uuid.UUID `json:"orderUUID"`
	Amount      uint64    `json:"amount"`
}

type acceptRequest struct {
	UserUUID    uuid.UUID `json:"userUUID"`
	ServiceUUID uuid.UUID `json:"serviceUUID"`
	OrderUUID   uuid.UUID `json:"orderUUID"`
	Amount      uint64    `json:"amount"`
}

func (u *userRoutes) append(c *gin.Context) {
	var req appendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "Error in request credentials")
		return
	}
	err := u.t.AppendBalance(c.Request.Context(), req.User, req.Sum)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in updating user balance")
		return
	}
	c.JSONP(http.StatusOK, nil)
}

func (u *userRoutes) getBalance(c *gin.Context) {
	userUUID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing user uuid")
		return
	}
	balance, err := u.t.GetBalance(c.Request.Context(), userUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in getting user balance")
		return
	}
	c.JSONP(http.StatusOK, balance)
}

type reserveResponse struct {
	Reserves []int64 `json:"reserveList"`
}

func (u *userRoutes) getReserve(c *gin.Context) {
	userUUID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing user uuid")
		return
	}
	reserve, err := u.t.GetReserve(c.Request.Context(), userUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in getting user balance")
		return
	}
	c.JSONP(http.StatusOK, reserveResponse{reserve})
}

func (u *userRoutes) reserveMoney(c *gin.Context) {
	var request reserveRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, "Error in request credentials")
		return
	}
	err := u.t.ReserveMoney(c.Request.Context(), request.UserUUID, request.ServiceUUID, request.OrderUUID, request.Amount)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in reserving user money")
		return
	}
	c.JSONP(http.StatusOK, nil)
}

func (u *userRoutes) acceptIncome(c *gin.Context) {
	var request acceptRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, "Error in request credentials")
		return
	}
	err := u.t.AcceptIncome(c.Request.Context(), request.UserUUID, request.ServiceUUID, request.OrderUUID, request.Amount)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in accepting income")
		return
	}
	c.JSONP(http.StatusOK, nil)
}
