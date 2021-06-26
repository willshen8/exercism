const OribitPeriodInEarthYear: { [key: string]: number } = {
    mercury: 0.2408467,
    venus: 0.61519726,
    earth: 1,
    mars: 1.8808158,
    jupiter: 11.862615,
    saturn: 29.447498,
    uranus: 84.016846,
    neptune: 164.79132
  }
  const yearInSeconds = 31557600

  class SpaceAge {
    public seconds: number;
    constructor(seconds: number) {
        this.seconds = seconds
    }

    calculateSpaceAge(planet:string):number {
        return Math.round(this.seconds / (OribitPeriodInEarthYear[planet] * yearInSeconds) * 100) / 100
    }

    onEarth() {
        return this.calculateSpaceAge('earth')
    }

    onMercury(){
        return this.calculateSpaceAge('mercury')
    }

    onVenus(){
        return this.calculateSpaceAge('venus')
    }

    onMars(){
        return this.calculateSpaceAge('mars')
    }

    onJupiter(){
        return this.calculateSpaceAge('jupiter')
    }

    onSaturn(){
        return this.calculateSpaceAge('saturn')
    }

    onUranus(){
        return this.calculateSpaceAge('uranus')
    }

    onNeptune(){
        return this.calculateSpaceAge('neptune')
    }
  }

  export default SpaceAge
