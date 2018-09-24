package main

import (
	"fmt"
	"strings"
	"time"
)

const MovieDuration = 5000

type game struct {
	currentMovie chan func(string, int64)
	NewPlayer    chan func() string
	GetPlayers   chan func([]player)
	GuessTitle   chan func(map[string]player, string)
}

type player struct {
	Name  string
	Score *int
}

type movie struct {
	Title         string
	MsToNextMovie int64
}

func InitGame() game {

	movies := []string{"movie 1", "movie 2", "movie 3"}
	playerMap := make(map[string]player)

	currentMovie := make(chan func(string, int64))
	newPlayer := make(chan func() string)
	getPlayers := make(chan func([]player))
	guessTitle := make(chan func(map[string]player, string))
	game := game{currentMovie, newPlayer, getPlayers, guessTitle}

	go func() {
		incrementMovie := time.Tick(MovieDuration * time.Millisecond)
		currentMovieStartTime := time.Now()
		currentMovieIndex := 0

		for {
			select {
			case <-incrementMovie:
				currentMovieStartTime = time.Now()
				currentMovieIndex++
				if len(movies) == currentMovieIndex {
					currentMovieIndex = 0
				}
			case task := <-currentMovie:
				timeToNextMovie := MovieDuration - time.Now().Sub(currentMovieStartTime).Nanoseconds()/int64(time.Millisecond)
				task(movies[currentMovieIndex], timeToNextMovie)
			case task := <-newPlayer:
				playerName := task()
				playerMap[playerName] = player{playerName, new(int)}
			case task := <-getPlayers:
				playerList := make([]player, len(playerMap))
				for k := range playerMap {
					playerList = append(playerList, playerMap[k])
				}
				task(playerList)
			case task := <-guessTitle:
				task(playerMap, movies[currentMovieIndex])
			}
		}
	}()

	return game
}

func GetCurrentMovie(game game) movie {
	currentMovie := make(chan movie)
	game.currentMovie <- func(movieName string, timeToNextMovie int64) {
		currentMovie <- movie{movieName, timeToNextMovie}
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

func GetPlayers(game game) []player {
	getPlayers := make(chan []player)
	game.GetPlayers <- func(players []player) {
		getPlayers <- players
	}
	players := <-getPlayers
	close(getPlayers)
	return players
}

func GuessTitle(game game, playerName string, titleName string) string {
	guessTitle := make(chan string)
	game.GuessTitle <- func(playerMap map[string]player, currentMovieTitle string) {
		p := playerMap[playerName]
		fmt.Printf("%s %s\n", currentMovieTitle, titleName)
		if strings.EqualFold(currentMovieTitle, titleName) {
			*(p.Score) = *(p.Score) + 1
			guessTitle <- "ok"
		} else {
			guessTitle <- "ko"
		}
	}
	message := <-guessTitle
	close(guessTitle)
	return message
}
