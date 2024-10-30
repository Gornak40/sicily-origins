package cliengine

import (
	"log/slog"

	"github.com/Gornak40/sicily-origins/internal/gamecore"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Reset(seed int) {
	slog.Info("reset engine", slog.Int("seed", seed))
	// TODO
}

func (e *Engine) GetStatus() gamecore.GameStatus {
	// TODO
	return gamecore.StatusProgress
}
