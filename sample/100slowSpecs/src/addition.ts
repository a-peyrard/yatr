export function addition(x: number, y: number): number {
  return x + y;
}

// dummy sleep function as we can not use async...
export function sleepSync(milliseconds: number) {
  const threshold = Date.now() + milliseconds;
  while (Date.now() < threshold) {
    // Busy-wait: Do nothing
  }
}
