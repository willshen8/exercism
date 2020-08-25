export class GradeSchool {
  constructor() {
    this.school = {}
  }

  roster() {
    const rosterCopy = JSON.parse(JSON.stringify(this.school))
    const sortObject = obj => Object.keys(obj).sort().reduce((res, key) => (res[key] = obj[key], res), {})
    return sortObject(rosterCopy)
  }

  add(name, grade) {
    if (this.school[grade]) {
      this.school[grade].push(name)
    } else {
      this.school[grade] = [name]
    }
  }

  grade(gradeNumber) {
    const rosterCopy = JSON.parse(JSON.stringify(this.school))
    if (rosterCopy[gradeNumber])
      return rosterCopy[gradeNumber].sort()
    else
      return []
  }
}