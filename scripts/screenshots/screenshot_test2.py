import sys
import asyncio
from playwright.async_api import async_playwright

async def main():
    async with async_playwright() as p:
        browser = await p.chromium.launch()
        page = await browser.new_page()
        await page.goto('http://localhost:8080/editorial/docs/style-guide.html')
        
        # Get the innerHTML of the second card
        html = await page.evaluate('''() => {
            const cards = document.querySelectorAll('.card');
            return cards[1].innerHTML;
        }''')
        print(html)
        await browser.close()

asyncio.run(main())
