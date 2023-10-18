package workflows

import (
	"fmt"

	"go.temporal.io/sdk/client"
)

type Service struct {
	Client client.Client
}

func NewService(hostPort string) (*Service, error) {

	c, err := client.Dial(client.Options{
		HostPort: hostPort,
	})

	if err != nil {
		return nil, fmt.Errorf("unable to create Temporal client: %w", err)
	}
	return &Service{Client: c}, nil
}

func (t *Service) StartMigrationWorkflow(clusterID string, databaseID int) (string, error) {
	fmt.Printf("clusterID: %v\n", clusterID)
	fmt.Printf("databaseID: %v\n", databaseID)
	// use t.Client to start the workflow
	return "", nil
}

// ... other methods to interact with Temporal ...

func (t *Service) Close() {
	t.Client.Close()
}
