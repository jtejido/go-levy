package main

import (
	"fmt"
	"github.com/jtejido/go-levy/levy"
)

func main() {
	alphas := []float64{0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 1.95, 1.99}
	gamma := 1.0
	n := 1
	l := levy.NewLevy()

	fmt.Printf("This displays Table I of Mantegna's paper....\n")
	fmt.Printf("alpha     levy function     sigmax         K          C\n")
	for _, alpha := range alphas {
		z, err := l.Levy(alpha, gamma, n)
		if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%.5f    %.5f          %.5f     %.5f     %.5f\n", alpha, z, l.Sigmax(alpha), l.K(alpha), l.C(alpha))
			}
		
	}
}
