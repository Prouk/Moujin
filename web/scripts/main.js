document.addEventListener("DOMContentLoaded", function() {
    const darkThemeMq = window.matchMedia("(prefers-color-scheme: dark)")
    if (darkThemeMq.matches) {
        document.documentElement.classList.add('dark-mode')
    }
    const darkModeButton = document.getElementById('dark-mode-button')
    darkModeButton.onclick = function () {
        document.documentElement.classList.toggle('dark-mode')
    }
});

function toggleFullCard(el) {
    el.classList.toggle("full")
}