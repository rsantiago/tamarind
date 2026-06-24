import os, re

def relative_luminance(r, g, b):
    def adjust(c):
        c = c / 255.0
        return c / 12.92 if c <= 0.03928 else ((c + 0.055) / 1.055) ** 2.4
    return 0.2126 * adjust(r) + 0.7152 * adjust(g) + 0.0722 * adjust(b)

def contrast_ratio(l1, l2):
    if l1 < l2: l1, l2 = l2, l1
    return (l1 + 0.05) / (l2 + 0.05)

def hex_to_rgb(hex_str):
    hex_str = hex_str.lstrip('#')
    if len(hex_str) == 3: hex_str = ''.join(c+c for c in hex_str)
    return tuple(int(hex_str[i:i+2], 16) for i in (0, 2, 4))

def rgb_to_hex(r, g, b):
    return f"#{int(r):02x}{int(g):02x}{int(b):02x}"

def generate_palette(is_dark):
    # Base hues (R, G, B, M, Y, C, etc)
    hues = [
        (0, 100, 255),    # Blue
        (255, 150, 0),    # Orange
        (150, 0, 255),    # Purple
        (0, 200, 100),    # Green
        (255, 50, 100),   # Pink/Red
        (0, 200, 255),    # Cyan
        (200, 200, 0),    # Yellow
        (100, 100, 255),  # Light Blue
        (255, 0, 200)     # Magenta
    ]
    
    palette = []
    for i, (r, g, b) in enumerate(hues):
        if is_dark:
            # For dark background, we need light colors
            # Alternate between VERY light and MEDIUM light
            if i % 2 == 0:
                # Very light
                r = r + (255 - r) * 0.8
                g = g + (255 - g) * 0.8
                b = b + (255 - b) * 0.8
            else:
                # Medium
                r = r + (255 - r) * 0.4
                g = g + (255 - g) * 0.4
                b = b + (255 - b) * 0.4
        else:
            # For light background, we need dark colors
            # Alternate between VERY dark and MEDIUM dark
            if i % 2 == 0:
                # Very dark
                r = r * 0.4
                g = g * 0.4
                b = b * 0.4
            else:
                # Medium
                r = r * 0.8
                g = g * 0.8
                b = b * 0.8
        palette.append(rgb_to_hex(r, g, b))
    return palette

def process_css(css):
    # Remove existing --chart-X definitions
    css = re.sub(r'\s*--chart-\d+:[^;]+;', '', css)
    
    # We need to insert the light palette before the closing brace of :root
    light_pal = generate_palette(False)
    dark_pal = generate_palette(True)
    
    def light_repl(m):
        insert = "\n"
        for i, c in enumerate(light_pal):
            insert += f"    --chart-{i+1}: {c};\n"
        return insert + m.group(1)
        
    css = re.sub(r'(\n\})', light_repl, css, count=1)
    
    # Now for dark mode
    def dark_repl(m):
        insert = "\n"
        for i, c in enumerate(dark_pal):
            insert += f"        --chart-{i+1}: {c};\n"
        return insert + m.group(1)
        
    if '[data-theme="dark"]' in css:
        # Find the closing brace of data-theme="dark"
        # Since it's regex, it's tricky. Let's just find [data-theme="dark"] { ... }
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
                # idx-1 is the closing brace
                css = css[:idx-1] + "\n" + "".join([f"    --chart-{i+1}: {c};\n" for i, c in enumerate(dark_pal)]) + css[idx-1:]
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

