package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Bound struct {
	min float64
	max float64
}

type SearchSpace struct {
	bounds []Bound
}

type Vector struct {
	values []float64
}

type Solution struct {
	vector Vector
	cost   float64
}

func objFunc(v Vector) float64 {
	sum := 0.0
	for _, x := range v.values {
		sum += x * x
	}
	return sum
}

func randomVec(space SearchSpace) Vector {
	rndVector := Vector{}
	dims := len(space.bounds)
	rndVector.values = make([]float64, dims)
	for i, rng := range space.bounds {
		rnd := rand.Float64()
		rndVector.values[i] = rng.min + ((rng.max - rng.min) *  rnd)
	}

	return rndVector
}

func rndSearch(space SearchSpace, maxIter int) Solution {
	best := Solution{
		cost: math.MaxFloat32,
	}
	for i := 0; i < maxIter; i++ {
		rndVec := randomVec(space)
		candidate := Solution{
			vector: rndVec,
			cost: objFunc(rndVec),
		}

		if candidate.cost < best.cost {
			best = candidate
		}
	}
	return best
}

func main() {
	searchSpace := SearchSpace{
		bounds: []Bound{
			Bound{min: -5, max: 5},
			Bound{min: -5, max: 5},
		},
	}

	best := rndSearch(searchSpace, 1000000)
	fmt.Println("Solution: {}", best)
}
