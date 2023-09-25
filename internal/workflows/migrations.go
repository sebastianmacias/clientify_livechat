package workflows

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/workflow"
)

func RunMigrations(ctx workflow.Context, clusterID string, databaseID int) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: 60 * 60 * 24, // 24 hours
		StartToCloseTimeout:    60 * 60 * 24, // 24 hours
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var migrationResult string
	err := workflow.ExecuteActivity(ctx, RunMigrationsActivity, clusterID, databaseID).Get(ctx, &migrationResult)
	if err != nil {
		return "", err
	}

	var updateDBResult string
	err = workflow.ExecuteActivity(ctx, SetDatabasePendingMigrationsFalseActivity, clusterID, databaseID).Get(ctx, &updateDBResult)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Migration result: %s, Update DB result: %s", migrationResult, updateDBResult), nil
}

func RunMigrationsActivity(ctx context.Context, clusterID string, databaseID int) (string, error) {
	// Here you will put the real migration logic
	fmt.Printf("ctx.Err(): %v\n", ctx.Err())

	return fmt.Sprintf("Database with ID %d from cluster %s was migrated", databaseID, clusterID), nil
}

func SetDatabasePendingMigrationsFalseActivity(ctx context.Context, clusterID string, databaseID int) (string, error) {
	// Here you will put the real database update logic
	fmt.Printf("ctx.Err(): %v\n", ctx.Err())
	return fmt.Sprintf("Database with ID %d from cluster %s was updated", databaseID, clusterID), nil
}
