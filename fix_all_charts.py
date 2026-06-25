import os, re, colorsys

def relative_luminance(r, g, b):
    def adjust(c):
        c = c / 255.0
        return c / 12.92 if c <= 0.03928 else ((c + 0.055) / 1.055) ** 2.4
    return 0.2126 * adjust(r) + 0.7152 * adjust(g) + 0.0722 * adjust(b)

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

def force_luminance(h, s, target_lum):
    low, high = 0.0, 1.0
    best_l = 0.5
    for _ in range(20):
        mid = (low + high) / 2
        r, g, b = colorsys.hls_to_rgb(h/360.0, mid, s)
        lum = relative_luminance(r*255, g*255, b*255)
        if lum < target_lum:
            low = mid
        else:
            high = mid
        best_l = mid
    return hsl_to_hex(h, s, best_l)

def generate_theme_palette(primary_hex, bg_hex):
    bg_r, bg_g, bg_b = hex_to_rgb(bg_hex)
    bg_lum = relative_luminance(bg_r, bg_g, bg_b)
    
    if bg_lum > 0.5: 
        lum1 = 0.50 
        lum2 = 0.20 
    else: 
        lum1 = 0.45 
        lum2 = 0.80 
        
    pr, pg, pb = hex_to_rgb(primary_hex)
    ph, pl, ps = colorsys.rgb_to_hls(pr/255.0, pg/255.0, pb/255.0)
    
    hues = [
        (ph * 360, min(ps, 0.85)), 
        (180, 0.8), 
        (15,  0.8), 
        (300, 0.8), 
        (140, 0.7), 
        (45,  0.9), 
        (210, 0.7), 
        (345, 0.7), 
        (270, 0.6), 
    ]
    
    palette = []
    for i in range(9):
        h, s = hues[i]
        target = lum1 if i % 2 == 0 else lum2
        palette.append(force_luminance(h, s, target))
        
    return palette

def extract_color(css, var_name, default):
    match = re.search(r'{}\s*:\s*(#[0-9a-fA-F]{{3,6}})'.format(var_name), css)
    if match: return match.group(1)
    return default

def process_css(css):
    light_primary = '#3b82f6'
    light_bg_val = '#ffffff' # keep track of what :root defines
    
    css = re.sub(r'\s*--chart-\d+:[^;]+;', '', css)

    regexes = [
        r':root\s*\{',
        r'\[data-theme="light"\]\s*\{',
        r'\[data-theme="dark"\]\s*\{',
        r'@media\s*\(prefers-color-scheme:\s*dark\)\s*\{\s*:root\s*\{'
    ]
    
    # Process :root first to capture fallbacks
    for idx_reg, reg in enumerate(regexes):
        match = re.search(reg, css)
        if match:
            start_brace = match.end() - 1
            if css[start_brace] != '{': continue
            brace_count = 1
            if 'prefers-color-scheme' in reg: brace_count = 2
            
            end_brace = start_brace + 1
            while end_brace < len(css) and brace_count > 0:
                if css[end_brace] == '{': brace_count += 1
                elif css[end_brace] == '}': brace_count -= 1
                end_brace += 1
                
            if brace_count == 0:
                block_css = css[start_brace:end_brace]
                
                # Setup fallbacks based on if it's light or dark block
                if 'dark' in reg:
                    # For dark mode, if they don't define a BG, they inherit the LIGHT BG! 
                    # Yes, standard CSS cascade!
                    bg_def = light_bg_val 
                else:
                    bg_def = '#ffffff'
                    
                primary = extract_color(block_css, '--primary-color', light_primary)
                bg = extract_color(block_css, '--background-color', bg_def)
                if extract_color(block_css, '--card-bg', None):
                    bg = extract_color(block_css, '--card-bg', bg_def)
                    
                if reg == r':root\s*\{':
                    light_primary = primary
                    light_bg_val = bg
                    
                pal = generate_theme_palette(primary, bg)
                
                insert = "\n"
                for i, c in enumerate(pal):
                    insert += f"    --chart-{i+1}: {c};\n"
                    
                if 'prefers-color-scheme' in reg:
                    last_rbrace = block_css.rfind('}')
                    sec_last_rbrace = block_css.rfind('}', 0, last_rbrace)
                    new_block = block_css[:sec_last_rbrace] + insert + block_css[sec_last_rbrace:]
                else:
                    last_rbrace = block_css.rfind('}')
                    new_block = block_css[:last_rbrace] + insert + block_css[last_rbrace:]
                    
                css = css[:start_brace] + new_block + css[end_brace:]

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

