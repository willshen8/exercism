export const steps = (num) => {
  if (num <= 0) {
    throw new Error("Only positive numbers are allowed");
  }
  let steps = 0;
  while (num != 1) {
    if (num % 2 == 0) {
      num = num / 2;
    } else {
      num = num * 3 + 1;
    }
    steps++;
  }
  return steps;
};
