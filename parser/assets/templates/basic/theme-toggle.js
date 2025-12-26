/**
 * Theme Toggle Script
 * Handles light/dark theme switching with localStorage persistence
 */

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
            // localStorage can fail on file:// or privacy modes; fall through to window.name
        }

        const nameData = safeParseWindowName();
        return nameData[WINDOW_NAME_KEY] || null;
    }

    function writeStoredTheme(theme) {
        try {
            localStorage.setItem(THEME_KEY, theme);
        } catch (err) {
            // Ignore write failures (private mode / file:// origins)
        }

        const nameData = safeParseWindowName();
        nameData[WINDOW_NAME_KEY] = theme;
        try {
            window.name = JSON.stringify(nameData);
        } catch (err) {
            // Swallow; window.name may be blocked in some contexts
        }
    }

    // Get saved theme or detect system preference
    function getInitialTheme() {
        const savedTheme = readStoredTheme();
        if (savedTheme === 'light' || savedTheme === 'dark') {
            return savedTheme;
        }

        // Check system preference
        if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
            return 'dark';
        }

        return 'light';
    }

    // Apply theme to document
    function applyTheme(theme) {
        html.setAttribute('data-theme', theme);
        writeStoredTheme(theme);
    }

    // Toggle between themes
    function toggleTheme() {
        const currentTheme = html.getAttribute('data-theme') || 'light';
        const newTheme = currentTheme === 'light' ? 'dark' : 'light';
        applyTheme(newTheme);
    }

    // Initialize theme
    applyTheme(getInitialTheme());

    // Add click listener to toggle button
    if (toggleButton) {
        toggleButton.addEventListener('click', toggleTheme);
    }

    // Listen for system theme changes
    if (window.matchMedia) {
        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
            // Only auto-switch if user hasn't manually set a preference
            if (!readStoredTheme()) {
                applyTheme(e.matches ? 'dark' : 'light');
            }
        });
    }

    // Sync theme across tabs/windows where localStorage is available
    window.addEventListener('storage', (event) => {
        if (event.key === THEME_KEY && event.newValue) {
            html.setAttribute('data-theme', event.newValue);
        }
    });
})();
