package entity

import "time"

type Rate struct {
	Id        int64     `db:"id"`
	Market    string    `db:"market" insert:"market"`
	Ask       float64   `db:"ask" insert:"ask"`
	Bid       float64   `db:"bid" insert:"bid"`
	CreatedAt time.Time `db:"created_at" insert:"created_at"`
}

func GetRatesTable() string {
	return "rates"
}

func (p *Rate) GetTableName() string {
	return GetRatesTable()
}
