package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)



var (
  dbDriver = "postgres"
  dbSource = "postgresql://root:MikeSera2022@localhost:5432/devcamper?sslmode=disable"
 TestQueries *Queries
)

func TestMain(m *testing.M){

  conn, err :=  sql.Open(dbDriver, dbSource) 
  if err != nil {
    log.Fatalf("could not connec to database %v", err)
  }

  TestQueries = New(conn)
  
  code := m.Run()
  _ = conn.Close()
  os.Exit(code)
}
