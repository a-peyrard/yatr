const { expect } = require('expect');

const readline = require('readline');

// Register global test functions
global.describe = function (name, fn) {
    console.log(`Running test suite: ${name}`);
    fn();  // This will run the test cases defined in this `describe` block
};

global.it = function (name, fn) {
    try {
        console.log(`Running test case: ${name}`);
        fn();  // Run the actual test function
        console.log(`Test case passed: ${name}`);
    } catch (error) {
        console.error(`Test case failed: ${name}`, error);
    }
};

global.beforeEach = function (fn) {
    console.log('Running beforeEach setup');
    fn();  // Run the setup function
};

global.afterEach = function (fn) {
    console.log('Running afterEach teardown');
    fn();  // Run the teardown function
};

global.expect = expect;

const path = process.argv[2];  // Pass the test file as an argument
// require('ts-node').register(); // For TypeScript, otherwise skip this line
require(path);
