import os, re

pal_light = ["#2563eb", "#10b981", "#d946ef", "#f59e0b", "#7c3aed", "#0ea5e9", "#e11d48", "#84cc16", "#475569"]
pal_dark  = ["#3b82f6", "#6ee7b7", "#c084fc", "#fde047", "#f472b6", "#67e8f9", "#fb7185", "#bef264", "#94a3b8"]

def process_css(css):
    css = re.sub(r'\s*--chart-\d+:[^;]+;', '', css)
    
    def light_repl(m):
        insert = "\n"
        for i, c in enumerate(pal_light):
            insert += f"    --chart-{i+1}: {c};\n"
        return insert + m.group(1)
        
    css = re.sub(r'(\n\})', light_repl, css, count=1)
    
    if '[data-theme="dark"]' in css:
        match = re.search(r'\[data-theme="dark"\]\s*\{', css)
        if match:
            start = match.end()
            brace_count = 1
            idx = start
            while idx < len(css) and brace_count > 0:
                if css[idx] == '{': brace_count += 1
                elif css[idx] == '}': brace_count -= 1
                idx += 1
            if brace_count == 0:
                css = css[:idx-1] + "\n" + "".join([f"    --chart-{i+1}: {c};\n" for i, c in enumerate(pal_dark)]) + css[idx-1:]
    return css

templates_dir = "/home/rsantiago/Documents/atman-multi-agents/tamarind/parser/assets/templates"
for theme in os.listdir(templates_dir):
    theme_dir = os.path.join(templates_dir, theme)
    if os.path.isdir(theme_dir):
        css_file = os.path.join(theme_dir, "style.css")
        if os.path.exists(css_file):
            with open(css_file, "r") as f:
                css = f.read()
            new_css = process_css(css)
            with open(css_file, "w") as f:
                f.write(new_css)
            print(f"Processed {theme}")

