export class ResistorColor {
  private colors: string[]

  constructor(colors: string[]) {
    if (colors.length < 2) {
      throw new Error('At least two colors need to be present')
    }
    this.colors = colors
  }
  
  value():number {

    return parseInt(this.colors.slice(0,2).reduce((accumulator:string, color:string) => accumulator + resistorColors.get(color),''))
  }
  
}
export default ResistorColor


const resistorColors = new Map<string, string>([
  ['black', '0'],
  ['brown', '1'],
  ['red', '2'],
  ['orange', '3'],
  ['yellow', '4'],
  ['green', '5'],
  ['blue', '6'],
  ['violet', '7'],
  ['grey', '8'],
  ['white', '9'],
])
