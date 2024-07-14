package search

import "math"

/**
 * Given two crystal balls that will break
 * if dropped from high enough distance,
 * determine the exact spot in which
 * it will break in the most optimized way.
 */

func TwoCrystalBalls(levels []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(levels)))))

	i := jumpAmount
	for ; i < len(levels); i += jumpAmount {
		if levels[i] {
			break
		}
	}

	i -= jumpAmount
	for j := 0; j <= jumpAmount && i < len(levels); j, i = j+1, i+1 {
		if levels[i] {
			return i
		}
	}

	return -1
}
