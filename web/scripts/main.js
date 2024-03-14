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

function goFullCard(el) {
    el.classList.add("full")
}

function goSmallCard(e) {
    e.target.parentElement.classList.remove('full');
    e.stopPropagation();
}