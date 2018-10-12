![](https://upload.wikimedia.org/wikipedia/commons/thumb/6/6e/Levy0_distributionPDF.svg/325px-Levy0_distributionPDF.svg.png)

# go-levy
go-levy is a golang port of Mantegna's Levy Stable Stochastic Process Generation Algorithm, introduced in the 1994 edition of [Physical Review E](https://journals.aps.org/pre/).


I decided to place it here after finding an interest in a specific metaheuristic (Cuckoo Search AKA obligate brood parasitism) as a form of Relevance Feedback in Information Retrieval.


The function needs as input the distribution’s parameters α ∈[0.3,1.99] and c > 0, the number of iterations n and the number of random points to be produced; without this last input, the output consists of a single number. If an input parameter is outside the valid range, an error message is displayed (It produces NaN, and I'd rather make use of an error as a value instead of NaN).



P.S. the sample usage is found in gl.go
