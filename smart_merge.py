import os
import shutil

d1 = "parser/assets/docs"
d2 = "parser/assets/structure/docs"

files1 = set(os.listdir(d1))
files2 = set(os.listdir(d2))

# Ensure all files from d1 are in d2, using the largest one
for f in files1:
    p1 = os.path.join(d1, f)
    p2 = os.path.join(d2, f)
    
    if not os.path.exists(p2):
        print(f"Copying missing file {f} to structure/docs")
        shutil.copy2(p1, p2)
    else:
        s1 = os.path.getsize(p1)
        s2 = os.path.getsize(p2)
        if s1 > s2 + 100:  # Only overwrite if it's significantly larger
            print(f"Overwriting {f} in structure/docs (docs is larger: {s1} > {s2})")
            shutil.copy2(p1, p2)

print("Merge complete.")
