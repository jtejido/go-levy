package levy

import (
    "math"
    "fmt"
)

// Fast, accurate algorithm for numerical simulation of Levy stable stochastic processes 
// Mantegna, R, N. 1994
type Levy struct {}

func NewLevy() *Levy {
	return new(Levy)
}

// Stochastic variable
func (levy Levy) Vf(alpha float64) float64 {
	var x, y float64
	x = randNormal(0, 1)
	y = randNormal(0, 1)

	x = x * levy.Sigmax(alpha)

	return x / math.Pow(math.Abs(y), 1.0 / alpha)
}

func (levy Levy) Sigmax(alpha float64) float64 {
	numerator := math.Gamma(alpha + 1.0) * math.Sin(math.Pi * alpha / 2.0)
	denominator := math.Gamma((alpha + 1)/2.0) * alpha * math.Pow(2.0, (alpha - 1.0) / 2.0)

	return math.Pow(numerator / denominator, 1.0 / alpha)
}

func (levy Levy) K(alpha float64) float64 {
	k := alpha * math.Gamma((alpha + 1.0)/(2.0 * alpha))/ math.Gamma(1.0 / alpha)
	k *= math.Pow(alpha * math.Gamma((alpha + 1.0)/2.0) / (math.Gamma(alpha + 1.0) * math.Sin(math.Pi * alpha / 2.0)), 1.0 / alpha)

	return k
}

func (levy Levy) C(alpha float64) float64 {
    x := []float64{0.75, 0.8, 0.9, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 1.95, 1.99}
    y := []float64{2.2085, 2.483, 2.7675, 2.945, 2.941, 2.9005, 2.8315, 2.737, 2.6125, 2.4465, 2.206, 1.7915, 1.3925, 0.6089}
	li := NewLinear()
	li.Fit(x, y)
	
	estimate, err := interpolate(li, alpha)
	
    if err != nil {
		panic(err)
	}
	
	return estimate
}

func (levy Levy) Levy(alpha, gamma float64, n int) (float64, error) {
    var v, w, z float64

    if gamma <= 0 {
        return z, fmt.Errorf("gamma out of range %f", gamma) 
    }

    if n < 0 {
        return z, fmt.Errorf("iteration less than zero %f", n) 
    }

	if alpha >= 0.3 && alpha <= 1.99 {

            w = 0
            for i := 0; i <= n; i++ {
                v = levy.Vf(alpha)

                for v < -10 {
                        v = levy.Vf(alpha)
                }
                    w += v * ((levy.K(alpha) - 1.0) * math.Exp(-v / levy.C(alpha)) + 1.0)
            }
            // The Levy random variable
            z = 1.0 / math.Pow(float64(n), 1.0 / alpha) * w * gamma

        return z, nil
    }

    return z, fmt.Errorf("alpha out of range %f", alpha)
}