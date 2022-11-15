package main

import (
	"net/http"

	"github.com/Sharon-Liu-go/go-training-test/crawler"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Account  string `form:"account" json:"account" binding:"required" `
	Password string `form:"password" json:"password" binding:"required"`
}

type ResFmt struct {
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
			rtn := ResFmt{}
			rtn.Code = 1
			rtn.Msg = err.Error()
			c.JSON(http.StatusBadRequest, rtn)
			return
		}

		if user.Account != userStorage.Account || user.Password != userStorage.Password {
			rtn := ResFmt{}
			rtn.Code = 1
			rtn.Msg = "unauthorized"
			c.JSON(http.StatusUnauthorized, rtn)
			return
		}

		c.Redirect(http.StatusFound, "/welcome")
	})

	router.LoadHTMLGlob("views/*.html")

	router.GET("/welcome", func(c *gin.Context) {
		account := userStorage.Account
		imgSrc, count := crawler.CollyUseTemplate()
		rtn := ResFmt{}
		rtn.Code = 0
		rtn.Msg = "success"
		rtn.Data = gin.H{"account": account, "imgsSrc": imgSrc, "count": count}
		c.HTML(http.StatusOK, "index.html", gin.H{"account": account, "imgsSrc": imgSrc, "count": count})
		//c.JSON(http.StatusOK, rtn)
	})

	router.Run(":3000") //放在router. API的前面，就會404
}
