package gorm_test

import (
	"github.com/jt05610/calib/gorm"
	"math/rand"
	"testing"
)

func checkTolerance(t *testing.T, expected, actual, tolerance float64) {
	if actual < expected-tolerance || actual > expected+tolerance {
		t.Errorf("expected %f, got %f", expected, actual)
	}
}

func TestCal(t *testing.T) {
	cal := gorm.Cal("steps", "masses")
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

	result := cal.Result()
	if result == nil {
		t.Error("nil result")
	}
	checkTolerance(t, slope, result.Slope(), 0.5)
	checkTolerance(t, 0, result.Intercept(), 3)
	checkTolerance(t, 1, result.R2(), tol*2)
}
