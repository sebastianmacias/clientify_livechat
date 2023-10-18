package cmd

import (
	"fmt"
	"log"

	// "github.com/sebastianmacias/clientify_livechat/internal/workflows"

	"github.com/spf13/cobra"
	// "go.temporal.io/sdk/client"
	// "go.temporal.io/sdk/worker"
)

var queueName string

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Starts the temporal worker",
	Long:  `Starts the temporal worker with specified queue name and worker count.`,
	Run: func(cmd *cobra.Command, args []string) {

		validQueueNames := []string{"pending_migrations"}
		isValidQueueName := false
		for _, validQueueName := range validQueueNames {
			fmt.Println(validQueueName + "-" + queueName)

			if queueName == validQueueName {
				isValidQueueName = true
				break
			}
		}
		if !isValidQueueName {
			log.Fatal("invalid queue_name. available options are: 'pending_migrations'")
		}

		if queueName == "" {
			log.Fatal("queue_name must be provided")
		}

		// serviceConfig := config.DefaultServiceConfigFromEnv()

		// c, err := client.Dial(client.Options{
		// 	HostPort: serviceConfig.Workflow.ConnectionString(),
		// })

		// if err != nil {
		// 	log.Fatalf("unable to create Temporal client: %v", err)
		// }

		// defer c.Close()

		// w := worker.New(c, queueName, worker.Options{})

		// if queueName == "pending_migrations" {
		// 	w.RegisterWorkflow(workflows.RunMigrations)
		// 	w.RegisterActivity(workflows.RunMigrationsActivity)
		// 	w.RegisterActivity(workflows.SetDatabasePendingMigrationsFalseActivity)
		// }

		// Register other workflows and activities based on queueName here

		// err = w.Run(worker.InterruptCh())
		// if err != nil {
		// 	log.Fatalf("unable to start worker: %v", err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
	workerCmd.Flags().StringVarP(&queueName, "queue_name", "q", "", "The name of the queue that the worker will listen on")
}
