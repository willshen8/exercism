export const score = (x, y) => {
  const radius = Math.sqrt(x * x + y * y);
  if (radius <= 1) return 10;
  if (radius <= 5) return 5;
  if (radius <= 10) return 1;
  return 0;
};
