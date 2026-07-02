import os

d1 = "parser/assets/docs"
d2 = "parser/assets/structure/docs"

files1 = set(os.listdir(d1))
files2 = set(os.listdir(d2))

print("Common files:")
for f in sorted(files1.intersection(files2)):
    p1 = os.path.join(d1, f)
    p2 = os.path.join(d2, f)
    s1 = os.path.getsize(p1)
    s2 = os.path.getsize(p2)
    if s1 != s2:
        print(f"  {f}: docs={s1} bytes, structure={s2} bytes")
        
