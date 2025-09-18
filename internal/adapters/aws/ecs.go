package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ECS struct {
	client *ecs.Client
}

func New(cfg aws.Config) *ECS {
	return &ECS{client: ecs.NewFromConfig(cfg)}
}

func (e *ECS) RunTasks(ctx context.Context, cluster, taskDef string) (string, error) {
	out, err := e.client.RunTask(ctx, &ecs.RunTaskInput{
		Cluster:        &cluster,
		TaskDefinition: &taskDef,
	})
	if err != nil {
		return "", fmt.Errorf("		failed to run task: %w", err)
	}
	if len(out.Tasks) == 0 {
		return "", fmt.Errorf("no task started")
	}

	return *out.Tasks[0].TaskArn, nil
}
