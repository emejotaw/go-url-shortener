package sqlite

import (
	"testing"

	"github.com/emejotaw/go-url-shortener/config"
	"github.com/emejotaw/go-url-shortener/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type sqliteRepositoryTest struct {
	db *gorm.DB
}

var (
	srt = &sqliteRepositoryTest{
		db: config.GetConnection(),
	}
)

func TestNewSqliteRepository(t *testing.T) {

	config.AutoMigrate()
	repository := NewSqliteRepository(srt.db)

	assert.NotNil(t, repository)
}

func TestCreate(t *testing.T) {

	type testCase struct {
		longURL     string
		shortURL    string
		expectedErr error
	}
	repository := NewSqliteRepository(srt.db)
	tests := []testCase{
		{longURL: "https://google.com?query='Some test query'", shortURL: "abc", expectedErr: nil},
		{longURL: "https://globoesporte.com", shortURL: "xyz", expectedErr: nil},
		{longURL: "https://youtube.com", shortURL: "xxx", expectedErr: nil},
	}

	for _, test := range tests {

		url := entity.New(test.longURL, test.shortURL)
		err := repository.Create(url)
		assert.Nil(t, err)
	}
}

func TestFindByUrlKey(t *testing.T) {

	type testCase struct {
		longURL     string
		urlKey      string
		expectedErr error
	}
	repository := NewSqliteRepository(srt.db)
	tests := []testCase{
		{longURL: "https://google.com?query='Some test query'", urlKey: "abc", expectedErr: nil},
		{longURL: "https://globoesporte.com", urlKey: "xyz", expectedErr: nil},
		{longURL: "https://youtube.com", urlKey: "xxx", expectedErr: nil},
	}

	for _, test := range tests {

		url, _ := repository.FindByUrlKey(test.urlKey)
		assert.NotNil(t, url)
	}

}

func TestFindByID(t *testing.T) {

	type testCase struct {
		id int
	}

	tests := []testCase{
		{id: 1},
		{id: 2},
		{id: 3},
		{id: 4},
	}

	repository := NewSqliteRepository(srt.db)

	for _, test := range tests {

		_, err := repository.FindByID(test.id)

		assert.Nil(t, err)
	}
}

func TestUpdate(t *testing.T) {

	type testCase struct {
		id          int
		longURL     string
		urlKey      string
		expectedErr error
	}
	tests := []testCase{
		{id: 1, longURL: "https://somecrazyurl.com.br'", urlKey: "abc", expectedErr: nil},
		{id: 2, longURL: "https://longlonglonglonglonglongurl.com", urlKey: "xyz", expectedErr: nil},
		{id: 3, longURL: "https://anotherlongurl.com", urlKey: "xxx", expectedErr: nil},
	}

	repository := NewSqliteRepository(srt.db)

	for _, test := range tests {

		url, err := repository.FindByID(test.id)

		if err != nil {
			t.Fatalf("error when executing tests, err: %v", err.Error())
		}

		url.LongUrl = test.longURL
		url.UrlKey = test.urlKey
		err = repository.Update(url)

		assert.Nil(t, err)
	}
}

func TestDeleteByID(t *testing.T) {

	type testCase struct {
		id int
	}

	tests := []testCase{
		{id: 1},
		{id: 2},
		{id: 3},
	}

	repository := NewSqliteRepository(srt.db)

	for _, test := range tests {

		err := repository.DeleteByID(test.id)

		assert.Nil(t, err)
	}
}
