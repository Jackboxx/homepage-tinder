const puppeteer = require('puppeteer');

const getScreenshot = async (url, name) => {
    const browser = await puppeteer.launch({ 
        headless: false,
        defaultViewport: {width: 1920, height: 1080},
        args: [
            '--disable-extensions-except=src/cli/IDCAC-chrome',
            '--load-extension=src/cli/IDCAC-chrome',
        ]
    });

    const page = await browser.newPage();
    await page.goto(url,  {waitUntil: 'domcontentloaded'});



    await page.waitForNavigation({
        waitUntil: 'networkidle0',
        timeout: 5000
    }).catch(err => {});

    await page.screenshot({ path: `web/images/${name}.png`, type: 'png' });

    await page.close();
    await browser.close();
};

const url = process.argv[2];
const domain = process.argv[3];

getScreenshot(url, domain);
