import glob
import os
import re

for filename in glob.glob("parser/internal/builder/plugin_*.go"):
    if filename == "parser/internal/builder/plugin_extension.go":
        continue
    if filename == "parser/internal/builder/plugin_chart.go":
        continue
    if filename == "parser/internal/builder/plugin_chart_register.go":
        continue

    with open(filename, "r") as f:
        content = f.read()

    # Find the New...Plugin function
    m = re.search(r'func (New[A-Za-z0-9_]+Plugin)\(\)\s*\*?[A-Za-z0-9_]+Plugin', content)
    if not m:
        m = re.search(r'func (New[A-Za-z0-9_]+Plugin)\(\)\s*\*?[A-Za-z0-9_]+', content)

    if m:
        func_name = m.group(1)
        if "func init()" not in content:
            init_code = f"\nfunc init() {{\n\tRegisterDefaultPlugin(func() ShortcodePlugin {{ return {func_name}() }})\n}}\n"
            with open(filename, "a") as f:
                f.write(init_code)
            print(f"Updated {filename} with {func_name}")
        else:
            print(f"Skipped {filename} (already has init)")
    else:
        print(f"Could not find constructor in {filename}")
