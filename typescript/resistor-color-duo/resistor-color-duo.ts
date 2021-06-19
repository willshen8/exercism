export class ResistorColor {
  constructor(private colors: Color[]) {
    if (colors.length < 2) {
      throw new Error('At least two colors need to be present')
    }
  }
  
  value():number {
    return resistorColors.indexOf(this.colors[0]) * 10 + resistorColors.indexOf(this.colors[1])
  }
  
}
export default ResistorColor

const resistorColors = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
] as const

export type Color = typeof resistorColors[number]