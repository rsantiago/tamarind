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
            return Array.from(cards).map(c => ({
                id: c.querySelector('h3')?.id,
                cardHeight: c.clientHeight,
                cardDisplay: window.getComputedStyle(c).display,
                cardFlexDirection: window.getComputedStyle(c).flexDirection,
                divMarginTop: window.getComputedStyle(c.querySelector('div')).marginTop,
                divTop: c.querySelector('div').offsetTop
            }));
        }''')
        print(info)
        await browser.close()

asyncio.run(main())
