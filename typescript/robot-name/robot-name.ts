const ASCII_A = "A".charCodeAt(0)
const ASCII_Z = "Z".charCodeAt(0)
const MAX_NUM = 1000
const generateRandomLetter = (): string => String.fromCharCode(Math.floor(Math.random() * (ASCII_Z - ASCII_A + 1) + ASCII_A))
const generateRandomNumber = (): number => Math.floor(Math.random() * MAX_NUM)
const usedNames = new Set();
export default class Robot {
  private _name: string

  constructor() {
    this._name = this.generateUniqueRobotName()
  }

  private generateUniqueRobotName(): string {
    let newRobotName = this.NewRobotName()
    do {
      newRobotName = this.NewRobotName()
    } while (usedNames.has(newRobotName))
    usedNames.add(newRobotName)
    return newRobotName
  }

  private NewRobotName(): string {
    const randomLetter1: string = generateRandomLetter()
    const randomLetter2: string = generateRandomLetter()
    const randomNum: string = generateRandomNumber().toString().padStart(3, "0")
    return `${randomLetter1}${randomLetter2}${randomNum}`
  }

  public get name(): string {
    return this._name
  }

  public resetName(): void {
    this._name = this.generateUniqueRobotName()
  }

  public static releaseNames(): void {
    usedNames.clear()
  }
}
