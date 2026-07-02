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
    return tuple(int(hex_str[i:i+2], 16) for i in (0, 2, 4))

pal_light = [
    "#1a365d", # Navy
    "#d97706", # Mustard
    "#166534", # Forest
    "#9f1239", # Burgundy
    "#475569", # Slate
    "#b45309", # Rust
    "#4c1d95", # Eggplant
    "#65a30d", # Olive
    "#3f3f46", # Charcoal
]

def check(pal, bg):
    lums = [relative_luminance(*hex_to_rgb(c)) for c in pal]
    for i in range(8):
        cr = contrast_ratio(lums[i], lums[i+1])
        if cr < 1.35:
            print(f"FAILED ADJACENT: {pal[i]} vs {pal[i+1]}: {cr:.2f}")
    for i in range(9):
        if contrast_ratio(lums[i], bg) < 1.35:
            print(f"FAILED BG: {pal[i]}")
    print("Checked.")

print("Light:")
check(pal_light, 1.0)
