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

palette = [
    "#3b82f6", # 1
    "#10b981", # 2
    "#58508d", # 3
    "#ffa600", # 4
    "#2f4b7c", # 5
    "#00ba38", # 6
    "#1e2e33", # 7
    "#ff6361", # 8
    "#a05195", # 9
]

lums = []
for c in palette:
    r, g, b = hex_to_rgb(c)
    l = relative_luminance(r, g, b)
    lums.append(l)

ok = True
for i in range(8):
    cr = contrast_ratio(lums[i], lums[i+1])
    if cr < 1.35:
        print(f"FAILED ADJACENT: {palette[i]} vs {palette[i+1]}: {cr:.2f}")
        ok = False

if ok:
    print("NEW DEFAULT COMBINATION IS PERFECT!")
