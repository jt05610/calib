package gorm_test

import (
	"github.com/jt05610/calib"
	"github.com/jt05610/calib/gorm"
	"math/rand"
	"os"
	"sync"
	"testing"
)

func setUp() calib.ItemService[*gorm.Cal] {
	err := os.Remove("test.db")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	return gorm.SqliteService("test.db")
}

func TestSqliteService(t *testing.T) {
	var wg sync.WaitGroup
	s := setUp()
	nEntries := 100
	nCals := 10
	for i := 0; i < nCals; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 1. create and save new calibration
			cal := gorm.NewCal("steps", "masses").(*gorm.Cal)
			err := s.Add(cal)
			if err != nil {
				t.Error(err)
			}
			// 2. add points to calibration
			slope := 5.0
			tol := 0.20
			for i := 0; i < nEntries; i++ {
				v := float64(i) * slope
				for j := 0; j < 5; j++ {
					y := rand.NormFloat64()*(v*tol) + v
					err := cal.AddPoint(float64(i), y)
					if err != nil {
						t.Error(err)
					}
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
		}()
	}
	wg.Wait()
	// 4. load the most recent regression
	allCals, err := s.GetAll() // assuming cal.ID is updated during s.Add(cal)
	if err != nil {
		t.Error(err)
	}
	if len(allCals) < nCals {
		t.Error("not enough cals")
	}
}
