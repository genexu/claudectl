package main

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/fx"

	"claudectl/internal/utils"
	"claudectl/internal/view"
)

func RunTUI(
	lc fx.Lifecycle,
	shutdowner fx.Shutdowner,
	model *view.Model,
	logger *utils.Logger,
	cfg Config,
) {
	logger.Info("claudectl starting", "version", version)

	p := tea.NewProgram(model, tea.WithAltScreen())
	model.SetProgram(p)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("starting TUI")

			go func() {
				if _, err := p.Run(); err != nil {
					logger.Error("TUI crashed", "error", err)
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
				logger.Info("TUI exited, triggering shutdown")
				shutdowner.Shutdown()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("stopping TUI")

			p.Quit()

			select {
			case <-model.Shutdown():
				logger.Info("TUI shutdown complete")
			case <-ctx.Done():
				logger.Warn("TUI shutdown timeout")
				return ctx.Err()
			}

			return nil
		},
	})
}
