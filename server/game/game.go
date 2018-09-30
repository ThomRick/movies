package game

import (
	"errors"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const MovieDuration = 30000

type Game struct {
	movies       []movieInfo
	currentMovie atomic.Value
	players      atomic.Value
	playerMux    sync.Mutex
}

type Player struct {
	Name  string
	Score *int
}

type movieInfo struct {
	title string
	file  string
}

type Movie struct {
	File          string
	Title         string
	MsToNextMovie int64
}

type movie struct {
	idx              int
	currentStartTime time.Time
}

func New() *Game {
	g := Game{
		movies: []movieInfo{movieInfo{title: "Fight Club", file: "movie1.mp4"},
			movieInfo{title: "American Psyco", file: "movie2.mp4"},
			movieInfo{title: "American History X", file: "movie3.mp4"},
		},
	}

	go func() {
		incrementMovie := time.Tick(MovieDuration * time.Millisecond)
		g.currentMovie.Store(movie{currentStartTime: time.Now()})
		g.players.Store(make(map[string]Player))

		for {
			select {
			case <-incrementMovie:
				currentMovie := g.currentMovie.Load().(movie)
				currentMovie.idx++
				if len(g.movies) == currentMovie.idx {
					currentMovie.idx = 0
				}
				currentMovie.currentStartTime = time.Now()
				g.currentMovie.Store(currentMovie)
			}
		}
	}()

	return &g
}

func (g *Game) GetCurrentMovie() *Movie {
	m := g.currentMovie.Load().(movie)
	movieInfo := g.movies[m.idx]
	return &Movie{Title: movieInfo.title, File: movieInfo.file, MsToNextMovie: MovieDuration - time.Now().Sub(m.currentStartTime).Nanoseconds()/int64(time.Millisecond)}
}

func (g *Game) NewPlayer(playerName string) error {
	if _, ok := g.getPlayers()[playerName]; ok {
		return errors.New("player already exist")
	}

	g.playerMux.Lock()
	g.getPlayers()[playerName] = Player{Name: playerName, Score: new(int)}
	g.playerMux.Unlock()

	return nil
}

func (g *Game) GetPlayers() *[]Player {
	var playerList []Player
	players := g.getPlayers()
	for k := range players {
		playerList = append(playerList, players[k])
	}
	return &playerList
}

func (g *Game) GuessTitle(playerName string, titleName string) string {
	g.playerMux.Lock()
	m := g.currentMovie.Load().(movie)
	var message string
	if strings.EqualFold(g.movies[m.idx].title, titleName) {
		*(g.getPlayers()[playerName].Score)++
		message = "ok"
	} else {
		message = "ko"
	}
	g.playerMux.Unlock()
	return message
}

func (g *Game) getPlayers() map[string]Player {
	return g.players.Load().(map[string]Player)
}
