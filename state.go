package main

import (
	"github.com/L0PE/gator/internal/config"
	"github.com/L0PE/gator/internal/database"
)

type state struct {
	conf *config.Config
	db *database.Queries
}
