package gorm_test

import (
	"github.com/jt05610/calib"
	"github.com/jt05610/calib/gorm"
	"os"
	"testing"
)

func setUp() calib.ItemService[calib.Calibration] {
	err := os.Remove("test.db")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	return gorm.SqliteService("test.db")
}

func TestSqliteService(t *testing.T) {
	setUp()
}
