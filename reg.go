package calib

type RegResult interface {
	Item
	R2() float64
	Intercept() float64
	Slope() float64
}

func LinReg(points []Point) RegResult {
	xs := make([]float64, len(points))
	ys := make([]float64, len(points))
	for i, p := range points {
		xs[i] = p.X()
		ys[i] = p.Y()
	}

}
