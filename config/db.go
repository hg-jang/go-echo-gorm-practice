package config

import "fmt"

const (
	DB_USER		= ""
	DB_PASSWORD	= ""
	DB_NAME		= ""
	DB_HOST		= "localhost"
	DB_PORT		= "5432"
	DB_TYPE		= "postgres"
)

func GetDBType() string {
	return DB_TYPE
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DB_HOST,
		DB_PORT,
		DB_USER,
		DB_NAME,
		DB_PASSWORD,
	)

	return dataBase
}