package gorm

import (
	"github.com/jt05610/calib"
	"gorm.io/gorm"
	"time"
)

type Cal struct {
	gorm.Model
	XLabel string
	YLabel string
	Points []*Point
	Result *RegResult
}

func (c *Cal) AddPoint(x, y float64) error {
	if c.Result != nil {
		return calib.ErrCalibrationComplete
	}
	if c.Points == nil {
		c.Points = make([]*Point, 0)
	}
	c.Points = append(c.Points, &Point{X: x, Y: y})
	return nil
}

func (c *Cal) Ref() uint {
	return c.ID
}

func (c *Cal) CreateTime() *time.Time {
	return &c.CreatedAt
}

func (c *Cal) UpdateTime() *time.Time {
	return &c.UpdatedAt
}

func (c *Cal) DeleteTime() *time.Time {
	return &c.DeletedAt.Time
}

func (c *Cal) Regress() calib.RegResult {
	if c.Result == nil {
		xs := make([]float64, len(c.Points))
		ys := make([]float64, len(c.Points))
		for i, p := range c.Points {
			xs[i] = p.X
			ys[i] = p.Y
		}
		c.Result = Regress(xs, ys).(*RegResult)
	}
	return c.Result
}

func NewCal(xLabel, yLabel string) calib.Calibration {
	return &Cal{
		XLabel: xLabel,
		YLabel: yLabel,
	}
}
