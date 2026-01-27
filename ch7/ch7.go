package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Piece struct {
	Color_ string
	X_     int
	Y_     int
}

// Methods can only be defined at package level
func (p Piece) Color() string {
	return p.Color_
}

type Tree struct {
	val_          int
	left_, right_ *Tree
}

func (t *Tree) Insert(val int) *Tree {
	if t == nil {
		return &Tree{val_: val}
	}
	if val < t.val_ {
		t.left_ = t.left_.Insert(val)
	} else if val > t.val_ {
		t.right_ = t.right_.Insert(val)
	}
	return t

}

func (t *Tree) Contains(val int) bool {
	switch {
	case t == nil:
		return false
	case val < t.val_:
		return t.left_.Contains(val)
	case val > t.val_:
		return t.right_.Contains(val)
	default:
		return true
	}
}

const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4
	Field5 = iota
)

type Queen Piece

type Team struct {
	name    string
	members []string
}

type League struct {
	teams []Team
	wins  map[string]int
}

func (l *League) MatchResult(t1 string, t2 string, score1 int, score2 int) {
	if score1 > score2 {
		l.wins[t1]++
	} else if score2 > score1 {
		l.wins[t2]++
	}
}

func (l *League) Ranking() []string {
	type Pair struct {
		name string
		wins int
	}
	board := []Pair{}
	for _, team := range l.teams {
		value := l.wins[team.name]
		board = append(board, Pair{name: team.name, wins: value})
	}

	sort.Slice(board, func(i, j int) bool {
		return board[i].wins > board[j].wins
	})

	res := []string{}

	for _, team := range board {
		res = append(res, team.name)
	}

	return res

}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	var s string = "Printing teams\n"
	io.WriteString(w, s)
	var res []string = r.Ranking()
	for _, val := range res {
		var x string = val + "\n"
		io.WriteString(w, x)
	}
}

func main() {
	fmt.Println("Chapter 7: Types, Methods, Interfaces")
	var t *Tree
	x := []int{1, 2, 3, 4, 5}
	for _, val := range x {
		fmt.Printf("Inserting %d\n", val)
		t = t.Insert(val)
	}
	for index := range x {
		if t.Contains(index) {
			fmt.Printf("Tree contains %d\n", index)
		} else {
			fmt.Printf("Tree does not contain %d\n", index)
		}
	}

	fmt.Println(Field1, Field2, Field3, Field4, Field5)

	var teams_in []Team = []Team{{"Esben", []string{"Esben"}}, {"Ben", []string{"Ben"}}}
	var l League
	l.teams = teams_in
	l.wins = make(map[string]int)
	l.MatchResult("Esben", "Ben", 1, 0)

	// this works because *League implements Ranker interface! (I think)
	RankPrinter(&l, os.Stdout)

}
