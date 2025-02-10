package database

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
  _"github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	_ "github.com/joho/godotenv/autoload"
)




func mustStartPostgresContainer() (func(context.Context, ...testcontainers.TerminateOption) error, error) {
	var (
		dbName = "database"
		dbPwd  = "password"
		dbUser = "user"
	)

	dbContainer, err := postgres.Run(
		context.Background(),
		"postgres:17-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPwd),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}
	os.Setenv("BLUEPRINT_DB_DATABASE", dbName)
	os.Setenv("BLUEPRINT_DB_PASSWORD", dbPwd)
	os.Setenv("BLUEPRINT_DB_USERNAME", dbUser)

	dbHost, err := dbContainer.Host(context.Background())
  log.Println("dbHost: ",dbHost)
	if err != nil {
		return dbContainer.Terminate, err
	}

	dbPort, err := dbContainer.MappedPort(context.Background(), "5432/tcp")
  log.Println("port: ", dbPort)
	if err != nil {
		return dbContainer.Terminate, err
	}

	os.Setenv("BLUEPRINT_DB_HOST", dbHost)
	os.Setenv("BLUEPRINT_DB_PORT", dbPort.Port())
  log.Printf("dbName: %v, dbUser: %v, dbPwd: %v, dbHost: %v, dbPort: %v", dbName, dbUser, dbPwd, dbHost, dbPort)

	return dbContainer.Terminate, err
}

func TestMain(m *testing.M) {
	teardown, err := mustStartPostgresContainer()
	if err != nil {
		log.Fatalf("could not start postgres container: %v", err)
	}
  m.Run()

	if teardown != nil && teardown(context.Background()) != nil {
		log.Fatalf("could not teardown postgres container: %v", err)
	}
}

func TestNew(t *testing.T) {
	srv := New()
  log.Println("service",srv)
	if srv == nil {
		t.Fatalf("New() returned nil")
	}
}

func TestHealth(t *testing.T) {
	srv := New()

	stats := srv.Health()

	if stats["status"] != "up" {
		t.Fatalf("expected status to be up, got %s", stats["status"])
	}

	if _, ok := stats["error"]; ok {
		t.Fatalf("expected error not to be present")
	}

	if stats["message"] != "It's healthy" {
		t.Fatalf("expected message to be 'It's healthy', got %s", stats["message"])
	}
}


