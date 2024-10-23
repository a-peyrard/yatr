"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const addition_1 = require("./addition");
describe('addition with 1', () => {
    it('should validate addition with 2', () => {
        const a = 1;
        const b = 2;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(3);
        (0, addition_1.sleepSync)(10);
    });
    it('should validate addition with 3', () => {
        const a = 1;
        const b = 3;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(4);
        (0, addition_1.sleepSync)(10);
    });
    it('should validate addition with 4', () => {
        const a = 1;
        const b = 4;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(5);
        (0, addition_1.sleepSync)(10);
    });
    it('should validate addition with 5', () => {
        const a = 1;
        const b = 5;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(6);
        (0, addition_1.sleepSync)(10);
    });
    it('should validate addition with 6', () => {
        const a = 1;
        const b = 6;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(7);
        (0, addition_1.sleepSync)(10);
    });
});
describe('addition with 2', () => {
    it('should validate addition with 2', () => {
        const a = 2;
        const b = 2;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(4);
        (0, addition_1.sleepSync)(10);
    });
    it('should validate addition with 3', () => {
        const a = 2;
        const b = 3;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(5);
        (0, addition_1.sleepSync)(10);
    });
    it('should validate addition with 4', () => {
        const a = 2;
        const b = 4;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(6);
        (0, addition_1.sleepSync)(10);
    });
    it('should validate addition with 5', () => {
        const a = 2;
        const b = 5;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(7);
        (0, addition_1.sleepSync)(10);
    });
    it('should validate addition with 6', () => {
        const a = 2;
        const b = 6;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(8);
        (0, addition_1.sleepSync)(10);
    });
});
//# sourceMappingURL=addition21.spec.js.map