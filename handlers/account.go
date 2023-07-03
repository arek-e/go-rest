// Package handlers containing CRUD for Accounts
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/arek-e/lanexpense/domain"
	"github.com/arek-e/lanexpense/services"
)

type AccountHandlers struct {
	Platform *services.AccountService
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// GetAllAccounts
func (h *AccountHandlers) GetAllAccounts(ctx *gin.Context) {
	var response Response

	books, err := h.Platform.GetAllAccounts()
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = books
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

// GetAccount
func (h *AccountHandlers) GetAccount(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	book, err := h.Platform.GetAccount(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = book
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

// CreateAccount
func (h *AccountHandlers) CreateAccount(ctx *gin.Context) {
	var response Response
	var request domain.Account
	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	account, code, err := h.Platform.CreateAccount(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		ctx.JSON(code, response)
		return
	}

	response.Data = account
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

// UpdateAccount
func (h *AccountHandlers) UpdateAccount(ctx *gin.Context) {
	var response Response
	var request domain.Account
	ctx.Header("Content-Type", "application/json")
	id := ctx.Param("id")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	request.ID = id

	account, err := h.Platform.UpdateAccount(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = account
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

// DeleteAccount
func (h *AccountHandlers) DeleteAccount(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	if err := h.Platform.DeleteAccount(id); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = nil
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

// // AccountsCreate function for creating an new Account
// func AccountsCreate(c *gin.Context) {
// 	// Get data of request body
// 	var AccountCreateRequest struct {
// 		Name             string          `json:"name"`
// 		Currency         domain.Currency `json:"currency"`
// 		AccountType      *string         `json:"accountType"`
// 		AccountNumber    *string         `json:"accountNumber"`
// 		StartBalance     *float64        `json:"startBalance"`
// 		StartBalanceDate string          `json:"startBalanceDate"`
// 	}
// 	c.Bind(&AccountCreateRequest)
//
// 	// Parse StartBalanceDate string to time.Time
// 	startBalanceDate, err := time.Parse("2006-01-02", AccountCreateRequest.StartBalanceDate)
// 	if err != nil {
// 		// Handle error parsing date
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid start balance date",
// 		})
// 		return
// 	}
//
// 	// Create a Account
// 	account := domain.Account{
// 		Name:          AccountCreateRequest.Name,
// 		Currency:      AccountCreateRequest.Currency,
// 		AccountType:   AccountCreateRequest.AccountType,
// 		AccountNumber: AccountCreateRequest.AccountNumber,
// 		StartBalance:  AccountCreateRequest.StartBalance,
// 	}
//
// 	account.StartBalanceDate = &startBalanceDate
//
// 	result := initializers.DB.Create(&account)
// 	if result.Error != nil {
// 		c.Status(http.StatusBadRequest)
// 		return
// 	}
//
// 	// Return the Account
// 	c.JSON(http.StatusOK, gin.H{
// 		"account": account,
// 	})
// }
//
// func AccountsIndex(c *gin.Context) {
// 	// Get all Accounts
// 	var accounts []domain.Account
// 	initializers.DB.Find(&accounts)
//
// 	// Return the Accounts
// 	c.JSON(http.StatusOK, gin.H{
// 		"accounts": accounts,
// 	})
// }
//
// func AccountsShow(c *gin.Context) {
// 	// Get the id from the param
// 	id := c.Param("id")
//
// 	// Retreive the first record from db
// 	var account domain.Account
// 	initializers.DB.First(&account, id)
//
// 	// Return the account
// 	c.JSON(http.StatusOK, gin.H{
// 		"account": account,
// 	})
// }
//
// // AccountsUpdate updates a account by ID
// func AccountsUpdate(c *gin.Context) {
// 	// Get the id from the param
// 	id := c.Param("id")
//
// 	// Get data of request body
// 	var AccountCreateRequest struct {
// 		Name             string          `json:"name"`
// 		Currency         domain.Currency `json:"currency"`
// 		AccountType      *string         `json:"accountType"`
// 		AccountNumber    *string         `json:"accountNumber"`
// 		StartBalance     *float64        `json:"startBalance"`
// 		StartBalanceDate string          `json:"startBalanceDate"`
// 	}
// 	c.Bind(&AccountCreateRequest)
//
// 	// Retreive the first record from db
// 	var account domain.Account
// 	initializers.DB.First(&account, id)
//
// 	// Update
// 	initializers.DB.Model(&account).Updates(domain.Account{
// 		Name:          AccountCreateRequest.Name,
// 		Currency:      AccountCreateRequest.Currency,
// 		AccountType:   AccountCreateRequest.AccountType,
// 		AccountNumber: AccountCreateRequest.AccountNumber,
// 		StartBalance:  AccountCreateRequest.StartBalance,
// 	})
//
// 	// Parse StartBalanceDate string to time.Time
// 	startBalanceDate, err := time.Parse("2006-01-02", AccountCreateRequest.StartBalanceDate)
// 	if err != nil {
// 		// Handle error parsing date
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid start balance date",
// 		})
// 		return
// 	}
//
// 	account.StartBalanceDate = &startBalanceDate
//
// 	// Return the Account
// 	c.JSON(http.StatusOK, gin.H{
// 		"account": account,
// 	})
// }
//
// func AccountsDelete(c *gin.Context) {
// 	// Get the id from the param
// 	id := c.Param("id")
//
// 	initializers.DB.Delete(&domain.Account{}, id)
//
// 	// Return the Account
// 	c.Status(http.StatusOK)
// }
