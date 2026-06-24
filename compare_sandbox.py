import os
import filecmp

d1 = "writer-sandbox/docs"
d2 = "parser/assets/structure/docs"

files1 = set(os.listdir(d1))
files2 = set(os.listdir(d2))

print("Missing in structure:")
for f in sorted(files1 - files2):
    print("  " + f)

print("\nDifferences:")
for f in sorted(files1.intersection(files2)):
    p1 = os.path.join(d1, f)
    p2 = os.path.join(d2, f)
    if not filecmp.cmp(p1, p2, shallow=False):
        print(f"  {f} differs! sandbox={os.path.getsize(p1)}, structure={os.path.getsize(p2)}")

