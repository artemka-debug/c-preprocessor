#!/usr/bin/env node

const testsRunner = require('./run.js');
const clc = require('cli-color');

(() => {
    console.log(clc.green('Welcome C Preprocessor tests!'));
    if (process.argv.length <= 2) {
        return console.error(clc.yellow('Usage: ./cli.js <path-to-preprocessor-executable>'));
    }
    testsRunner(process.argv[2]);
})();
