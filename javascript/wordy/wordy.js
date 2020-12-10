export const answer = question => {
  const validOperators = ["multiplied", "divided", "minus", "plus"]
  const isOperator = input => validOperators.includes(input)
  const isOperand = input => Number.isInteger(parseInt(input))
  
  const calculate = (operandA, operator, operandB) => {
    if (!(isOperand(operandA) && isOperator(operator) && isOperand(operandB))) {
      throw new Error('Syntax error')
    }
    switch(operator) {
      case "multiplied":
        return operandA * operandB
      case "divided":
        return operandA / operandB
      case "minus":
        return operandA - operandB
      case "plus":
        return operandA + operandB
    }
  }

  const parsedCommands = []
  const parsedQuestion = question.replace('?', '').split(" ")
  parsedQuestion.map((word, index) => {
    if (index === 0 && word !== "What") {
      throw new Error('Unknown operation')
    }
    if(isOperand(word) || isOperator(word)) {
      parsedCommands.push(word)
    }
  })
  if (parsedCommands.length === 1 && parsedQuestion.length>3) throw new Error('Unknown operation')
  if (parsedCommands.length === 1 && isOperand(parsedCommands[0])) return parseInt(parsedCommands[0])
  let result = calculate(parseInt(parsedCommands[0]), parsedCommands[1], parseInt(parsedCommands[2]))
  for (let i=3; i<parsedCommands.length; i+=2){
    result = calculate(result, parsedCommands[i], parseInt(parsedCommands[i+1]))
  }
  return result
};
