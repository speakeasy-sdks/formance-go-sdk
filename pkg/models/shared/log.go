package shared

import (
	"time"
)

type LogTypeEnum string

const (
	LogTypeEnumNewTransaction LogTypeEnum = "NEW_TRANSACTION"
	LogTypeEnumSetMetadata    LogTypeEnum = "SET_METADATA"
)

type Log struct {
	Data map[string]interface{} `json:"data"`
	Date time.Time              `json:"date"`
	Hash string                 `json:"hash"`
	ID   int64                  `json:"id"`
	Type LogTypeEnum            `json:"type"`
}
