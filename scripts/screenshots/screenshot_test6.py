import sys
import asyncio
from playwright.async_api import async_playwright

async def main():
    async with async_playwright() as p:
        browser = await p.chromium.launch()
        page = await browser.new_page()
        await page.goto('http://localhost:8080/atlas/docs/landing-demo.html')
        
        info = await page.evaluate('''() => {
            const terminals = document.querySelectorAll('.terminal');
            return Array.from(terminals).map(t => {
                const buttons = Array.from(t.querySelectorAll('.terminal-tab-btn'));
                const panes = Array.from(t.querySelectorAll('.terminal-tab-pane'));
                return {
                    isWrappedInPre: t.parentElement.tagName === 'CODE' && t.parentElement.parentElement.tagName === 'PRE',
                    buttons: buttons.map(b => b.outerHTML),
                    panes: panes.map(p => p.id)
                };
            });
        }''')
        print(info)
        await browser.close()

asyncio.run(main())
