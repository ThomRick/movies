package main

import "time"

func initGame() chan func(string) {
	movies := []string{"movie 1", "movie 2", "movie 3"}

	gameChan := make(chan func(string))

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
			case task := <-gameChan:
				task(movies[currentMovieIndex])
			}
		}
	}()

	return gameChan
}

func GetCurrentMovie(game chan func(string)) string {
	currentMovieChan := make(chan string)
	game <- func(movie string) {
		currentMovieChan <- movie
	}
	movie := <-currentMovieChan
	close(currentMovieChan)
	return movie
}
