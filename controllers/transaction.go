package controllers

import (
	"go-bank/database"
	"go-bank/helpers"
	"go-bank/repositories"
	"go-bank/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Deposit(ctx *gin.Context) {
	var transactionInfo structs.TransactionInfo
	var transaction structs.Transaction
	var account structs.AccountBalance

	err := ctx.ShouldBindJSON(&transactionInfo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	if transactionInfo.Amount < 50000 {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Amount must be at least 50000", nil)
		return
	}

	account.AccountNo = transactionInfo.AccountNo

	account, err = repositories.GetAccountBalance(database.DbConnection, account)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to deposit", nil)
		return
	}

	transaction.AccountNo = transactionInfo.AccountNo
	transaction.TransactionType = "deposit"
	transaction.Amount = transactionInfo.Amount
	transaction.Description = transactionInfo.Description

	transaction, err = repositories.InsertTransaction(database.DbConnection, transaction)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to deposit", nil)
		return
	}

	account.Balance += transactionInfo.Amount

	err = repositories.UpdateAccountBalance(database.DbConnection, account)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to deposit", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "Deposit successful", transaction)
}

func Withdraw(ctx *gin.Context) {
	var transactionInfo structs.TransactionInfo
	var transaction structs.Transaction
	var account structs.AccountBalance

	err := ctx.ShouldBindJSON(&transactionInfo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	if transactionInfo.Amount < 50000 {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Amount must be at least 50000", nil)
		return
	}

	account.AccountNo = transactionInfo.AccountNo

	account, err = repositories.GetAccountBalance(database.DbConnection, account)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to withdraw", nil)
		return
	}

	if account.Balance <= 100000 || account.Balance-transactionInfo.Amount <= 100000 {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Minimum balance must be 100000", nil)
		return
	}

	if account.Balance < transactionInfo.Amount {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Insufficient balance", nil)
		return
	}

	transaction.AccountNo = transactionInfo.AccountNo
	transaction.TransactionType = "withdraw"
	transaction.Amount = transactionInfo.Amount
	transaction.Description = transactionInfo.Description

	transaction, err = repositories.InsertTransaction(database.DbConnection, transaction)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to withdraw", nil)
		return
	}

	account.AccountNo = transactionInfo.AccountNo
	account.Balance -= transactionInfo.Amount

	err = repositories.UpdateAccountBalance(database.DbConnection, account)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to withdraw", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "Withdraw successful", transaction)
}

func Transfer(ctx *gin.Context) {
	var transferInfo structs.TransferInfo
	var transaction structs.Transaction
	var accountFrom, accountTo structs.AccountBalance

	err := ctx.ShouldBindJSON(&transferInfo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	if transferInfo.Amount <= 0 {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Amount must be greater than 0", nil)
		return
	}

	accountFrom.AccountNo = transferInfo.FromAccountNo
	accountTo.AccountNo = transferInfo.ToAccountNo

	accountFrom, err = repositories.GetAccountBalance(database.DbConnection, accountFrom)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to transfer", nil)
		return
	}

	accountTo, err = repositories.GetAccountBalance(database.DbConnection, accountTo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to transfer", nil)
		return
	}

	if accountFrom.Balance <= 100000 || accountFrom.Balance-transferInfo.Amount <= 100000 {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Minimum balance must be 100000", nil)
		return
	}

	if accountFrom.Balance < transferInfo.Amount {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Insufficient balance", nil)
		return
	}

	transaction.AccountNo = transferInfo.FromAccountNo
	transaction.TransactionType = "send"
	transaction.Amount = transferInfo.Amount
	transaction.Description = transferInfo.Description

	transaction, err = repositories.InsertTransaction(database.DbConnection, transaction)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to transfer", nil)
		return
	}

	transaction.AccountNo = transferInfo.ToAccountNo
	transaction.TransactionType = "receive"
	transaction, err = repositories.InsertTransaction(database.DbConnection, transaction)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to transfer", nil)
		return
	}

	accountFrom.AccountNo = transferInfo.FromAccountNo
	accountFrom.Balance -= transferInfo.Amount

	err = repositories.UpdateAccountBalance(database.DbConnection, accountFrom)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to transfer", nil)
		return
	}

	accountTo.AccountNo = transferInfo.ToAccountNo
	accountTo.Balance += transferInfo.Amount

	err = repositories.UpdateAccountBalance(database.DbConnection, accountTo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Failed to transfer", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "Transfer successful", transferInfo)
}
