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
    if (this.school[gradeNumber])
      return JSON.parse(JSON.stringify(this.school[gradeNumber]))
    else
      return []
  }
}