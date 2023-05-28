package calib

type Calibration interface {
	Item
	XLabel() string
	YLabel() string
	Points() []Point
	AddPoint(x, y float64) error
	Result() RegResult
}
