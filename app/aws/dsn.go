package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func BuildDSN(cfg aws.Config, dbName string) string {
	user := getSecret(cfg, "DB_USER")
	password := getSecret(cfg, "DB_PASSWORD")

	dbInfo, err := getDBInfo(cfg, dbName)
	if err != nil {
		return err.Error()
	}

	host := dbInfo["endpoint"]
	port := dbInfo["port"]
	database := dbInfo["name"]

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)
}
