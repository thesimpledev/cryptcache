package main

import (
	"log/slog"
	"os"
)

type application struct {
	logger *slog.Logger
	nouns  map[string]func(string, []string)
}

func newApplication() application {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	var nouns = map[string]func(string, []string){
		"project": projectHandler,
		"profile": profileHandler,
		"secret":  secretHandler,
	}

	app := &application{
		logger: logger,
		nouns:  nouns,
	}

	return *app

}

func (app *application) parse() {
	app.commandLenValidator(len(os.Args))

	noun := os.Args[1]
	verb := os.Args[2]
	args := os.Args[3:]

	if cmd, exists := app.nouns[noun]; exists {
		cmd(verb, args)
		os.Exit(0)
	}

	app.logger.Error("Invalid Command")
	os.Exit(1)

}

func (app *application) commandLenValidator(length int) {
	errors := map[int]string{
		1: "Expected a noun (e.g., project, profile, value)",
		2: "Expected a verb (e.g., create, delete, update, merge)",
		3: "Expected one or more arguments",
	}

	if msg, ok := errors[length]; ok {
		app.logger.Error(msg)
		os.Exit(1)
	}

}
