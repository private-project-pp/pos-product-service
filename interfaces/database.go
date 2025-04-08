package interfaces

import (
	_ "github.com/lib/pq"
	"github.com/private-project-pp/product-rpc-service/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(configs config.ConfigApp) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(configs.Db.Host))
	if err != nil {
		return db, err
	}
	return db, nil
}
