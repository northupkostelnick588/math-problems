function calculate(x, y) {
  if (x === 0 && y === 0) {
    throw new Error("Divide by zero error");
  }
  return x / y;
}
