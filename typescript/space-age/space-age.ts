const OribitPeriodInEarthYear: { [key: string]: number } = {
    Mercury: 0.2408467,
    Venus: 0.61519726,
    Earth: 1,
    Mars: 1.8808158,
    Jupiter: 11.862615,
    Saturn: 29.447498,
    Uranus: 84.016846,
    Neptune: 164.79132
  }
  const yearInSeconds = 31_557_600

  interface SpaceAge {
    [key: string]: any
}

  class SpaceAge {
    constructor(public readonly seconds: number) {
        for (const planet in OribitPeriodInEarthYear) {
            this[`on${planet}`] = () => this.calculateSpaceAge(planet)
        }
    }

    calculateSpaceAge(planet:string):number {
        return Math.round(this.seconds / (OribitPeriodInEarthYear[planet] * yearInSeconds) * 100) / 100
    }
  }

  export default SpaceAge
