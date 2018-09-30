package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mouminoux/movies/server/game"
)

const authCookieName = "authorization"

func main() {
	g := game.New()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	r.Use(middleware())

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/login")
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html.tmpl", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		playerName := c.PostForm("name")
		err := g.NewPlayer(playerName)
		if err != nil {
			c.Redirect(302, "/login?message="+err.Error())
			return
		}

		c.SetCookie(authCookieName, playerName, 0, "/", "", false, true)

		c.Redirect(302, "/game")
	})

	r.GET("/game", func(c *gin.Context) {
		c.HTML(200, "game.html.tmpl", nil)
	})

	r.GET("/game/current-movie", func(c *gin.Context) {
		c.JSON(200, gin.H{"movie": g.GetCurrentMovie()})
	})

	r.GET("/game/players", func(c *gin.Context) {
		c.JSON(200, gin.H{"players": g.GetPlayers()})
	})

	r.POST("/game/answer", func(c *gin.Context) {
		titleName := c.PostForm("title")
		playerName, _ := c.Cookie(authCookieName)
		c.JSON(200, gin.H{"message": g.GuessTitle(playerName, titleName)})
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
