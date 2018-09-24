package main

import "time"

type game struct {
	currentMovie chan func(string)
	NewPlayer    chan func() string
	GetPlayers   chan func(map[string]player)
}

type player struct {
	Name  string
	Score int
}

func InitGame() game {

	movies := []string{"movie 1", "movie 2", "movie 3"}
	playerMap := make(map[string]player)

	currentMovie := make(chan func(string))
	newPlayer := make(chan func() string)
	getPlayers := make(chan func(map[string]player))
	game := game{currentMovie, newPlayer, getPlayers}

	go func() {
		incrementMovie := time.Tick(5 * time.Second)
		currentMovieIndex := 0

		for {
			select {
			case <-incrementMovie:
				currentMovieIndex++
				if len(movies) == currentMovieIndex {
					currentMovieIndex = 0
				}
			case task := <-currentMovie:
				task(movies[currentMovieIndex])
			case task := <-newPlayer:
				playerName := task()
				playerMap[playerName] = player{playerName, 0}
			case task := <-getPlayers:
				task(playerMap)
			}
		}
	}()

	return game
}

func GetCurrentMovie(game game) string {
	currentMovie := make(chan string)
	game.currentMovie <- func(movie string) {
		currentMovie <- movie
	}
	movie := <-currentMovie
	close(currentMovie)
	return movie
}

func NewPlayer(game game, playerName string) {
	game.NewPlayer <- func() string {
		return playerName
	}
}

func GetPlayers(game game) map[string]player {
	getPlayers := make(chan map[string]player)
	game.GetPlayers <- func(players map[string]player) {
		getPlayers <- players
	}
	players := <-getPlayers
	close(getPlayers)
	return players
}
