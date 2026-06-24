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

pal_light = ["#2563eb", "#10b981", "#d946ef", "#f59e0b", "#7c3aed", "#0ea5e9", "#e11d48", "#84cc16", "#475569"]
pal_dark  = ["#3b82f6", "#6ee7b7", "#c084fc", "#fde047", "#f472b6", "#67e8f9", "#fb7185", "#bef264", "#94a3b8"]

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
print("Dark:")
check(pal_dark, 0.0)

