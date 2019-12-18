package space

const SECONDS_IN_EARTH_YEAR float64 = 31557600

type Planet string

func Age(seconds float64, planet Planet) float64 {
	return seconds / (SECONDS_IN_EARTH_YEAR * numberOfEarthYearsForOneYearOnPlanet(planet))
}

func numberOfEarthYearsForOneYearOnPlanet(planet Planet) float64 {
	switch planet {
	case "Mercury":
		return 0.2408467
	case "Venus":
		return 0.61519726
	case "Earth":
		return 1
	case "Mars":
		return 1.8808158
	case "Jupiter":
		return 11.862615
	case "Saturn":
		return 29.447498
	case "Uranus":
		return 84.016846
	case "Neptune":
		return 164.79132
	default:
		return 1
	}
}
