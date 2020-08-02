package postgres

import (
	"github.com/Rahul12344/skelego"

	"github.com/Rahul12344/skelego/services/storage"
)

//Migrate Initialization-time migration of schemas
func Migrate(logger skelego.Logging, schemas ...storage.Schema) {
	for _, schema := range schemas {
		schema.Migrate(logger)
		logger.LogEvent("Migrated schema: %s", schema.TableName())
	}
}
