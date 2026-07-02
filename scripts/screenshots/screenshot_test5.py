import sys
import asyncio
from playwright.async_api import async_playwright

async def main():
    async with async_playwright() as p:
        browser = await p.chromium.launch()
        page = await browser.new_page()
        await page.goto('http://localhost:8080/editorial/docs/style-guide.html')
        
        info = await page.evaluate('''() => {
            const cards = document.querySelectorAll('.card');
            return Array.from(cards).map(c => {
                const div = c.querySelector('div');
                const badge = div ? div.querySelector('span') : null;
                return {
                    id: c.querySelector('h3')?.id,
                    divLeft: div ? div.offsetLeft : null,
                    badgeLeft: badge ? badge.offsetLeft : null,
                    pLeft: c.querySelector('p') ? c.querySelector('p').offsetLeft : null,
                    h3Left: c.querySelector('h3') ? c.querySelector('h3').offsetLeft : null
                };
            });
        }''')
        print(info)
        await browser.close()

asyncio.run(main())
