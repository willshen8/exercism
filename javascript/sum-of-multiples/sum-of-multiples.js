export const sum = (baseValues, levelCompleted) => {
  if (levelCompleted <= 1 || baseValues.length === 0) return 0;

  let uniqueMultiplesSet = new Set();
  baseValues.forEach((baseValue) => {
    if (baseValue === 0) return;

    for (
      let multiple = baseValue;
      multiple < levelCompleted;
      multiple += baseValue
    ) {
      uniqueMultiplesSet.add(multiple);
    }
  });

  return sumOfSet(uniqueMultiplesSet);
};

const sumOfSet = (set) => [...set].reduce((acc, curr) => acc + curr);
