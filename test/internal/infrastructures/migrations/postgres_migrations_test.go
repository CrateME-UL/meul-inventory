package postgres_migrations

import (
	"context"
	"fmt"
	"testing"

	migrations "meul/inventory/internal/infrastructures/drivers/postgres/migrations"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func Test_whenRunUpMigrations_thenApplyAllMigrationsFilesInOrderASC(t *testing.T) {
	ctx := context.Background()
	dbname := "testdb"
	user := "testuser"
	password := "testpassword"

	ctr, err := postgres.Run(
		ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbname),
		postgres.WithUsername(user),
		postgres.WithPassword(password),
		postgres.BasicWaitStrategies(),
	)

	require.NoError(t, err)
	defer testcontainers.CleanupContainer(t, ctr)

	dbURL, err := ctr.ConnectionString(ctx)

	require.NoError(t, err)

	migrationConfig := migrations.MigrationConfig{
		DatabaseURL:   migrations.DatabaseURL(fmt.Sprintf("%s?sslmode=disable", dbURL)),
		MigrationPath: migrations.MigrationPath("file://../../../../internal/infrastructures/drivers/postgres/migrations/"),
	}

	migrationHandler := migrations.MigrationHandler{
		MigrationConfig: &migrationConfig,
	}

	err = migrationHandler.RunUp()

	require.NoError(t, err)

	err = ctr.Snapshot(ctx, postgres.WithSnapshotName("initial-state"))
	require.NoError(t, err)

	t.Run("whenMigrationRunDown_thenRollbackAllMigrations", func(t *testing.T) {
		t.Cleanup(func() {
			err := ctr.Restore(ctx)
			require.NoError(t, err)
		})

		err = migrationHandler.RunDown()

		require.NoError(t, err)
	})
}
