package app

import (
	"gobarber/internal/infra/database"
	"gobarber/internal/infra/http"
)

func Start() error {
	if err := database.Connect(); err != nil {
		return err
	}

	srv := http.NewServer()
	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
