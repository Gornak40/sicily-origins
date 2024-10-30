package gamecore

type GameStatus int // TODO: add deck nation

const (
	StatusRedWinner GameStatus = iota
	StatusBlackWinner
	StatusDraw
)

type Player struct {
	Deck      DeckHand
	Line      DistrictLine
	Score     int
	CapoBoost bool
	GangTrap  bool
}

func (p *Player) CapoInt() int {
	if p.CapoBoost {
		return 2
	}

	return 0
}

type GameState struct {
	Red     Player
	Black   Player
	Turn    int
	Pending int
}

func initState(_ int) GameState {
	// TODO: use real seed
	var d DeckHand
	d.Append(CardAce)
	d.Append(CardDon)
	d.Append(CardJoker)
	d.Append(CardSoldier)
	d.Append(CardKiller)
	d.Append(CardGang)
	d.Append(CardOfficer)
	d.Append(CardKiller)

	return GameState{
		Red: Player{
			Deck: d,
		},
		Black: Player{
			Deck: d,
		},
	}
}

type GameRound struct {
	RedCardIdx   int
	BlackCardIdx int
}
