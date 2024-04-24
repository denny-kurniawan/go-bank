package controllers

import (
	"fmt"
	"go-bank/database"
	"go-bank/helpers"
	"go-bank/repositories"
	"go-bank/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	users, err := repositories.GetUsers(database.DbConnection)

	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
	} else {
		helpers.APIResponse(ctx, http.StatusOK, "Users retrieved successfully", users)
	}
}

func CreateUser(ctx *gin.Context) {
	var user structs.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	fmt.Println(user)

	user, err = repositories.InsertUser(database.DbConnection, user)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusInternalServerError, "Failed to create user", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "User created successfully", user)
}

func UpdateUser(ctx *gin.Context) {
	var user structs.User

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	user.ID = uint64(id)

	user, err = repositories.UpdateUser(database.DbConnection, user)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusInternalServerError, "Failed to update user", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "User updated successfully", user)
}

func DeleteUser(ctx *gin.Context) {
	var user structs.User

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	user.ID = uint64(id)

	err = repositories.DeleteUser(database.DbConnection, user)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusInternalServerError, "Failed to delete user", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "User deleted successfully", nil)
}
