package main

import (
	"net/http"

	"github.com/Sharon-Liu-go/go-training-test/crawler"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	userStorage := Login{"Sharon", "1111"}

	router := gin.Default()
	store := cookie.NewStore([]byte("secret00000"))

	router.Use(sessions.Sessions("mySession", store))
	router.LoadHTMLGlob("views/*.html")

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{"title": "login Page"})
	})

	router.POST("/login", func(c *gin.Context) {
		var user Login

		if err := c.ShouldBind(&user); err != nil {
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

		session := sessions.Default(c)
		if session.Get("account") != user.Account {
			// 设置session数据
			session.Set("account", user.Account)
			session.Save()
		}
		c.Redirect(http.StatusFound, "/welcome")
	})

	router.GET("/welcome", func(c *gin.Context) {
		account := userStorage.Account

		session := sessions.Default(c)
		if session.Get("account") != account {
			c.JSON(http.StatusOK, "無權瀏覽該頁面!")
			return
		}

		imgSrc, count := crawler.CollyUseTemplate()

		c.HTML(http.StatusOK, "index.html", gin.H{"account": account, "imgsSrc": imgSrc, "count": count})
		/**rtn := ResFmt{}
		rtn.Code = 0
		rtn.Msg = "success"
		rtn.Data = gin.H{"account": account, "imgsSrc": imgSrc, "count": count}
		c.JSON(http.StatusOK, rtn)**/
	})

	router.GET("/logout", func(c *gin.Context) {
		account := userStorage.Account

		session := sessions.Default(c)
		if session.Get("account") != account {
			c.JSON(http.StatusOK, "無權瀏覽該頁面!")
			return
		}
		session.Clear() //移除session
		session.Save()  //儲存
		c.Redirect(http.StatusFound, "/login")
	})

	router.Run(":3000") //放在router. API的前面，就會404
}
