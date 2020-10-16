const ASCII_A = 65
const ASCII_Z = 90
const NUMBERS = 10
const TOTAL_NUMBER_OF_NAMES = 26 * 26 * 10 * 10 * 10;

export class Robot {
    constructor() {
        this.usedNames = new Set()
        this._name = this.generateRobotName()
    }

    get name() { return this._name }

    generateRobotName = () => {
        const alpha1 = String.fromCharCode(Math.floor(Math.random() * (ASCII_Z - ASCII_A) + ASCII_A))
        const alpha2 = String.fromCharCode(Math.floor(Math.random() * (ASCII_Z - ASCII_A) + ASCII_A))
        const num1 = Math.floor(Math.random() * NUMBERS)
        const num2 = Math.floor(Math.random() * NUMBERS)
        const num3 = Math.floor(Math.random() * NUMBERS)
        if (this.usedNames.has(`${alpha1}${alpha2}${num1}${num2}${num3}`)) {
            return this.generateRobotName()
        }
        this.usedNames.add(`${alpha1}${alpha2}${num1}${num2}${num3}`)
        return alpha1 + alpha2 + num1 + num2 + num3
    }

    reset = () => this._name = this.generateRobotName()
}

Robot.releaseNames = () => {};