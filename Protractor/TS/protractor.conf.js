exports.config = {
    directConnect: true,
    framework: 'jasmine',
    SELENIUM_PROMISE_MANAGER: false,
    specs: ['src/*.spec.ts'],
    beforeLaunch: function () {
        require('ts-node').register({
            project: 'tsconfig.json'
        });
    }
};
