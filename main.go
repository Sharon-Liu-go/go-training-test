package main

import (
	"fmt"
	"net/http"

	"github.com/Sharon-Liu-go/go-training-test/crawler"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Account  string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type resFmt struct {
	Code int
	Msg  string
	Data interface{}
}

func main() {
	userStorage := Login{"Sharon", "wohaha1111"}

	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var user Login
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.Account != userStorage.Account || user.Password != userStorage.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.Redirect(http.StatusFound, "/welcome")
	})

	router.GET("/welcome", func(c *gin.Context) {
		account := userStorage.Account
		imgsSrc, count := crawler.CollyUseTemplate()
		fmt.Println("imgsSrc:", imgsSrc)
		fmt.Println("count:", count)
		rtn := resFmt{}
		rtn.Code = 0
		rtn.Msg = "success"
		rtn.Data = gin.H{"account": account, "imgsSrc": imgsSrc}
		c.JSON(http.StatusOK, rtn)
	})

	router.Run(":3000")
}
