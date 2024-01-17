// Open Search
document.querySelector("#search-box").addEventListener("click", function () {
    openSearch();
});

// Close Search
document.querySelectorAll("#search-overlay, #close-search").forEach(el => {
    el.addEventListener("click", function () {
        closeSearch();
    });
});

// Show results
document.querySelector("#input-search").addEventListener("keyup", function(){
    showMatches();
});

function showMatches() {
    // Récupérer la valeur de la recherche depuis l'élément input
    var toSearch = document.querySelector("#input-search").value.toLowerCase(); // Convertir en minuscules pour une recherche insensible à la casse
    var resCount = 0; // Initialiser le compteur de résultats

    // Parcourir tous les éléments avec la classe "search-info"
    document.querySelectorAll(".search-info").forEach(artist => {
        // Nettoyer le contenu HTML précédent (supprimer les balises <strong>)
        artist.innerHTML = artist.innerHTML.replaceAll('<strong class="found">', '');
        artist.innerHTML = artist.innerHTML.replaceAll('</strong>', '');

        // Obtenir le contenu textuel en minuscules pour une recherche insensible à la casse
        var content = artist.innerText.toLowerCase();

        // Vérifier si la recherche est vide ou si le contenu inclut le terme recherché
        if (toSearch.length === 0 || content.includes(toSearch)) {
            // Si la recherche est vide ou le contenu inclut le terme recherché
            // Remplacer le terme recherché par une version soulignée dans le texte
            artist.innerHTML = artist.innerText.replace(
                new RegExp('(' + toSearch + ')', 'gi'),
                '<strong class="found">$1</strong>'
            );
            artist.parentElement.style.display = "flex"; // Afficher l'élément parent avec flex
            resCount++; // Incrémenter le compteur de résultats
        } else {
            artist.parentElement.style.display = "none"; // Masquer l'élément parent
        }
    });

    // Afficher ou masquer l'élément "no-matches" en fonction du nombre de résultats
    document.querySelector("#no-matches").style.display = resCount !== 0 || toSearch.length === 0 ? "none" : "block";
}


// Acces to result element
document.querySelectorAll(".search-body").forEach(el => {
    el.addEventListener("click", function (){
        closeSearch();
        let id = this.getAttribute("id").split("_")[1];
        goToCard(id);
        openModal(id);
    });
});

function openSearch() {
    document.querySelector("#search-overlay").style.visibility = "visible";
    document.querySelector("#search-overlay").style.opacity = "1";
    document.querySelector("#search-overlay").style.backdropFilter = "blur(6px)";
    document.querySelector("#search-results").style.right = "0";

    // Disable page scrollbar
    document.body.style.overflow = "hidden";
}

function goToCard(id){
    document.querySelector("#card_"+id).scrollIntoView({behavior: 'smooth'});
}

function openModal(id) {
    document.querySelector("#see_"+id).click()
}

function closeSearch() {
    document.querySelector("#search-overlay").style.visibility = "hidden";
    document.querySelector("#search-overlay").style.backdropFilter = "blur(0px)";
    document.querySelector("#search-overlay").style.opacity = "0";
    document.querySelector("#search-results").style.right = "-360px";

    // Enable page scrollbar
    document.body.style.overflow = "auto";
}