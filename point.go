package calib

type Point interface {
	Item
	X() float64
	Y() float64
}
