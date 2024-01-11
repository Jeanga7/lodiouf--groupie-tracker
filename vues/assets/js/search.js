document.querySelector("#search-box").addEventListener("click", function () {
    document.querySelector("#search-overlay").style.visibility = "visible";
    document.querySelector("#search-overlay").style.backdropFilter = "blur(6px)";
    document.querySelector("#search-overlay").style.opacity = "0.7";
    document.querySelector("#search-results").style.right = "0";
});

document.querySelector("#search-overlay").addEventListener("click", function () {
    document.querySelector("#search-overlay").style.visibility = "hidden";
    document.querySelector("#search-overlay").style.backdropFilter = "blur(0)";
    document.querySelector("#search-overlay").style.opacity = "0";
    document.querySelector("#search-results").style.right = "-360px";
});