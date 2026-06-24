import os
import glob

badge = """
    <!-- Ghost Badge -->
    <a href="https://usetamarind.com" target="_blank" class="tamarind-ghost-badge">
        Use Tamarind
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"></path><polyline points="15 3 21 3 21 9"></polyline><line x1="10" y1="14" x2="21" y2="3"></line></svg>
    </a>
</body>
"""

base_dir = "/home/rsantiago/Documents/atman-multi-agents/tamarind/parser/assets/templates"
for root, dirs, files in os.walk(base_dir):
    for f in ["page.mdt", "articles.mdt"]:
        if f in files:
            path = os.path.join(root, f)
            with open(path, "r") as file:
                content = file.read()
            if "tamarind-ghost-badge" not in content:
                content = content.replace("</body>", badge)
                with open(path, "w") as file:
                    file.write(content)
                print(f"Fixed {path}")
