export const COLORS = ["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"]
export const colorCode = color => COLORS.indexOf(color)
export const decodedValue = colors => colorCode(colors[0]) * 10 + colorCode(colors[1])