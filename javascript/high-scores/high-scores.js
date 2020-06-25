export class HighScores {
  constructor(input) {
    this._scores = input
  }

  get scores() {
    return this._scores
  }

  get latest() {
    return this._scores[this._scores.length - 1]
  }

  get personalBest() {
    if (!this._personalBest) {
      return this._scores.reduce((a, b) => Math.max(a, b))
    }
    return this._personalBest
  }

  get personalTopThree() {
    if (!this._personalTopThree) {
      return this._scores.sort((a, b) => b - a).slice(0, 3)
    }
    return !this._personalTopThree
  }
}