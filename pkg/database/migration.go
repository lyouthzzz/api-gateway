package database

import (
	"github.com/jinzhu/gorm"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/model"
)

func MigrateTables(tx *gorm.DB, tables ...model.HasTableName) error {
	for _, t := range tables {
		if tx = tx.AutoMigrate(t); tx.Error != nil {
			return tx.Error
		}
		if hi, ok := t.(model.HasIndexes); ok {
			for indexName, columns := range hi.Indexes() {
				if tx = tx.Model(t).AddIndex(indexName, columns...); tx.Error != nil {
					return tx.Error
				}
			}
		}
		if hi, ok := t.(model.HasUniqueIndexes); ok {
			for indexName, columns := range hi.UniqueIndexes() {
				if tx = tx.Model(t).AddUniqueIndex(indexName, columns...); tx.Error != nil {
					return tx.Error
				}
			}
		}
	}
	for _, t := range tables {
		if hf, ok := t.(model.HasForeignKeys); ok {
			for col, foreignKey := range hf.ForeignKeys() {
				if tx = tx.Model(t).AddForeignKey(col, foreignKey, "RESTRICT", "CASCADE"); tx.Error != nil {
					return tx.Error
				}
			}
		}
	}
	return nil
}
