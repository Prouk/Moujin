if (document.readyState === "complete") {
    ready();
} else {
    document.addEventListener("DOMContentLoaded", ready);
}

function ready() {
    console.log("ready")
    const html = document.getElementsByTagName( 'html' )[0]
    const darkmodeButton = document.getElementsByClassName("darkmode")[0]
    function switchDarkMode() {
        console.log('clicky click')
        if (darkmodeButton.classList.contains("on")) {
            darkmodeButton.classList.remove("on")
            html.classList.remove("dark")
        } else {
            darkmodeButton.classList.add("on")
            html.classList.add("dark")
        }
    }
    darkmodeButton.onclick = switchDarkMode

}