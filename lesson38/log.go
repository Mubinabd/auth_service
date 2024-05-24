package main

import (
	"log/slog"
)

type t struct {
	R int
	A string
}

func main() {
	t := t{12, "sdfds"}
	slog.Info("hello world", "user", t, "num", slog.Int("son", 543))
	slog.Error("hello world", "user", t)

	slog.SetLogLoggerLevel(slog.LevelDebug)

	//l := slog.NewJSONHandler(nil, &slog.HandlerOptions{Level: slog.LevelDebug})
	//
	//l,

	//log := slog.Logger{}
	//slog.SetDefault(&log)
	slog.Debug("hello world", "user", t)
}
