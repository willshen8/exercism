//  Helper functions
const isShouting = (message: string) => message.toUpperCase() === message && /[A-Z]/.test(message)
const isSilence = (message: string) => message === ""
const isAsking = (message: string) => message[message.length - 1] === '?'

class Bob {
  hey(comment: string): string {
    return SITUATIONS.find(situation => situation.test(comment.trim()))!.answer;
  }
}

// List of situations
const SITUATIONS: Array<{ answer: string, test: (s: string) => boolean }> = [{
  answer: "Fine. Be that way!",
  test: isSilence
},
{
  answer: "Calm down, I know what I'm doing!",
  test: (message: string) => isAsking(message) && isShouting(message)
},
{
  answer: "Sure.",
  test: isAsking
},
{
  answer: "Whoa, chill out!",
  test: isShouting
},
{
  answer: "Whatever.",
  test: () => true
}
]

export default Bob

