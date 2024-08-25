//
// This is only a SKELETON file for the 'Flatten Array' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const flatten = (input) => {
  let result = [];
  for (let i = 0; i < input.length; i++) {
    if (Array.isArray(input[i])) {
      result = result.concat(flatten(input[i]));
    } else {
      if (input[i] !== null && input[i] !== undefined) {
        result.push(input[i]);
      }
    }
  }
  return result;
};
