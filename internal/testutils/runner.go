package testutils

import (
	"context"
	"testing"
)

// Run runs the tests and returns an exit code to pass to os.Exit.
func Run(m *testing.M) int {
	close := InitRDB()
	defer close()

	ctx := context.Background()

	createTablesIfNotExist(ctx)
	truncateTables(ctx)

	return m.Run()
}
