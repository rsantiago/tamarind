import glob

alert = """
{{ callout type="info" title="Architecture Note" }}
Curious about how Tamarind mathematically guarantees elegant, theme-adaptive, and highly legible color contrast sequences across all 31 themes? Check out the [Hidden Chart Color Theory](/docs/hidden-color-theory.html) architecture document.
{{ /callout }}

Tamarind includes"""

files = [
    "parser/assets/docs/shortcode-charts.md",
    "parser/assets/structure/docs/shortcode-charts.md",
    "writer-sandbox/docs/shortcode-charts.md"
]

for f in files:
    try:
        with open(f, "r") as file:
            content = file.read()
            
        new_content = content.replace("Tamarind includes", alert)
        
        with open(f, "w") as file:
            file.write(new_content)
        print(f"Updated {f}")
    except Exception as e:
        print(f"Error on {f}: {e}")
