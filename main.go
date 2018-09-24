package main

import (
	"github.com/gin-gonic/gin"
)

const authCookieName = "authorization"

func main() {
	game := InitGame()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	r.Use(middleware())

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html.tmpl", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		playerName := c.PostForm("name")
		NewPlayer(game, playerName)
		c.SetCookie(authCookieName, playerName, 0, "/", "", false, true)

		c.Redirect(302, "/game")
	})

	r.GET("/game", func(c *gin.Context) {
		c.HTML(200, "game.html.tmpl", nil)
	})

	r.GET("/game/current-movie", func(c *gin.Context) {
		c.JSON(200, gin.H{"movie": GetCurrentMovie(game)})
	})

	r.GET("/game/players", func(c *gin.Context) {
		c.JSON(200, gin.H{"players": GetPlayers(game)})
	})

	r.Static("/static", "static")

	r.Run() // listen and serve on 0.0.0.0:8080
}

func middleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie(authCookieName)
		if cookie == "" && c.Request.URL.Path != "/login" {
			c.Redirect(302, "/login")
			return
		}
		c.Next()
	}
}
