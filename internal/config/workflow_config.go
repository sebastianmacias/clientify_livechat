package config

import (
	"fmt"
)

type Workflow struct {
	Host string
	Port int
}

// ConnectionString generates a connection string to be passed to workflow client
func (w Workflow) ConnectionString() string {
	return fmt.Sprintf("%s:%d", w.Host, w.Port)
}
