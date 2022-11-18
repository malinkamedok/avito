package v1

import (
	"avito/internal/entity"
	"avito/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"time"
)

type userRoutes struct {
	t usecase.UserContract
	r usecase.ReportContract
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.UserContract, r usecase.ReportContract) {
	us := userRoutes{t: t, r: r}

	handler.POST("/append", us.append)
	handler.GET("/get-balance/:id", us.getBalance)
	handler.POST("/reserve-money", us.reserveMoney)
	handler.GET("/get-reserve/:id", us.getReserve)
	handler.POST("/accept-income", us.acceptIncome)
	handler.POST("/transfer-money", us.transferMoney)
	handler.POST("/unreserve-money", us.unreserveMoney)
	handler.GET("/get-transactions-by-date/:id/:limit/:offset", us.getTransactionListByDate)
	handler.GET("/get-transactions-by-sum/:id/:limit/:offset", us.getTransactionListBySum)
	handler.GET("/get-all-transactions/:date", us.getAllTransactions)
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

// AppendBalance godoc
// @Summary append balance to user
// @Tags Posts
// @Description create or update user balance
// @Param     request body appendRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/append [post]
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

// GetBalance godoc
// @Summary get user balance
// @Tags Gets
// @Description get user balance
// @Param       id  path 	 string  true "user id"
// @Success     200 {object} uint64
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/get-balance/{id} [get]
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

// GetReserve godoc
// @Summary get user reserve
// @Tags Gets
// @Description get user reserve
// @Param       id  path 	 string  true "user id"
// @Success     200 {object} uint64
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/get-reserve/{id} [get]
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

// ReserveMoney godoc
// @Summary reserve user money
// @Tags Posts
// @Description create user reserve
// @Param     request body reserveRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/reserve-money [post]
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

// AcceptIncome godoc
// @Summary accept income
// @Tags Posts
// @Description the service was provided. accept income
// @Param     request body acceptRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/accept-income [post]
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

// TransferMoney godoc
// @Summary transfer money
// @Tags Posts
// @Description user to user money transfer
// @Param     request body transferRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/transfer-money [post]
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

// UnreserveMoney godoc
// @Summary unreserve user money
// @Tags Posts
// @Description the service was not provided. return money to user balance
// @Param     request body unreserveRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/unreserve-money [post]
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

// GetTransactionsListByDate godoc
// @Summary get transactions list by date
// @Tags Gets
// @Description get transactions list by date
// @Param       id  path 	 string  true "user id"
// @Param		limit path uint64 true "amount of rows"
// @Param		offset path uint64 true "amount of skipped rows"
// @Success     200 {object} transactionListResponse
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/get-transactions-by-date/{id}/{limit}/{offset} [get]
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

// GetTransactionsListBySum godoc
// @Summary get transactions list by sum
// @Tags Gets
// @Description get transactions list by sum
// @Param       id  path 	 string  true "user id"
// @Param		limit path uint64 true "amount of rows"
// @Param		offset path uint64 true "amount of skipped rows"
// @Success     200 {object} transactionListResponse
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/get-transactions-by-sum/{id}/{limit}/{offset} [get]
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

// GetALLTransactions godoc
// @Summary get all transactions list in one month
// @Tags Gets
// @Description get all transactions list in one month
// @Param       date  path 	 string  true "YYYY-Mmm (example: 2022-Nov)"
// @Success     200 {object} transactionListResponse
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/get-all-transactions/{date} [get]
func (u *userRoutes) getAllTransactions(c *gin.Context) {
	date, err := time.Parse("2006-Jan", c.Param("date"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "error in parsing date")
		return
	}
	res, err := u.r.GenerateReportByPeriod(c.Request.Context(), date)
	if err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "cannot generate .csv report")
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=report.csv")
	c.Data(http.StatusOK, "text/csv", res.Bytes())
}
