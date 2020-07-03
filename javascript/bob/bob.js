export const hey = (message) => {
  let trimmedMessage = message.trim()
  return SITUATIONS
    .find(situation => situation.test(trimmedMessage))
    .answer;
}

//  Helper functions
const isShouting = message => /^[A-Z]+$/.test(message.replace(/[^a-zA-Z]/g, ""))
const isSilence = message => message.replace(/[^a-zA-Z0-9!?]/g, "") === ""
const isAsking = message => message[message.length - 1] === '?'

// List of situations
const SITUATIONS = [{
    answer: "Fine. Be that way!",
    test: isSilence
  },
  {
    answer: "Calm down, I know what I'm doing!",
    test: function (message) {
      return isAsking(message) && isShouting(message)
    }
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
    test: function () {
      return true
    }
  }
]