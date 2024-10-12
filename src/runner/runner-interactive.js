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

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false
});

let number_of_tests = 0;
rl.on('line', (filePath) => {
    if (number_of_tests === 0) {
        console.time('100 tests');
    }
    // Run the spec file received via stdin
    console.log(`Running test file: ${filePath}`);
    require(filePath);

    console.log(`Finished running: ${filePath}`);

    number_of_tests++;
    if (number_of_tests === 100) {
        console.timeEnd('100 tests');
        process.exit(0);
    }
});
