package hamming

import (
	"errors"
)

// Calculate number of different characters between 2 strings
func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
    	return 0, errors.New("both strings do not have the same length")
    }

	distance := 0

    for i := 0; i < len(a); i++ {
    	if a[i:i + 1] != b[i:i + 1] {
    		distance += 1
    	}
    }

    return distance, nil
}
