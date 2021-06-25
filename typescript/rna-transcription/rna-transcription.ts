const DNAtoRNA : { [key: string]: string } = {
  'G':'C',
  'C':'G',
  'T':'A',
  'A':'U',
} as const
class Transcriptor {
  toRna(dnaStrand:string): string {
    return dnaStrand.split('').map((item: string) => {
      if (!DNAtoRNA[item]) throw new Error('Invalid input DNA.')
      return DNAtoRNA[item]
    }).join("")
  }
}



export default Transcriptor
