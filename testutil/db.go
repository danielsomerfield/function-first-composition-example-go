package testutil

import (
	"context"
	"database/sql"
	"function-first-composition-example-go/review-server/db"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"strconv"
	"time"
)

type database struct {
	dbContainer *postgres.PostgresContainer
	db          *sql.DB
	datasource  *db.DataSource
}

func (d *database) Stop(ctx context.Context) error {
	return d.dbContainer.Terminate(ctx)
}

func (d *database) Exec(sql string, toBind ...interface{}) error {
	_, err := d.db.Exec(sql, toBind...)
	if err != nil {
		return err
	}
	return nil
}

func (d *database) DataSource() *db.DataSource {
	return d.datasource
}

func (d *database) GetDB() *sql.DB {
	return d.db
}

type DBImplementation interface {
	GetDB() *sql.DB
}

type DB interface {
	Stop(ctx context.Context) error
	Exec(sql string, toBind ...interface{}) error
	DataSource() *db.DataSource
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

	ports, err := dbContainer.Ports(context.Background())
	if err != nil {
		panic(err)
	}

	portBinding := ports["5432/tcp"][0]
	port, err := strconv.Atoi(portBinding.HostPort)
	if err != nil {
		panic(err)
	}

	instance := database{
		dbContainer: dbContainer,
		db:          sqlDB,
		datasource: &db.DataSource{
			Host:     "127.0.0.1",
			Port:     port,
			User:     dbUsername,
			Password: dbPassword,
			DbName:   databaseName,
		},
	}

	if err != nil {
		panic(err)
	}

	return &instance
}
