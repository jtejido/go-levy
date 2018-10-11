package main

import (
	"fmt"
	"github.com/jtejido/go-levy"
)

func main() {
	alphas := []float64{0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 1.95, 1.99}
	gamma := 1.0
	n := 1
	l := levy.NewLevy()

	fmt.Printf("This displays Table I of Mantegna's paper (except the function result) ....\n")
	fmt.Printf("alpha     levy function     sigmax         K          C\n")
	for _, alpha := range alphas {
		fmt.Printf("%.5f    %.5f          %.5f     %.5f     %.5f\n", alpha, l.Levy(alpha, gamma, n), l.Sigmax(alpha), l.K(alpha), l.C(alpha))
	}
}
