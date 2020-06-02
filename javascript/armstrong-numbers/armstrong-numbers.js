export const isArmstrongNumber = input => {
  return Array.from(input.toString()).map(Number).reduce((total, digit) => total + Math.pow(digit, input.toString().split('').length), 0) === input
};