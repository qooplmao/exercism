package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	converted := ""

	if 0 == number % 3 {
		converted = converted + "Pling"
	}

	if 0 == number % 5 {
		converted = converted + "Plang"
	}

	if 0 == number % 7 {
		converted = converted + "Plong"
	}

	if "" == converted {
		converted = strconv.Itoa(number)
	}

	return converted
}
