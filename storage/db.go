package storage

import (
	config "go-echo-gorm-practice/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func NewDB(tableSchema string) *gorm.DB {
	var err error

	var connectionString string = config.GetPostgresConnectionString()

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: tableSchema,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return DB
}

func GetDB() *gorm.DB {
	return DB
}