"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const addition_1 = require("./addition");
describe('addition1', () => {
    it('should validate addition', () => {
        const a = 1;
        const b = 2;
        const result = (0, addition_1.addition)(a, b);
        expect(result).toBe(3);
    });
});
//# sourceMappingURL=addition38.spec.js.map