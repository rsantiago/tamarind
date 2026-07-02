const puppeteer = require('puppeteer');

(async () => {
    console.log("Launching browser...");
    const browser = await puppeteer.launch({
        args: ['--no-sandbox', '--disable-setuid-sandbox', '--window-size=1200,800'],
        defaultViewport: { width: 1200, height: 800 }
    });
    const page = await browser.newPage();
    console.log("Navigating to http://localhost:8080/gram/docs/shortcode-charts.html...");
    await page.goto('http://localhost:8080/gram/docs/shortcode-charts.html', { waitUntil: 'networkidle2' });

    console.log("Taking screenshot of initial state...");
    await page.screenshot({ path: 'initial.png', fullPage: true });

    // Click maximize on Multiline chart
    console.log("Maximizing Multiline Chart...");
    await page.evaluate(() => {
        const btns = Array.from(document.querySelectorAll('.chart-maximize-btn'));
        const btn = btns.find(b => b.closest('.tamarind-multiline'));
        if (btn) btn.click();
    });
    await new Promise(r => setTimeout(r, 1000));
    await page.screenshot({ path: 'multiline_maximized.png' });

    // Click close on Multiline chart
    await page.evaluate(() => {
        const btns = Array.from(document.querySelectorAll('.chart-maximize-btn'));
        const btn = btns.find(b => b.closest('.tamarind-multiline'));
        if (btn) btn.click();
    });
    await new Promise(r => setTimeout(r, 500));

    // Click maximize on Grouped Bar Chart
    console.log("Maximizing Grouped Bar Chart...");
    await page.evaluate(() => {
        const btns = Array.from(document.querySelectorAll('.chart-maximize-btn'));
        const btn = btns.find(b => b.closest('.tamarind-groupedbar'));
        if (btn) btn.click();
    });
    await new Promise(r => setTimeout(r, 1000));
    await page.screenshot({ path: 'groupedbar_maximized.png' });

    // Click close
    await page.evaluate(() => {
        const btns = Array.from(document.querySelectorAll('.chart-maximize-btn'));
        const btn = btns.find(b => b.closest('.tamarind-groupedbar'));
        if (btn) btn.click();
    });
    await new Promise(r => setTimeout(r, 500));

    // Click maximize on Horizontal Bar Chart
    console.log("Maximizing Horizontal Bar Chart...");
    await page.evaluate(() => {
        const btns = Array.from(document.querySelectorAll('.chart-maximize-btn'));
        const btn = btns.find(b => b.closest('.tamarind-hbarchart'));
        if (btn) btn.click();
    });
    await new Promise(r => setTimeout(r, 1000));
    await page.screenshot({ path: 'hbar_maximized.png' });

    // Click close
    await page.evaluate(() => {
        const btns = Array.from(document.querySelectorAll('.chart-maximize-btn'));
        const btn = btns.find(b => b.closest('.tamarind-hbarchart'));
        if (btn) btn.click();
    });
    await new Promise(r => setTimeout(r, 500));
    
    // Click maximize on Pie Chart
    console.log("Maximizing Pie Chart...");
    await page.evaluate(() => {
        const btns = Array.from(document.querySelectorAll('.chart-maximize-btn'));
        const btn = btns.find(b => b.closest('.tamarind-piechart'));
        if (btn) btn.click();
    });
    await new Promise(r => setTimeout(r, 1000));
    await page.screenshot({ path: 'pie_maximized.png' });

    console.log("Closing browser...");
    await browser.close();
})();
