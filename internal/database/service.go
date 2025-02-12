package database

import (
	"database/sql"
	"fmt"
	"log"
)


type Service interface{
  Health() map[string]string
  Close() error
}


type service struct{
  db *sql.DB
}

func New(dataSourceName string) (Service, error){
  db, err := sql.Open("postgres", dataSourceName)
  if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

  if err := db.Ping(); err != nil {
    db.Close()
    return nil,fmt.Errorf("failed to connect to database: %w", err)

  }
  log.Println("Database connection established")
  return &service{db: db}, nil
}


func (s *service) Health() map[string]string {
  err := s.db.Ping()
  if err != nil {
    return map[string]string{"status": "down", "error": err.Error()}
  }
  return map[string]string{"status": "up"}
}

func (s *service) Close() error {
  
  return s.db.Close()
}
