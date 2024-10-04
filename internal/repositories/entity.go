package repositories

import (
	"github.com/elgris/stom"
)

type tagMap string

const (
	tagInsert tagMap = "insert"
)

// Entity defines entity interface.
type Entity interface {
	// GetTableName get entity table name.
	GetTableName() string
}

func toMapInsert(row Entity) map[string]interface{} {
	return toMap(row, tagInsert)
}

func toMap(row Entity, tag tagMap) map[string]interface{} {
	tstom := stom.MustNewStom(row)
	tstom.SetTag(string(tag))
	m, err := tstom.ToMap(row)
	if err != nil {
		panic(err)
	}
	return m
}
