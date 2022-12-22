package graph

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sRRRs-7/loose_style.git/cfg"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/dataloaders"
	"github.com/sRRRs-7/loose_style.git/token"
)

type Resolver struct {
	store       db.Store
	tokenMaker  token.Maker
	config      cfg.Config
	dataloaders dataloaders.Retriever
	tx          *pgxpool.Pool
}
