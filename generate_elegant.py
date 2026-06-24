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
    hex_str = hex_str.lstrip('#')
    if len(hex_str) == 3: hex_str = ''.join(c+c for c in hex_str)
    return tuple(int(hex_str[i:i+2], 16) for i in (0, 2, 4))

def rgb_to_hex(r, g, b):
    return f"#{int(r):02x}{int(g):02x}{int(b):02x}"

def hsl_to_hex(h, s, l):
    r, g, b = colorsys.hls_to_rgb(h/360.0, l, s)
    return rgb_to_hex(r*255, g*255, b*255)

def adjust_color_for_contrast(hex_color, target_luminance_type, bg_lum):
    # target_luminance_type: "dark" or "light"
    # We tweak the HSL lightness until it passes
    r, g, b = hex_to_rgb(hex_color)
    h, l, s = colorsys.rgb_to_hls(r/255.0, g/255.0, b/255.0)
    
    # Professional desaturation
    if s > 0.65: s = 0.65 
    
    # Adjust lightness based on target type
    if target_luminance_type == "dark":
        # Force it to be a dark variant
        l = min(l, 0.35)
    else:
        # Force it to be a light variant
        l = max(l, 0.65)
        
    # Check contrast vs background
    while True:
        r_new, g_new, b_new = colorsys.hls_to_rgb(h, l, s)
        lum = relative_luminance(r_new*255, g_new*255, b_new*255)
        if contrast_ratio(lum, bg_lum) > 1.4:
            break
        if target_luminance_type == "dark":
            l -= 0.05
            if l <= 0: break
        else:
            l += 0.05
            if l >= 1: break
            
    return hsl_to_hex(h*360, s, l)

# Hues for our elegant palette
# 1. Primary (Extracted)
# 2. Forest Green (H=140)
# 3. Mustard (H=45)
# 4. Slate (H=210, low saturation)
# 5. Burgundy (H=345)
# 6. Olive (H=75)
# 7. Navy (H=220)
# 8. Rust/Terracotta (H=15)
# 9. Charcoal (H=0, S=0)

def generate_theme_palette(primary_hex, bg_hex):
    bg_r, bg_g, bg_b = hex_to_rgb(bg_hex)
    bg_lum = relative_luminance(bg_r, bg_g, bg_b)
    is_dark_bg = bg_lum < 0.5
    
    # If dark bg, base colors should be light (so target "light" first)
    # If light bg, base colors should be dark (so target "dark" first)
    targets = []
    for i in range(9):
        if is_dark_bg:
            targets.append("light" if i % 2 == 0 else "dark")
        else:
            targets.append("dark" if i % 2 == 0 else "light")
            
    # Base raw colors
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
        
        # Ensure it contrasts with previous
        if prev_lum is not None:
            r, g, b = hex_to_rgb(c)
            lum = relative_luminance(r, g, b)
            if contrast_ratio(lum, prev_lum) < 1.35:
                # Force more contrast
                h, l, s = colorsys.rgb_to_hls(r/255.0, g/255.0, b/255.0)
                if targets[i] == "dark":
                    l = max(0, l - 0.15)
                else:
                    l = min(1, l + 0.15)
                r, g, b = colorsys.hls_to_rgb(h, l, s)
                c = rgb_to_hex(r*255, g*255, b*255)
                lum = relative_luminance(r*255, g*255, b*255)
        
        palette.append(c)
        r, g, b = hex_to_rgb(c)
        prev_lum = relative_luminance(r, g, b)
        
    return palette

# Test
print("Light theme palette (Primary: #2563eb, BG: #ffffff):")
print(generate_theme_palette("#2563eb", "#ffffff"))
print("Dark theme palette (Primary: #60a5fa, BG: #1e293b):")
print(generate_theme_palette("#60a5fa", "#1e293b"))
