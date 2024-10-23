"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.addition = addition;
exports.sleepSync = sleepSync;
function addition(x, y) {
    return x + y;
}
function sleepSync(milliseconds) {
    const threshold = Date.now() + milliseconds;
    while (Date.now() < threshold) {
    }
}
//# sourceMappingURL=addition.js.map