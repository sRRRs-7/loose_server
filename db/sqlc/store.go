package db

import "github.com/jackc/pgx/v4/pgxpool"

type Store interface {
	Querier
	// oooTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

type SQLStore struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *SQLStore {
	return &SQLStore{
		Queries: New(pool),
		db:      pool,
	}
}
