import { addition } from './addition';

describe('addition1', () => {
  it('should validate addition', () => {
    // GIVEN
    const a = 1;
    const b = 2;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(3);
  });
});
