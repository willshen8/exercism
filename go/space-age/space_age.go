package space

type Planet string

const (
	EarthDays       = 365.25
	MercuryDays     = 0.2408467
	VenusDays       = 0.61519726
	MarsDays        = 1.8808158
	JupiterDays     = 11.862615
	SaturnDays      = 29.447498
	UranusDays      = 84.016846
	NeptuneDays     = 164.79132
	SecondsInOneDay = 86400
)

func Age(secs float64, p Planet) float64 {
	switch p {
	case "Earth":
		return secs / (EarthDays * SecondsInOneDay)
	case "Mercury":
		return secs / (MercuryDays * EarthDays * SecondsInOneDay)
	case "Venus":
		return secs / (VenusDays * EarthDays * SecondsInOneDay)
	case "Mars":
		return secs / (MarsDays * EarthDays * SecondsInOneDay)
	case "Jupiter":
		return secs / (JupiterDays * EarthDays * SecondsInOneDay)
	case "Saturn":
		return secs / (SaturnDays * EarthDays * SecondsInOneDay)
	case "Uranus":
		return secs / (UranusDays * EarthDays * SecondsInOneDay)
	case "Neptune":
		return secs / (NeptuneDays * EarthDays * SecondsInOneDay)
	default:
		return 0.0
	}
}
