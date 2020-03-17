const clc = require('cli-color');
const fs = require('fs');

function run(executable) {
    console.log(clc.green('Running tests for ', executable));
    const testsRoot = './samples';
    const testsModules = fs.readdirSync(testsRoot);
    console.log(testsModules);
    for (const testModule of testsModules) {
        const testModuleRoot = testsRoot + '/' + testModule;
        const testCasesFiles = fs.readdirSync(testModuleRoot);
        const testCases = testCasesFiles.filter(file => /[a-z\-_]+.c/i);
        const expectedOutPut = testCasesFiles.filter(file => /([a-z0-9\-_]+)(_expected)(.c)/i.test(file));
        console.log(testCases);

    }
}

module.exports = run;
