package repositories

import (
	"context"
	"github.com/Masterminds/squirrel"
	"test/infrastructure/db"
	"test/internal/entity"
)

type RateQuery interface {
	Insert(ctx context.Context, rate *entity.Rate) error
}

type rateQuery struct {
	conn db.Connection
}

func NewRateQuery(conn db.Connection) RateQuery {
	return &rateQuery{conn: conn}
}

func (q *rateQuery) Insert(ctx context.Context, rate *entity.Rate) error {
	_, err := squirrel.
		Insert(entity.GetRatesTable()).
		SetMap(toMapInsert(rate)).
		RunWith(q.conn.DB).
		PlaceholderFormat(squirrel.Dollar).
		ExecContext(ctx)

	return err
}
