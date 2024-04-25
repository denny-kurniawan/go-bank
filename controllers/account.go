package controllers

import (
	"fmt"
	"go-bank/database"
	"go-bank/helpers"
	"go-bank/repositories"
	"go-bank/structs"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAccountsByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")

	accounts, err := repositories.GetAccountsByUserID(database.DbConnection, userID)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to get accounts", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "Accounts retrieved successfully", accounts)
}

func GetAccountDetails(ctx *gin.Context) {
	var account structs.AccountDetails

	accountNo := ctx.Param("accountNo")
	account.AccountNo = accountNo

	userID := ctx.Param("userID")
	parsedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	account.UserID = parsedUserID

	account, err = repositories.GetAccountDetails(database.DbConnection, account)
	if err != nil {
		fmt.Println(err)
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to get account", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "Account retrieved successfully", account)
}

func CreateAccount(ctx *gin.Context) {
	var account structs.Account

	err := ctx.ShouldBindJSON(&account)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	// Create a new random source and generator
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Generate a 10-digit account number
	accountNumber := r.Int63n(9000000000) + 1000000000

	// Convert the account number to a string
	account.AccountNo = strconv.FormatInt(accountNumber, 10)

	// Insert the account into the database
	account, err = repositories.InsertAccount(database.DbConnection, account)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to create account", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "Account created successfully", account)
}

func DeleteAccountsByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")

	err := repositories.DeleteAccountsByUserID(database.DbConnection, userID)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to delete accounts", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "Accounts deleted successfully", nil)
}

func DeleteAccountByAccountNo(ctx *gin.Context) {
	var account structs.Account

	accountNo := ctx.Param("accountNo")
	account.AccountNo = accountNo

	userID := ctx.Param("userID")
	parsedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	account.UserID = parsedUserID

	err = repositories.DeleteAccountByAccountNo(database.DbConnection, account)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to delete account", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "Account deleted successfully", nil)
}
