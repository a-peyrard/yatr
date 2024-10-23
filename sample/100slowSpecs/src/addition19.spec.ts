import {addition, sleepSync} from './addition';

describe('addition with 1', () => {
  it('should validate addition with 2', () => {
    // GIVEN
    const a = 1;
    const b = 2;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(3);
    sleepSync(10);
  });

  it('should validate addition with 3', () => {
    // GIVEN
    const a = 1;
    const b = 3;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(4);
    sleepSync(10);
  });

  it('should validate addition with 4', () => {
    // GIVEN
    const a = 1;
    const b = 4;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(5);
    sleepSync(10);
  });

  it('should validate addition with 5', () => {
    // GIVEN
    const a = 1;
    const b = 5;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(6);
    sleepSync(10);
  });

  it('should validate addition with 6', () => {
    // GIVEN
    const a = 1;
    const b = 6;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(7);
    sleepSync(10);
  });
});

describe('addition with 2', () => {
  it('should validate addition with 2', () => {
    // GIVEN
    const a = 2;
    const b = 2;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(4);
    sleepSync(10);
  });

  it('should validate addition with 3', () => {
    // GIVEN
    const a = 2;
    const b = 3;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(5);
    sleepSync(10);
  });

  it('should validate addition with 4', () => {
    // GIVEN
    const a = 2;
    const b = 4;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(6);
    sleepSync(10);
  });

  it('should validate addition with 5', () => {
    // GIVEN
    const a = 2;
    const b = 5;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(7);
    sleepSync(10);
  });

  it('should validate addition with 6', () => {
    // GIVEN
    const a = 2;
    const b = 6;

    // WHEN
    const result = addition(a, b);

    // THEN
    expect(result).toBe(8);
    sleepSync(10);
  });
});
