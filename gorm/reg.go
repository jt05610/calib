package gorm

import (
	"github.com/jt05610/calib"
	"gonum.org/v1/gonum/stat"
	"gorm.io/gorm"
	"time"
)

type regResult struct {
	gorm.Model
	r2    float64
	inter float64
	slope float64
}

func (r *regResult) Ref() uint {
	return r.ID
}

func (r *regResult) CreateTime() *time.Time {
	return &r.CreatedAt
}

func (r *regResult) UpdateTime() *time.Time {
	return &r.UpdatedAt
}

func (r *regResult) DeleteTime() *time.Time {
	return &r.DeletedAt.Time
}

func (r *regResult) R2() float64 {
	return r.r2
}

func (r *regResult) Intercept() float64 {
	return r.inter
}

func (r *regResult) Slope() float64 {
	return r.slope
}

func Regress(xs, ys []float64) calib.RegResult {
	inter, slope := stat.LinearRegression(xs, ys, nil, false)
	return &regResult{
		r2:    stat.RSquared(xs, ys, nil, inter, slope),
		inter: inter,
		slope: slope,
	}
}
