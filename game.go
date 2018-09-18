package main

import "time"

func initGame() chan func(string) {
	movies := []string{"movie 1", "movie 2", "movie 3"}

	chanstrings := make(chan func(string))

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
			case task := <- chanstrings:
				task(movies[currentMovieIndex])
			}
		}
	}()

	return chanstrings
}
