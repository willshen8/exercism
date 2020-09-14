export class GradeSchool {
  constructor() {
    this.school = {}
  }

  roster() {
    return JSON.parse(JSON.stringify(this.school))
  }

  add(name, grade) {
    if (this.school[grade]) {
      this.school[grade].push(name)
    } else {
      this.school[grade] = [name]
    }
    this.school[grade].sort()
  }

  grade(gradeNumber) {
    const rosterCopy = JSON.parse(JSON.stringify(this.school))
    if (rosterCopy[gradeNumber])
      return rosterCopy[gradeNumber]
    else
      return []
  }
}