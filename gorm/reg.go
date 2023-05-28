package gorm

import (
	"github.com/jt05610/calib"
	"gonum.org/v1/gonum/stat"
	"gorm.io/gorm"
	"time"
)

type RegResult struct {
	gorm.Model
	R2        float64
	Intercept float64
	Slope     float64
	CalID     uint
}

func (r *RegResult) Ref() uint {
	return r.ID
}

func (r *RegResult) CreateTime() *time.Time {
	return &r.CreatedAt
}

func (r *RegResult) UpdateTime() *time.Time {
	return &r.UpdatedAt
}

func (r *RegResult) DeleteTime() *time.Time {
	return &r.DeletedAt.Time
}

func Regress(xs, ys []float64) calib.RegResult {
	inter, slope := stat.LinearRegression(xs, ys, nil, false)
	return &RegResult{
		R2:        stat.RSquared(xs, ys, nil, inter, slope),
		Intercept: inter,
		Slope:     slope,
	}
}
