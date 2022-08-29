package dbmodel

import (
	"github.com/uptrace/bun"
)

type Pokemon struct {
	bun.BaseModel `bun:"table:pokedex"`
	ID            int64    `bun:",pk,autoincrement"`
	Name          string   `bun:"name,notnull"`
	Description   string   `bun:"description,notnull"`
	Category      string   `bun:"category,notnull"`
	Type          []string `bun:"type,notnull"`
	Abilities     []string `bun:"abilities,notnull"`
}
