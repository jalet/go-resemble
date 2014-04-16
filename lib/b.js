var PhantomCSS = require('../node_modules/phantomcss/phantomcss.js');

PhantomCSS.init({
    'screenshotRoot': './screenshots',
    'libraryRoot': 'node_modules/phantomcss'
});

casper.start()
    .then(function () {
        PhantomCSS.compareAll();
    })
    .run(function() {
        phantom.exit(PhantomCSS.getExitStatus());

    });