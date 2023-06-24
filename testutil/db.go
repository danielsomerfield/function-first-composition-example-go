package testutil

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"
)

type db struct {
	dbContainer *postgres.PostgresContainer
	db          *sql.DB
}

func (d *db) Stop(ctx context.Context) error {
	return d.dbContainer.Terminate(ctx)
}

func (d *db) Exec(sql string, toBind ...interface{}) error {
	_, err := d.db.Exec(sql, toBind...)
	if err != nil {
		return err
	}
	return nil
}

func (d *db) GetDB() *sql.DB {
	return d.db
}

type DBImplementation interface {
	GetDB() *sql.DB
}

type DB interface {
	Stop(ctx context.Context) error
	Exec(sql string, toBind ...interface{}) error
}

func StartDB(ctx context.Context, initScript string) DB {
	databaseName := "postgres"
	dbUsername := "postgres"
	dbPassword := "postgres"
	dbContainer, err := postgres.RunContainer(
		ctx,
		postgres.WithDatabase(databaseName),
		postgres.WithUsername(dbUsername),
		postgres.WithPassword(dbPassword),
		postgres.WithInitScripts(initScript),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").
			WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		panic(err)
	}

	connectionString, err := dbContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		panic(err)
	}

	sqlDB, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	instance := db{
		dbContainer: dbContainer,
		db:          sqlDB,
	}

	if err != nil {
		panic(err)
	}

	return &instance
}
