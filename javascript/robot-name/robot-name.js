import { stringify } from "querystring"

const ASCII_A = 65
const ASCII_Z = 90
const MAX_NUM = 1000
const generateRandomLetter = () => String.fromCharCode(Math.floor(Math.random() * (ASCII_Z - ASCII_A) + ASCII_A))
const generateRandomNumber = () => Math.floor(Math.random() * MAX_NUM)
const usedNames = new Set()

export class Robot {
    constructor() {
        this._name = this.generateUniqueRobotName()
    }

    get name() { return this._name }

    generateUniqueRobotName = () => {
        let newRobotName = this.NewRobotName()
        while (usedNames.has(newRobotName)) {
            newRobotName = this.NewRobotName()
        }
        usedNames.add(newRobotName)
        return newRobotName
    }

    NewRobotName = () => {
        const randomLetter1 = generateRandomLetter()
        const randomLetter2 = generateRandomLetter()
        const randomNum = generateRandomNumber().toString().padStart(3, "0")
        return `${randomLetter1}${randomLetter2}${randomNum}`
    }

    reset = () => this._name = this.generateUniqueRobotName()
}
Robot.releaseNames = () => {};