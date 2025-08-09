package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()
	cfg := &config.Config{}
	config.LoadConfig(cfg)
	conn := &DBConnectionStruct{
		Conn: mockDB,
		Cfg:  cfg,
	}

	err = conn.OpenConnection()
	assert.NoError(t, err)
}

func TestCloseConnection(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	assert.NoError(t, err)
	cfg := &config.Config{}
	config.LoadConfig(cfg)
	conn := &DBConnectionStruct{
		Conn: mockDB,
		Cfg:  cfg,
	}

	conn.CloseConnection()

	assert.NoError(t, err)
}
