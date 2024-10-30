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

func (e *Engine) battle(_ GameRound) {
	// TODO: impl
	e.state.Turn++
}
