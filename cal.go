package calib

type Calibration interface {
	Item
	AddPoint(x, y float64) error
	Regress() RegResult
}
