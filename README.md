![](https://www.rspb.org.uk/globalassets/images/birds-and-wildlife/bird-species-illustrations/cuckoo_grey_female_1200x675.jpg?preset=landscape_mobile)

# go-levy
go-levy is a golang port of a Stochastic Process Generation Algorithm introduced by Rosario Nunzio Mantegna in 1994 edition of [Physical Review E](https://journals.aps.org/pre/).


I decided to place it here after finding an interest in a specific metaheuristic (Cuckoo Search AKA obligate brood parasitism) as a form of Relevance Feedback in Information Retrieval.

The function needs as input the distribution’s parameters α ∈[0.3,1.99] and c > 0, the number of iterations n and the number of random points to be produced; without this last input, the output consists of a single number. If an input parameter is outside the valid range, an error message is displayed (It produces NaN, and I'd rather make use of an error as a value instead of NaN).


Grab it by typing:

```go get github.com/jtejido/go-levy```


P.S. the implementation needed to run it is in gl.go
