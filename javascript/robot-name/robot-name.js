const ASCII_A = 65
const ASCII_Z = 90
const DIGITS = 10
const generateRandomLetter = () => String.fromCharCode(Math.floor(Math.random() * (ASCII_Z - ASCII_A) + ASCII_A))
const generateRandomDigit = () => Math.floor(Math.random() * DIGITS)

export class Robot {
    constructor() {
        this.usedNames = new Set()
        this._name = this.generateRobotName()
    }

    get name() { return this._name }

    generateRobotName = () => {
        const alpha1 = generateRandomLetter()
        const alpha2 = generateRandomLetter()
        const num1 = generateRandomDigit()
        const num2 = generateRandomDigit()
        const num3 = generateRandomDigit()
        const newRobotName = `${alpha1}${alpha2}${num1}${num2}${num3}`
        if (this.usedNames.has(newRobotName)) {
            return this.generateRobotName()
        }
        this.usedNames.add(newRobotName)
        return newRobotName
    }

    reset = () => this._name = this.generateRobotName()
}

Robot.releaseNames = () => {};