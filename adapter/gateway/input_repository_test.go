package gateway

import (
	"fmt"
	"testing"
	"time"

	"github.com/ShotaKitazawa/pi-temperature-api/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestInputGateway(t *testing.T) {

	t.Run("Get1()", func(t *testing.T) {
		t.Run("inputsテーブルから最近のデータ一つを得る", func(t *testing.T) {
			r := createRepositoryForTest(t, "../../testdb/inputs-test.sqlite3")

			expected := &domain.Input{
				ID:          1,
				Temperature: 10.2,
				Humidity:    20.4,
				CreatedAt:   time.Date(2019, time.April, 11, 22, 22, 22, 432590000, time.Local),
			}
			actual, err := r.Get1()
			if err != nil {
				t.Fatal("failed getting to db")
			}

			assert.Equal(t, expected.ID, actual.ID)
			assert.Equal(t, expected.Temperature, actual.Temperature)
			assert.Equal(t, expected.Humidity, actual.Humidity)
			assert.Equal(t, expected.CreatedAt.Year(), actual.CreatedAt.Year())
			assert.Equal(t, expected.CreatedAt.Month(), actual.CreatedAt.Month())
			assert.Equal(t, expected.CreatedAt.Day(), actual.CreatedAt.Day())
			assert.Equal(t, expected.CreatedAt.Hour(), actual.CreatedAt.Hour())
			assert.Equal(t, expected.CreatedAt.Minute(), actual.CreatedAt.Minute())
			assert.Equal(t, expected.CreatedAt.Second(), actual.CreatedAt.Second())
		})
		t.Run("inputsテーブルが空の場合エラーを返す", func(t *testing.T) {
			r := createRepositoryForTest(t, "../../testdb/empty-test.sqlite3")

			if _, err := r.Get1(); err == nil {
				t.Fatal("table entry in None, but not return nil")
			}

		})
	})

}

func createRepositoryForTest(t *testing.T, path string) *InputRepository {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		fmt.Println(err)
		t.Fatal("failed connecting to db")
	}
	return &InputRepository{DBConn: db}
}
