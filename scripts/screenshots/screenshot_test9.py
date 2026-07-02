import sys
import asyncio
from playwright.async_api import async_playwright

async def main():
    async with async_playwright() as p:
        browser = await p.chromium.launch()
        page = await browser.new_page()
        await page.goto('http://localhost:8080/atlas/docs/landing-demo.html')
        
        info = await page.evaluate('''() => {
            const tabs = document.querySelector('.tamarind-tabs');
            return tabs.outerHTML;
        }''')
        print(info)
        await browser.close()

asyncio.run(main())
