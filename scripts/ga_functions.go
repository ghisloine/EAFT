package scripts

import (
	"math/rand"
)

func MutNormalFloat64(genome []float64, rate float64, rng *rand.Rand) {
	for i := range genome {
		// Flip a coin and decide to mutate or not
		if rng.Float64() < rate {
			//genome[i] += rng.NormFloat64() * genome[i]
			if genome[i] < 0.5 {
				genome[i] = 0
			} else {
				genome[i] = 1
			}
		}
	}
}

// Returns a Binary Vector with just only 0 - 1.
func InitBinaryFloat64(n uint, lower, upper float64, rng *rand.Rand) (floats []float64) {
	floats = make([]float64, n)
	for i := range floats {
		floats[i] = lower + float64(rng.Intn(int(upper)))
		// fmt.Printf("Generated Number is %f\n", floats[i])
	}
	return
}

func ExecuteCode(cmd string) {
	//output := exec.Command(cmd)
	// TODO : Bu kisim'a geri kalan kodun execute edilmesi ile ilgili seyler yazilacak.
}
