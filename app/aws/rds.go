package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/drshooby/secrets-manager-practice/utils"
)

func getDBInfo(cfg aws.Config, dbName string) (map[string]string, error) {

	client := rds.NewFromConfig(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbSearchOutput, err := client.DescribeDBInstances(ctx, &rds.DescribeDBInstancesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("engine"),
				Values: []string{"mysql"},
			},
		},
	})

	instanceInfoMap := make(map[string]string, 0)

	if err != nil {
		return instanceInfoMap, err
	}

	for _, instance := range dbSearchOutput.DBInstances {
		if utils.SafeDeref(instance.DBName, "") == dbName {
			instanceInfoMap["name"] = dbName
			instanceInfoMap["endpoint"] = utils.SafeDeref(instance.Endpoint.Address, "")
			instanceInfoMap["port"] = fmt.Sprintf("%d", utils.SafeDeref(instance.Endpoint.Port, 0))
			return instanceInfoMap, nil
		}
	}
	return instanceInfoMap, fmt.Errorf("map is empty")
}
