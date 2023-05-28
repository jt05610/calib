package gorm

import (
	"github.com/jt05610/calib"
	"gorm.io/gorm"
	"time"
)

type cal struct {
	gorm.Model
	xLabel string
	yLabel string
	points []*point
	result *regResult
}

func (c *cal) AddPoint(x, y float64) error {
	if c.result != nil {
		return calib.ErrCalibrationComplete
	}
	if c.points == nil {
		c.points = make([]*point, 0)
	}
	c.points = append(c.points, &point{x: x, y: y})
	return nil
}

func (c *cal) Ref() uint {
	return c.ID
}

func (c *cal) CreateTime() *time.Time {
	return &c.CreatedAt
}

func (c *cal) UpdateTime() *time.Time {
	return &c.UpdatedAt
}

func (c *cal) DeleteTime() *time.Time {
	return &c.DeletedAt.Time
}

func (c *cal) XLabel() string {
	return c.xLabel
}

func (c *cal) YLabel() string {
	return c.yLabel
}

func (c *cal) Points() []calib.Point {
	ret := make([]calib.Point, len(c.points))
	for i, p := range c.points {
		ret[i] = p
	}
	return ret
}

func (c *cal) Result() calib.RegResult {
	if c.result == nil {
		xs := make([]float64, len(c.points))
		ys := make([]float64, len(c.points))
		for i, p := range c.points {
			xs[i] = p.X()
			ys[i] = p.Y()
		}
		c.result = Regress(xs, ys).(*regResult)
	}
	return c.result
}

func Cal(xLabel, yLabel string) calib.Calibration {
	return &cal{
		xLabel: xLabel,
		yLabel: yLabel,
	}
}
