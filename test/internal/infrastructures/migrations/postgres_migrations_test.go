package postgres_migrations

import (
	"context"
	"fmt"
	"testing"

	infrastructures_drivers_postgres "meul/inventory/internal/infrastructures/drivers/postgres"
	migrations "meul/inventory/internal/infrastructures/drivers/postgres/migrations"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func Test_whenRunUpMigrations_thenApplyAllMigrationsFilesInOrderASC(t *testing.T) {
	ctx := context.Background()
	dBName := "testdb"
	user := "testuser"
	password := "testpassword"

	ctr, err := postgres.Run(
		ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dBName),
		postgres.WithUsername(user),
		postgres.WithPassword(password),
		postgres.BasicWaitStrategies(),
	)

	require.NoError(t, err)
	defer testcontainers.CleanupContainer(t, ctr)

	mappedPort, err := ctr.MappedPort(ctx, "5432")
	require.NoError(t, err)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"localhost", user, password, dBName, mappedPort.Port(), "disable",
	)
	dbConfig := infrastructures_drivers_postgres.DbConfig{
		DSN: dsn,
	}

	require.NoError(t, err)

	migrationConfig := migrations.MigrationConfig{

		MigrationPath: migrations.MigrationPath("file://../../../../internal/infrastructures/drivers/postgres/migrations/sql"),
	}

	migrationHandler := migrations.MigrationHandler{
		DbConfig:        dbConfig,
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
