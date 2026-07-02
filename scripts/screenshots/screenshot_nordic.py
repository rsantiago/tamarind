import sys
import asyncio
from playwright.async_api import async_playwright

async def main():
    async with async_playwright() as p:
        browser = await p.chromium.launch()
        page = await browser.new_page()
        await page.goto('http://localhost:8080/nordic/docs/landing-demo.html')
        
        # Take screenshot of the tabs component
        element = await page.query_selector('.tamarind-tabs')
        if element:
            await element.screenshot(path='/home/rsantiago/.gemini/antigravity-cli/brain/d8aedfe9-c417-4b8e-ae3f-f74752eecf49/scratch/nordic_tabs.png')
            print("Screenshot saved.")
        else:
            print("Element not found.")
        await browser.close()

asyncio.run(main())
