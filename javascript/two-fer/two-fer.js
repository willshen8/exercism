export const twoFer = (...args) => {
  const [name] = args
  switch (name) {
    case undefined:
      return "One for you, one for me."
    default:
      return "One for " + name + ", one for me."
  }
};
