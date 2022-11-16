package v1

import (
	"avito/internal/entity"
	"avito/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
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
	handler.POST("/transfer-money", us.transferMoney)
	handler.POST("/unreserve-money", us.unreserveMoney)
	handler.GET("/get-transactions-by-date/:id/:limit/:offset", us.getTransactionListByDate)
	handler.GET("/get-transactions-by-sum/:id/:limit/:offset", us.getTransactionListBySum)
	handler.GET("/get-all-transactions/:service-id/:date", us.getAllTransactions)
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
	ServiceName string    `json:"serviceName"`
	OrderUUID   uuid.UUID `json:"orderUUID"`
	Amount      uint64    `json:"amount"`
}

type transferRequest struct {
	FirstUserUUID  uuid.UUID `json:"firstUserUUID"`
	SecondUserUUID uuid.UUID `json:"secondUserUUID"`
	Amount         uint64    `json:"amount"`
}

type unreserveRequest struct {
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
		errorResponse(c, http.StatusInternalServerError, "error in getting user reserve")
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
	err := u.t.AcceptIncome(c.Request.Context(), request.UserUUID, request.ServiceUUID, request.ServiceName, request.OrderUUID, request.Amount)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in accepting income")
		return
	}
	c.JSONP(http.StatusOK, nil)
}

func (u *userRoutes) transferMoney(c *gin.Context) {
	var request transferRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, "Error in request credentials")
		return
	}
	err := u.t.UserToUserMoneyTransfer(c.Request.Context(), request.FirstUserUUID, request.SecondUserUUID, request.Amount)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in transfering money")
		return
	}
	c.JSONP(http.StatusOK, nil)
}

func (u *userRoutes) unreserveMoney(c *gin.Context) {
	var request unreserveRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, "Error in request credentials")
		return
	}
	err := u.t.UnreserveMoney(c.Request.Context(), request.UserUUID, request.ServiceUUID, request.OrderUUID, request.Amount)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in unreserving money")
		return
	}
	c.JSONP(http.StatusOK, nil)
}

type transactionListResponse struct {
	List []entity.Transaction `json:"transactions"`
}

func (u *userRoutes) getTransactionListByDate(c *gin.Context) {
	userUUID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing user uuid")
		return
	}
	limit, err := strconv.ParseUint(c.Param("limit"), 10, 64)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing limit")
		return
	}
	offset, err := strconv.ParseUint(c.Param("offset"), 10, 64)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing offset")
		return
	}
	transactions, err := u.t.GetTransactionListByDate(c.Request.Context(), userUUID, limit, offset)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in getting transaction list by date")
		return
	}
	c.JSONP(http.StatusOK, transactionListResponse{List: transactions})
}

func (u *userRoutes) getTransactionListBySum(c *gin.Context) {
	userUUID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing user uuid")
		return
	}
	limit, err := strconv.ParseUint(c.Param("limit"), 10, 64)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing limit")
		return
	}
	offset, err := strconv.ParseUint(c.Param("offset"), 10, 64)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing offset")
		return
	}
	transactions, err := u.t.GetTransactionListBySum(c.Request.Context(), userUUID, limit, offset)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in getting transaction list by sum")
		return
	}
	c.JSONP(http.StatusOK, transactionListResponse{List: transactions})
}

type allTransactionsListResponse struct {
	List []entity.Report `json:"reports"`
}

func (u *userRoutes) getAllTransactions(c *gin.Context) {
	serviceUUID, err := uuid.Parse(c.Param("service-id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing service uuid")
		return
	}
	fmt.Println(c.Param("date"))
	date, err := time.Parse("2006-Jan", c.Param("date"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing date")
		return
	}
	reports, err := u.t.GetAllTransactions(c.Request.Context(), serviceUUID, date)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in getting transactions list")
		return
	}
	c.JSONP(http.StatusOK, allTransactionsListResponse{List: reports})
}
