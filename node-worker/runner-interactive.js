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

        // fixme
        //sleepSync(Math.floor(Math.random() * 100) + 1);  // Simulate a slow test case
        // fixme
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

console.log('Almost ready to accept spec files...');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false
});

rl.on('line', (filePath) => {
    // Run the spec file received via stdin
    console.log(`Running test file: ${filePath}`);
    require(filePath);

    console.log(`Finished running: ${filePath}`);
});

// fixme...
function sleepSync(milliseconds) {
    const start = Date.now();
    while (Date.now() - start < milliseconds) {
        // Busy-wait: Do nothing
    }
}