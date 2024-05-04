package job

import "log/slog"

type Hello struct {
}

func (Hello) Run() {
	slog.Info("Hello world")
}
