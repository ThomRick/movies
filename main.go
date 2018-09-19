package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const s = "authorization"

type user struct {
	Name  string
	Score int
}

func main() {
	gameChan := initGame()

	var usersMap = make(map[string]user)
	usersMap["toto"] = user{"toto", 1337}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	r.Use(middleware())

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html.tmpl", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		name := c.PostForm("name")
		usersMap[name] = user{name, 0}
		c.SetCookie(s, name, 0, "/", "", false, true)

		c.Redirect(302, "/game")
	})

	r.GET("/game", func(c *gin.Context) {
		c.HTML(200, "game.html.tmpl", GetCurrentMovie(gameChan))
	})

	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	r.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		c.HTML(200, "user.html.tmpl", usersMap[name])
	})

	r.POST("/user/add", func(c *gin.Context) {
		name := c.PostForm("name")
		score, err := strconv.Atoi(c.PostForm("score"))
		if err != nil {
			c.Error(err)
			return
		}

		println("%s %i", name, score)

		usersMap[name] = user{name, score}
		c.Redirect(302, "/users")
	})

	r.GET("/users", func(c *gin.Context) {
		c.HTML(200, "users.html.tmpl", usersMap)
	})

	r.Static("/static", "binaries")

	r.Run() // listen and serve on 0.0.0.0:8080
}

func middleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie(s)
		if cookie == "" && c.Request.URL.Path != "/login" {
			c.Redirect(302, "/login")
			return
		}
		c.Next()
	}
}
