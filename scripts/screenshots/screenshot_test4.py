import sys
import asyncio
from playwright.async_api import async_playwright

async def main():
    async with async_playwright() as p:
        browser = await p.chromium.launch()
        page = await browser.new_page()
        await page.goto('http://localhost:8080/editorial/docs/style-guide.html')
        
        info = await page.evaluate('''() => {
            const cards = document.querySelectorAll('.price-card');
            return Array.from(cards).map(c => {
                const badge = c.querySelector('.card-badge-poc');
                return {
                    height: c.clientHeight,
                    badgeTop: badge ? badge.offsetTop : null,
                    badgeText: badge ? badge.innerText : null
                };
            });
        }''')
        print(info)
        await browser.close()

asyncio.run(main())
