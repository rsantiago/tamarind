import os
import glob
import re

def hex_to_rgb(hex_code):
    hex_code = hex_code.lstrip('#')
    if len(hex_code) == 3:
        hex_code = ''.join([c*2 for c in hex_code])
    if len(hex_code) != 6:
        return 0, 0, 0
    return tuple(int(hex_code[i:i+2], 16) for i in (0, 2, 4))

def rgb_to_brightness(r, g, b):
    return (r * 299 + g * 587 + b * 114) / 1000

for filename in glob.glob('parser/assets/templates/*/style.css'):
    with open(filename, 'r') as f:
        content = f.read()
    
    m = re.search(r'--primary-color:\s*(#[A-Fa-f0-9]+);', content)
    if not m:
        continue
        
    color_hex = m.group(1)
    r, g, b = hex_to_rgb(color_hex)
    brightness = rgb_to_brightness(r, g, b)
    
    # If brightness is > 128, the color is light, text should be dark
    btn_text_color = '#000000' if brightness > 128 else '#ffffff'
    
    if '--btn-text-color' not in content:
        # Inject the variable right after primary-color
        content = re.sub(
            r'(--primary-color:\s*#[A-Fa-f0-9]+;.*)',
            rf'\1\n    --btn-text-color: {btn_text_color};',
            content
        )
        
        with open(filename, 'w') as f:
            f.write(content)
        print(f"Updated {filename} with text color {btn_text_color} (primary: {color_hex})")

