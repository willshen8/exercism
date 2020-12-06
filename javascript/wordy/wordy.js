export const answer = question => {
  const validOperators = ["multiplied", "divided", "minus", "plus"]
  const isOperator = input => validOperators.includes(input)
  const isOperand = input => Number.isInteger(parseInt(input))
  const parsedOperands = []
  const parsedOperators = []
  const calculate = (operandA, operandB, operation) => {
    console.log("Hello world: ", operandA, operandB)
    switch(operation) {
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

  question.replace('?', '').split(" ").map(word => {
    if(isOperand(word)) {
      parsedOperands.push(parseInt(word))
    }
    if (isOperator(word)) {
      parsedOperators.push(word)
    }
  })

  if (parsedOperands.length === 1) return parsedOperands[0]
  return calculate(parsedOperands[0], parsedOperands[1], parsedOperators[0])
};
