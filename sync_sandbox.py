import os
import shutil

d1 = "parser/assets/structure/docs"
d2 = "writer-sandbox/docs"

files = os.listdir(d1)
for f in files:
    p1 = os.path.join(d1, f)
    p2 = os.path.join(d2, f)
    if not os.path.exists(p2) or os.path.getsize(p1) != os.path.getsize(p2):
        print(f"Syncing {f} to sandbox")
        shutil.copy2(p1, p2)
print("Sync complete.")
