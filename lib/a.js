var PhantomCSS = require('../node_modules/phantomcss/phantomcss.js');

PhantomCSS.init({
    'screenshotRoot': './screenshots',
    'libraryRoot': 'node_modules/phantomcss'
});

casper.start(casper.cli.get('url'))
    .viewport(1440, 900)
    .then(function () {
        PhantomCSS.screenshot(casper.cli.get('selector'), 
                              casper.cli.get('selector').replace(/^[^a-z]/, ''));
    })
    .run(function() {
        phantom.exit(PhantomCSS.getExitStatus());
    });