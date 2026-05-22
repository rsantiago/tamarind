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

    applyTheme(getInitialTheme());

    if (toggleButton) {
        toggleButton.addEventListener('click', toggleTheme);
    }

    if (window.matchMedia) {
        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
            if (!readStoredTheme()) {
                applyTheme(e.matches ? 'dark' : 'light');
            }
        });
    }

    window.addEventListener('storage', (event) => {
        if (event.key === THEME_KEY && event.newValue) {
            html.setAttribute('data-theme', event.newValue);
        }
    });
})();
