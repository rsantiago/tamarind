import os, re, colorsys

def relative_luminance(r, g, b):
    def adjust(c):
        c = c / 255.0
        return c / 12.92 if c <= 0.03928 else ((c + 0.055) / 1.055) ** 2.4
    return 0.2126 * adjust(r) + 0.7152 * adjust(g) + 0.0722 * adjust(b)

def contrast_ratio(l1, l2):
    if l1 < l2: l1, l2 = l2, l1
    return (l1 + 0.05) / (l2 + 0.05)

def hex_to_rgb(hex_str):
    if not hex_str or not hex_str.startswith('#'): return (128,128,128)
    hex_str = hex_str.lstrip('#')
    if len(hex_str) == 3: hex_str = ''.join(c+c for c in hex_str)
    try:
        return tuple(int(hex_str[i:i+2], 16) for i in (0, 2, 4))
    except:
        return (128,128,128)

def rgb_to_hex(r, g, b):
    return f"#{int(max(0,min(255,r))):02x}{int(max(0,min(255,g))):02x}{int(max(0,min(255,b))):02x}"

def hsl_to_hex(h, s, l):
    r, g, b = colorsys.hls_to_rgb(h/360.0, l, s)
    return rgb_to_hex(r*255, g*255, b*255)

def adjust_color_for_contrast(hex_color, target_luminance_type, bg_lum):
    r, g, b = hex_to_rgb(hex_color)
    h, l, s = colorsys.rgb_to_hls(r/255.0, g/255.0, b/255.0)
    
    if s > 0.65: s = 0.65 
    
    if target_luminance_type == "dark": l = min(l, 0.35)
    else: l = max(l, 0.65)
        
    while True:
        r_new, g_new, b_new = colorsys.hls_to_rgb(h, l, s)
        lum = relative_luminance(r_new*255, g_new*255, b_new*255)
        if contrast_ratio(lum, bg_lum) > 1.4: break
        if target_luminance_type == "dark":
            l -= 0.05
            if l <= 0: break
        else:
            l += 0.05
            if l >= 1: break
            
    return hsl_to_hex(h*360, s, l)

def generate_theme_palette(primary_hex, bg_hex):
    bg_r, bg_g, bg_b = hex_to_rgb(bg_hex)
    bg_lum = relative_luminance(bg_r, bg_g, bg_b)
    is_dark_bg = bg_lum < 0.5
    
    targets = []
    for i in range(9):
        if is_dark_bg: targets.append("light" if i % 2 == 0 else "dark")
        else: targets.append("dark" if i % 2 == 0 else "light")
            
    raw_colors = [
        primary_hex,
        hsl_to_hex(140, 0.5, 0.5), # Forest
        hsl_to_hex(45, 0.7, 0.5),  # Mustard
        hsl_to_hex(210, 0.3, 0.5), # Slate
        hsl_to_hex(345, 0.6, 0.5), # Burgundy
        hsl_to_hex(75, 0.5, 0.5),  # Olive
        hsl_to_hex(220, 0.6, 0.5), # Navy
        hsl_to_hex(15, 0.6, 0.5),  # Rust
        hsl_to_hex(0, 0.0, 0.5),   # Charcoal
    ]
    
    palette = []
    prev_lum = None
    for i in range(9):
        c = adjust_color_for_contrast(raw_colors[i], targets[i], bg_lum)
        if prev_lum is not None:
            r, g, b = hex_to_rgb(c)
            lum = relative_luminance(r, g, b)
            if contrast_ratio(lum, prev_lum) < 1.35:
                h, l, s = colorsys.rgb_to_hls(r/255.0, g/255.0, b/255.0)
                if targets[i] == "dark": l = max(0, l - 0.15)
                else: l = min(1, l + 0.15)
                r, g, b = colorsys.hls_to_rgb(h, l, s)
                c = rgb_to_hex(r*255, g*255, b*255)
                lum = relative_luminance(r*255, g*255, b*255)
        
        palette.append(c)
        r, g, b = hex_to_rgb(c)
        prev_lum = relative_luminance(r, g, b)
        
    return palette

def extract_color(css, var_name, default):
    match = re.search(r'{}\s*:\s*(#[0-9a-fA-F]{{3,6}})'.format(var_name), css)
    if match: return match.group(1)
    return default

def process_css(css):
    # Extract colors from :root
    root_match = re.search(r':root\s*\{([^}]*)\}', css)
    root_css = root_match.group(1) if root_match else ""
    
    light_primary = extract_color(root_css, '--primary-color', '#3b82f6')
    light_bg = extract_color(root_css, '--background-color', '#ffffff')
    if extract_color(root_css, '--card-bg', None):
        light_bg = extract_color(root_css, '--card-bg', '#ffffff')
        
    pal_light = generate_theme_palette(light_primary, light_bg)
    
    dark_match = re.search(r'\[data-theme="dark"\]\s*\{([^}]*)\}', css)
    dark_css = dark_match.group(1) if dark_match else ""
    
    dark_primary = extract_color(dark_css, '--primary-color', light_primary)
    dark_bg = extract_color(dark_css, '--background-color', '#0f172a')
    if extract_color(dark_css, '--card-bg', None):
        dark_bg = extract_color(dark_css, '--card-bg', '#0f172a')
        
    pal_dark = generate_theme_palette(dark_primary, dark_bg)

    # Clean old ones
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

