package dataproviders

import (
	"testing"
)

func TestNewDataProvider(t *testing.T) {
	conn := Connection{
		Host:       "localhost",
		Port:       5432,
		User:       "nikitagryshchak",
		Password:   "",
		DbName:     "todo",
		DisableSSL: true,
	}
	dataProvider, err := NewDataProvider(conn)
	if err != nil {
		t.Fatalf("error while creating db: %v", err)
	}
	if err := dataProvider.db.Ping(); err != nil {
		t.Fatalf("error while pinging db: %v", err)
	}
}
