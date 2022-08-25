package db_model

import (
	"github.com/uptrace/bun"
)

type Pokemon struct {
	bun.BaseModel `bun:"table:Pokedex"`
	ID            int64    `bun:",pk,autoincrement"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Category      string   `json:"category"`
	Type          []string `json:"type"`
	Abilities     []string `json:"abilities"`
}
