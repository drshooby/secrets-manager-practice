package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/drshooby/secrets-manager-practice/utils"
)

func getSecret(cfg aws.Config, secretName string) string {
	client := secretsmanager.NewFromConfig(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	secretOutput, err := client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	})
	if err != nil {
		return ""
	}

	return utils.SafeDeref(secretOutput.SecretString, "")
}
