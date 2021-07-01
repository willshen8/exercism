const NUM_OF_ENGLISH_ALPHABET_LETTERS:number = 26

class Pangram {
    constructor(public readonly sentence: string) {
    }

    isPangram():boolean {
        const cleanedSentence = [...this.sentence.replace(/[^A-Za-z]/g, '')]
        const wordMap = new Map()

        cleanedSentence.forEach(letter => {
            if (!wordMap.has(letter)) {
                wordMap.set(letter, 1)
            }
        })
        return wordMap.size === NUM_OF_ENGLISH_ALPHABET_LETTERS
    }
}

export default Pangram