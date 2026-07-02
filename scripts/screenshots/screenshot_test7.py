import sys
import asyncio
from playwright.async_api import async_playwright

async def main():
    async with async_playwright() as p:
        browser = await p.chromium.launch()
        page = await browser.new_page()
        page.on("console", lambda msg: print(f"Browser console: {msg.text}"))
        page.on("pageerror", lambda err: print(f"Browser error: {err}"))
        
        await page.goto('http://localhost:8080/atlas/docs/landing-demo.html')
        
        # Click the Linux tab of the first terminal
        await page.click("button:has-text('Linux')")
        
        info = await page.evaluate('''() => {
            const activePanes = document.querySelectorAll('.terminal-tab-pane.active');
            return Array.from(activePanes).map(p => p.id);
        }''')
        print(f"Active panes: {info}")
        await browser.close()

asyncio.run(main())
