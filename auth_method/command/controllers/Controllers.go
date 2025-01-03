package controllers

// PRECISA REFATORAR O CODIGO E REPASSAR OS ERROS DOS STATUS HTTP.

import (
	"log"
	"net/http"
	"net/url"

	"github.com/auth/command/database/mysql/connection"
	"github.com/auth/command/model"
	"github.com/auth/command/util"
	"github.com/gin-gonic/gin"
)

const valueToGenerateToken int = 64

// Register
func HandleRegister(c *gin.Context) {
	var user model.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.ValidateUsersStruct(&user); err != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
		return
	}

	connection.DATABASE.Where(&model.Users{Username: user.Username}).First(&user)

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	if password, err := util.HashPassword(user.HashedPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {

		user.HashedPassword = password
		connection.DATABASE.Create(&user)

		c.JSON(http.StatusCreated, gin.H{
			"message": "successfully created",
		})
		return
	}
}

// Login
func HandleLogin(c *gin.Context) {
	var user model.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := user.ID
	password := user.HashedPassword

	connection.DATABASE.Where(&model.Users{Username: user.Username}).First(&user)
	if user.ID != 0 {
		if err := util.CheckHashPassword(user.HashedPassword, password); err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": "check username or password",
			})
			return

		}

		user.SessionToken = util.GenerateRandomToken(valueToGenerateToken)
		user.CSRFToken = util.GenerateRandomToken(valueToGenerateToken * 2)
		c.SetCookie("session_token", user.SessionToken, 60, "/", "localhost", false, true)
		c.SetCookie("csrf_token", user.CSRFToken, 60, "/", "localhost", false, false)

		connection.DATABASE.First(&user, id)

		connection.DATABASE.Save(&user)
		c.JSON(http.StatusOK, gin.H{
			"message": "login successfully",
		})
		return

	}

}

// Protected
func HandleProtected(c *gin.Context) {
	var user model.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	connection.DATABASE.Where(&model.Users{Username: user.Username}).First(&user)
	if user.ID != 0 {

		cookie, err := c.Request.Cookie("session_token")

		cookieValue, er := url.QueryUnescape(cookie.Value)
		if er != nil {
			log.Println(er.Error())
		}

		if err != nil || cookie.Value == "" || cookieValue != user.SessionToken {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

		}

		csrf := c.GetHeader("X-CSRF-Token")
		if csrf != user.CSRFToken || csrf == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

		}

		c.JSON(http.StatusOK, gin.H{
			"message": "CSRF validation sucessful! Welcome, " + user.Username,
		})

	}

}

// Logout
func HandleLogout(c *gin.Context) {
	var user model.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	connection.DATABASE.Where(&model.Users{Username: user.Username}).First(&user)
	if user.ID != 0 {

		cookie, err := c.Request.Cookie("session_token")

		cookieValue, er := url.QueryUnescape(cookie.Value)
		if er != nil {
			log.Println(er.Error())
		}

		if err != nil || cookie.Value == "" || cookieValue != user.SessionToken {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

		}

		csrf := c.GetHeader("X-CSRF-Token")
		if csrf != user.CSRFToken || csrf == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

		}
		connection.DATABASE.First(&user, user.ID)
		user.SessionToken = ""
		user.CSRFToken = ""
		c.SetCookie("session_token", user.SessionToken, -60, "/", "localhost", false, true)
		c.SetCookie("csrf_token", user.CSRFToken, -60, "/", "localhost", false, false)

		connection.DATABASE.Save(&user)
		c.JSON(http.StatusOK, gin.H{
			"message": "logout successfully",
		})
		return
	}

}
