package gamecore

type Engine struct {
	state GameState

	stateCh  chan GameState
	statusCh chan GameStatus
	roundCh  chan GameRound
}

func NewEngine(seed int) *Engine {
	return &Engine{
		state:    initState(seed),
		stateCh:  make(chan GameState),
		statusCh: make(chan GameStatus),
		roundCh:  make(chan GameRound),
	}
}

func (e *Engine) Run() {
gameLoop:
	e.stateCh <- e.state
	round := <-e.roundCh
	e.battle(round)
	switch {
	case e.state.Red.Score >= 4:
		e.statusCh <- StatusRedWinner
	case e.state.Black.Score >= 4:
		e.statusCh <- StatusBlackWinner
	case e.state.Turn == 8:
		e.statusCh <- StatusDraw
	default:
		goto gameLoop
	}
}

func (e *Engine) State() <-chan GameState {
	return e.stateCh
}

func (e *Engine) Status() <-chan GameStatus {
	return e.statusCh
}

func (e *Engine) Round() chan<- GameRound {
	return e.roundCh
}

//nolint:funlen,cyclop // It's the most important function in all project.
func (e *Engine) battle(r GameRound) {
	e.state.Turn++

	e.state.Red.GangTrap = false
	rc := e.state.Red.Deck.DropCard(r.RedCardIdx)
	rp := rc.Power() + e.state.Red.CapoInt()
	e.state.Red.CapoBoost = false
	rs := e.state.Pending + 1

	e.state.Black.GangTrap = false
	bc := e.state.Black.Deck.DropCard(r.BlackCardIdx)
	bp := bc.Power() + e.state.Black.CapoInt()
	e.state.Black.CapoBoost = false
	bs := e.state.Pending + 1

	if rc == CardKiller || bc == CardKiller {
		goto final
	}
	if rc == CardJoker || bc == CardJoker {
		rp, bp = 0, 0
	}
	if rc == CardAce && bc == CardDon {
		rs += 3
		rp = bp + 1
	}
	if bc == CardAce && rc == CardDon {
		bs += 3
		bp = rp + 1
	}
	if rc == CardGang && bc != CardGang {
		e.state.Black.GangTrap = true
	}
	if bc == CardGang && rc != CardGang {
		e.state.Red.GangTrap = true
	}
	if (rc == CardOfficer && bc != CardDon) || (bc == CardOfficer && rc != CardDon) {
		rp = -rp
		bp = -bp
	}
	if rc == CardSoldier {
		rs++
	}
	if bc == CardSoldier {
		bs++
	}
	if rc == CardCapo {
		e.state.Red.CapoBoost = true
	}
	if bc == CardCapo {
		e.state.Black.CapoBoost = true
	}

final:
	switch {
	case rp > bp:
		e.state.Red.Score += rs
		e.state.Pending = 0
	case bp > rp:
		e.state.Black.Score += bs
		e.state.Pending = 0
	default: // draw
		e.state.Pending++
	}
}
