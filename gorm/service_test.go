package gorm_test

import (
	"github.com/jt05610/calib"
	"github.com/jt05610/calib/gorm"
	"math/rand"
	"os"
	"testing"
	"time"
)

func setUp() calib.ItemService[*gorm.Cal] {
	err := os.Remove("test.db")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	return gorm.SqliteService("test.db")
}

func TestSqliteService(t *testing.T) {
	s := setUp()

	// 1. create and save new calibration
	cal := gorm.NewCal("steps", "masses").(*gorm.Cal)
	err := s.Add(cal)
	if err != nil {
		t.Error(err)
	}
	// 2. add points to calibration
	slope := 5.0
	tol := 0.05
	for i := 0; i < 100; i++ {
		v := float64(i) * slope
		y := rand.NormFloat64()*(v*tol) + v
		err := cal.AddPoint(float64(i), y)
		if err != nil {
			t.Error(err)
		}
	}
	// update calibration with the added points
	err = s.Update(cal)
	if err != nil {
		t.Error(err)
	}
	// 3. run the regression and check the result
	result := cal.Regress().(*gorm.RegResult)
	err = s.Update(cal)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Error("nil result")
	}
	checkTolerance(t, slope, result.Slope, 0.5)
	checkTolerance(t, 0, result.Intercept, 3)
	checkTolerance(t, 1, result.R2, tol*2)

	// 4. load the most recent regression
	latestCal, err := s.Get(cal.ID) // assuming cal.ID is updated during s.Add(cal)
	if err != nil {
		t.Error(err)
	}
	latestResult := latestCal.Result
	// 5. check if the most recent regression has an r2 > acceptance value and a date that is within a set time window
	acceptanceValue := 0.95
	if latestResult.R2 < acceptanceValue {
		t.Errorf("R2 is below the acceptance value: got %v, want > %v", latestResult.R2, acceptanceValue)
	}
	if time.Now().Sub(latestResult.CreatedAt) > 24*time.Hour {
		t.Error("The regression is more than 24 hours old")
	}
}
