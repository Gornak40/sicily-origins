package gamecore

type GameStatus int

const (
	StatusProgress GameStatus = iota
	StatusRedWinner
	StatusBlackWinner
	StatusDraw
)
