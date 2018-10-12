package levy

import (
    "math"
    "math/rand"
    "sort"
    "fmt"
)

// Gamma Function via Lanczos approximation formula. Depracated in favor of math.Gamma (Stirling approximation)
func gamma(x float64) float64 {
	return math.Exp(logGamma(x))
}

func logGamma(x float64) float64 {
	tmp := (x - 0.5) * math.Log(x + 4.5) - (x + 4.5)
      	ser := 1.0 + 76.18009172947146 / (x + 0) - 86.50532032941677 / (x + 1) + 24.01409824083091 / (x + 2) -  1.231739572450155 / (x + 3) +  0.001208650973866179 / (x + 4) -  0.000005395239384953 / (x + 5)
      	return tmp + math.Log(ser * math.Sqrt(2 * math.Pi))
}

func randNormal(m, s float64) float64 {
	return rand.NormFloat64() * s + m
}

func interpolate(l *Linear, val float64) (float64, error) {
	var est float64

	err := l.validate(val)
	if err != nil {
		return est, err
	}

	est = l.interpolate(val)
	return est, nil
}

type Linear struct {
	BaseInterpolation
}

func NewLinear() *Linear {
	li := &Linear{}
	return li
}

func (li *Linear) interpolate(val float64) float64 {
	var est float64

	l, r := li.searchNearestNeighbours(val, 0, len(li.XY)-1)

	lX := li.XY[l].X
	rX := li.XY[r].X
	lY := li.XY[l].Y
	rY := li.XY[r].Y

	est = lY + (rY-lY)/(rX-lX)*(val-lX)
	return est
}

func (li *Linear) validate(val float64) error {

	if val < li.XY[0].X {
		return fmt.Errorf("Out of bounds: %f less than %f", val, li.XY[0].X)
	}

	if val > li.XY[len(li.XY)-1].X {
		return fmt.Errorf("Out of bounds: %f greater than %f", val, li.XY[len(li.XY)-1].X)
	}

	return nil
}

func (li *Linear) searchNearestNeighbours(val float64, l, r int) (int, int) {
	middle := (l + r) / 2
	if (val >= li.XY[middle-1].X) && (val <= li.XY[middle].X) {
		return middle - 1, middle
	} else if val < li.XY[middle-1].X {
		return li.searchNearestNeighbours(val, l, middle-2)
	}
	return li.searchNearestNeighbours(val, middle+1, r)
}

type BaseInterpolation struct {
	XY []CoordinatePair
	X       []float64
	Y       []float64
}

func (b *BaseInterpolation) Fit(x, y []float64) error {
	if len(x) != len(y) {
		return fmt.Errorf("Pairs do not match X: %f Y: %f", x, y)
	}
	b.X = x
	b.Y = y
	b.XY = sliceToPairs(x, y)
	sortPairs(b.XY)
	return nil
}

type CoordinatePair struct {
	X float64
	Y float64
}

func sortPairs(cp []CoordinatePair) {
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].X < cp[j].X
	})
}

func sliceToPairs(x, y []float64) []CoordinatePair {
	cp := make([]CoordinatePair, len(x))
	for i := 0; i < len(x); i++ {
		cp = append(cp, CoordinatePair{X: x[i], Y: y[i]})
	}
	return cp
}
