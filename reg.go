package calib

type RegResult interface {
	Item
	R2() float64
	Intercept() float64
	Slope() float64
}
