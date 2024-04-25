package controllers

import (
	"fmt"
	"go-bank/database"
	"go-bank/helpers"
	"go-bank/repositories"
	"go-bank/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var user structs.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusInternalServerError, "Failed to hash password", nil)
		return
	}

	// Replace the plain text password with the hashed version
	user.Password = string(hashedPassword)

	var userResponse structs.UserResponse
	userResponse, err = repositories.InsertUser(database.DbConnection, user)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Username already taken", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "User created successfully", userResponse)
}

func Login(ctx *gin.Context) {
	var loginInfo structs.LoginInfo
	var user structs.User

	err := ctx.ShouldBindJSON(&loginInfo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	user, err = repositories.GetUserByUsername(database.DbConnection, loginInfo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		helpers.APIResponse(ctx, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}

	var session structs.Session

	// Generate a new session ID
	session.ID = helpers.GenerateUUID()
	session.UserID = user.ID

	// Generate JWT token
	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusInternalServerError, "Error generating token", nil)
		return
	}
	session.Token = token

	session, err = repositories.InsertSession(database.DbConnection, session)

	if err != nil {
		helpers.APIResponse(ctx, http.StatusInternalServerError, "Failed to create session", nil)
		return
	}

	var loginResponse structs.LoginResponse

	loginResponse.ID = user.ID
	loginResponse.Username = user.Username
	loginResponse.Token = session.Token
	loginResponse.ExpiresAt = session.ExpiresAt

	helpers.APIResponse(ctx, http.StatusOK, "Login successful", loginResponse)
}

func ChangePassword(ctx *gin.Context) {
	var changePasswordInfo structs.ChangePasswordInfo
	var loginInfo structs.LoginInfo
	var user structs.User

	err := ctx.ShouldBindJSON(&changePasswordInfo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	loginInfo.Username = changePasswordInfo.Username

	user, err = repositories.GetUserByUsername(database.DbConnection, loginInfo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Username does not exist", nil)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePasswordInfo.OldPassword)); err != nil {
		helpers.APIResponse(ctx, http.StatusUnauthorized, "Invalid password", nil)
		return
	}

	// Check if the new password and confirm new password match
	if changePasswordInfo.NewPassword != changePasswordInfo.ConfirmNewPassword {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Passwords do not match", nil)
		return
	}

	// Hash the new password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePasswordInfo.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusInternalServerError, "Failed to hash password", nil)
		return
	}

	changePasswordInfo.NewPassword = string(hashedPassword)

	err = repositories.UpdatePassword(database.DbConnection, changePasswordInfo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusInternalServerError, "Failed to update password", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "Your password has been successfully updated", nil)
}

func DeleteUser(ctx *gin.Context) {
	var loginInfo structs.LoginInfo

	err := ctx.ShouldBindJSON(&loginInfo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	user, err := repositories.GetUserByUsername(database.DbConnection, loginInfo)
	if err != nil {
		helpers.APIResponse(ctx, http.StatusBadRequest, "Username does not exist", nil)
		return
	}

	// Delete all sessions related to the user
	err = repositories.DeleteUserSessions(database.DbConnection, user.ID)
	if err != nil {
		fmt.Println(err)
		helpers.APIResponse(ctx, http.StatusInternalServerError, "Failed to delete user's sessions", nil)
		return
	}

	err = repositories.DeleteUser(database.DbConnection, loginInfo)
	if err != nil {
		fmt.Println(err)
		helpers.APIResponse(ctx, http.StatusInternalServerError, "Failed to delete user", nil)
		return
	}

	helpers.APIResponse(ctx, http.StatusOK, "User deleted successfully", nil)
}
