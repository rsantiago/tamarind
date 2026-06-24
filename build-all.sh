#!/bin/bash
# Exit immediately if any command exits with a non-zero status
set -e

# Compile the parser binary
echo "Compiling parser..."
(cd parser && go build -o ../tamarind)

# Refresh writer-sandbox structure from latest binary assets
echo "Initializing fresh writer-sandbox structure..."
rm -rf writer-sandbox
./tamarind init

echo "Cleaning previous builds..."
rm -rf public-all website
mkdir -p public-all

# Dynamically discover all themes (directories in parser/assets/templates except shared)
echo "Discovering available themes..."
THEMES=()
for dir in parser/assets/templates/*/ ; do
    theme=$(basename "$dir")
    if [ "$theme" != "shared" ] && [ "$theme" != "*" ]; then
        THEMES+=("$theme")
    fi
done

echo "Found ${#THEMES[@]} themes: ${THEMES[*]}"

# Build each discovered theme
SUCCESSFUL_THEMES=()
for theme in "${THEMES[@]}"; do
    echo "Building theme: $theme..."
    # Run the build, capture errors but don't exit script if one theme fails to compile
    if ./tamarind build -theme "$theme" -url "/$theme"; then
        mv website public-all/"$theme"
        SUCCESSFUL_THEMES+=("$theme")
    else
        echo "Warning: Failed to build theme $theme, skipping."
    fi
done

# Dynamically generate a beautiful unified index.html listing all compiled themes
echo "Generating dynamic gallery index.html..."
cat <<EOF > public-all/index.html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tamarind Theme Incubator & Gallery</title>
    <style>
        body { 
            font-family: 'Inter', -apple-system, system-ui, sans-serif; 
            max-width: 900px; 
            margin: 0 auto; 
            padding: 40px 24px; 
            line-height: 1.6; 
            color: #1f2937; 
            background: #f9fafb;
        }
        .header {
            text-align: center;
            margin-bottom: 48px;
            padding-bottom: 24px;
            border-bottom: 1px solid #e5e7eb;
        }
        h1 { 
            color: #111827; 
            font-size: 2.5rem; 
            font-weight: 800; 
            letter-spacing: -0.025em; 
            margin-bottom: 8px;
        }
        .subtitle {
            color: #4b5563;
            font-size: 1.125rem;
        }
        .grid { 
            display: grid; 
            grid-template-columns: repeat(auto-fill, minmax(250px, 1fr)); 
            gap: 20px; 
            padding: 0; 
            list-style: none; 
            margin-top: 32px;
        }
        .card {
            display: flex;
            flex-direction: column;
        }
        .card a { 
            display: flex; 
            flex-direction: column;
            justify-content: space-between;
            flex-grow: 1;
            padding: 24px; 
            background: #ffffff; 
            border-radius: 12px; 
            text-decoration: none; 
            color: #111827; 
            font-weight: 600; 
            border: 1px solid #e5e7eb; 
            transition: all 0.2s ease-in-out;
            box-shadow: 0 1px 3px rgba(0,0,0,0.05);
            box-sizing: border-box;
        }
        .card a:hover { 
            background: #ffffff; 
            transform: translateY(-4px); 
            box-shadow: 0 10px 15px -3px rgba(0,0,0,0.1), 0 4px 6px -2px rgba(0,0,0,0.05);
            border-color: #3b82f6;
        }
        .theme-name {
            font-size: 1.25rem;
            text-transform: capitalize;
            color: #3b82f6;
            margin-bottom: 8px;
        }
        .theme-desc {
            font-weight: 400;
            font-size: 0.875rem;
            color: #6b7280;
        }
        .badge {
            display: inline-block;
            padding: 2px 8px;
            font-size: 0.75rem;
            font-weight: 500;
            border-radius: 9999px;
            background: #eff6ff;
            color: #1d4ed8;
            align-self: flex-start;
            margin-top: 12px;
        }
    </style>
</head>
<body>
    <div class="header">
        <h1>Tamarind Theme Incubator</h1>
        <p class="subtitle">Welcome! Every theme is fully compliant and pre-compiled directly from the exact same Markdown source block.</p>
    </div>
    
    <h2>Available Themes (${#SUCCESSFUL_THEMES[@]})</h2>
    <ul class="grid">
EOF

for theme in "${SUCCESSFUL_THEMES[@]}"; do
    # Provide nice details/badges based on theme status or categories if desired
    # We can detect Phase 3 / Phase 2 based on the theme name for extra premium feel
    PHASE_BADGE="Phase 2"
    if [[ " atlas canvas prose console forge classic bento gallery " =~ " $theme " ]]; then
        PHASE_BADGE="Phase 3"
    fi
    
    cat <<LI >> public-all/index.html
        <li class="card">
            <a href="$theme/">
                <div>
                    <div class="theme-name">$theme</div>
                    <div class="theme-desc">Beautiful static site design compiled with the $theme theme templates.</div>
                </div>
                <span class="badge">$PHASE_BADGE</span>
            </a>
        </li>
LI
done

cat <<EOF >> public-all/index.html
    </ul>
</body>
</html>
EOF

# Mirror public-all to website for native ./tamarind serve compatibility
echo "Mirroring to website directory..."
cp -r public-all website

echo "Done! All themes built successfully in public-all/ and website/ directories."
