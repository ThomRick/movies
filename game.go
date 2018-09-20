package main

import "time"

type game struct {
	currentMovie chan func(string)
	players      chan func(string)
}

func InitGame() game {
	movies := []string{"movie 1", "movie 2", "movie 3"}

	currentMovie := make(chan func(string))
	players := make(chan func(string))

	game := game{currentMovie, players}

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
