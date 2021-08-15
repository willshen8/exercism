const ASCII_A = 65
const ASCII_Z = 90
const MAX_NUM = 1000
const generateRandomLetter = ():string => String.fromCharCode(Math.floor(Math.random() * (ASCII_Z - ASCII_A) + ASCII_A))
const generateRandomNumber = ():number => Math.floor(Math.random() * MAX_NUM)
const usedNames = new Set();
export default class Robot {
  _name: string
  
  constructor() {
    this._name = this.generateUniqueRobotName()
  }

  generateUniqueRobotName = (): string => {
    let newRobotName = this.NewRobotName()
    do {
      newRobotName = this.NewRobotName()
    } while (usedNames.has(newRobotName))
    usedNames.add(newRobotName)
    return newRobotName
  }

  NewRobotName = ():string => {
    let randomLetter1 = generateRandomLetter()
    let randomLetter2 = generateRandomLetter()
    let randomNum = generateRandomNumber().toString().padStart(3, "0")
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
