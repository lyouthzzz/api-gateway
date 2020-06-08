package model

import (
	"time"
)

type DeleteTs uint64

const (
	// 标志位:没删除
	DeletedOff DeleteTs = 0
	// 标志位:已删除
	DeletedOn DeleteTs = 1
)

type Model struct {
	Id        uint32    `gorm:"column:id;primary_key;auto_increment"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp(6);not null;default:now(6)"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp(6);not null;default:now(6)"`
	DeletedTs DeleteTs  `gorm:"column:deleted_ts;type:bigint unsigned;not null;default:0"`
}

type HasTableName interface {
	TableName() string
}

type HasIndexes interface {
	Indexes() map[string][]string
}

type HasUniqueIndexes interface {
	UniqueIndexes() map[string][]string
}

type HasForeignKeys interface {
	ForeignKeys() map[string]string
}
