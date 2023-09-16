package sqlite

import (
	"testing"

	"github.com/emejotaw/go-url-shortener/config"
	"github.com/stretchr/testify/assert"
)

func TestNewSqliteRepository(t *testing.T) {

	db := config.GetConnection()
	repository := NewSqliteRepository(db)

	assert.NotNil(t, repository)
}
