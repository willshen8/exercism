export const hey = (message) => {
  let trimmedMessage = message.trim()
  if (trimmedMessage.replace(/[^a-zA-Z0-9!?]/g, "") == "") return "Fine. Be that way!"
  if (/^[A-Z]+$/.test(trimmedMessage.replace(/[^a-zA-Z0-9]/g, "")) && trimmedMessage[trimmedMessage.length - 1] == '?') {
    return "Calm down, I know what I'm doing!"
  }
  if (trimmedMessage[trimmedMessage.length - 1] == '?') return "Sure."
  if (/^[A-Z]+$/.test(trimmedMessage.replace(/[^a-zA-Z]/g, ""))) return "Whoa, chill out!"
  return "Whatever."
}