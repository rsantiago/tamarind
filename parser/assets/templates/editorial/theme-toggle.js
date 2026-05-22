(function() {
    const THEME_KEY = 'site-theme';
    const WINDOW_NAME_KEY = '__site_theme__';
    const html = document.documentElement;
    const toggleButton = document.getElementById('theme-toggle');

    function safeParseWindowName() {
        try {
            const parsed = JSON.parse(window.name || '{}');
            return typeof parsed === 'object' && parsed !== null ? parsed : {};
        } catch (err) {
            return {};
        }
    }

    // Read theme from localStorage or window.name fallback
    function readStoredTheme() {
        try {
            const stored = localStorage.getItem(THEME_KEY);
            if (stored) return stored;
        } catch (err) {
            // ignore
        }

        const nameData = safeParseWindowName();
        return nameData[WINDOW_NAME_KEY] || null;
    }

    // Save theme to localStorage and window.name
    function writeStoredTheme(theme) {
        try {
            localStorage.setItem(THEME_KEY, theme);
        } catch (err) {
            // ignore
        }

        const nameData = safeParseWindowName();
        nameData[WINDOW_NAME_KEY] = theme;
        try {
            window.name = JSON.stringify(nameData);
        } catch (err) {
            // ignore
        }
    }

    function getInitialTheme() {
        const savedTheme = readStoredTheme();
        if (savedTheme === 'light' || savedTheme === 'dark') {
            return savedTheme;
        }

        if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
            return 'dark';
        }

        return 'light';
    }

    function applyTheme(theme) {
        html.setAttribute('data-theme', theme);
        writeStoredTheme(theme);
    }

    function toggleTheme() {
        const currentTheme = html.getAttribute('data-theme') || 'light';
        const newTheme = currentTheme === 'light' ? 'dark' : 'light';
        applyTheme(newTheme);
    }

    // Initialize theme immediately
    applyTheme(getInitialTheme());

    // Register toggle event listener on DOMContentLoaded or immediately
    if (toggleButton) {
        toggleButton.addEventListener('click', toggleTheme);
    } else {
        document.addEventListener('DOMContentLoaded', () => {
            const btn = document.getElementById('theme-toggle');
            if (btn) {
                btn.addEventListener('click', toggleTheme);
            }
        });
    }

    // Handle system theme preference changes
    if (window.matchMedia) {
        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
            if (!readStoredTheme()) {
                applyTheme(e.matches ? 'dark' : 'light');
            }
        });
    }

    // Sync theme across tabs
    window.addEventListener('storage', (event) => {
        if (event.key === THEME_KEY && event.newValue) {
            html.setAttribute('data-theme', event.newValue);
        }
    });
})();
