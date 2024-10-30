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
	e.stateCh <- e.state
	e.statusCh <- StatusRedWinner
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
