document.addEventListener("DOMContentLoaded", function(){
    var modalImage = document.querySelector("#group-image");
    var modalInfos = document.querySelector("#group-infos");
 
    var screenWidth = window.matchMedia("(min-width: 968px)");
    modalImage.onload = function() {
        resizeModalInfos(screenWidth);
    }
 
    screenWidth.addEventListener("change", function(e) {
        resizeModalInfos(e);
    });

    // Afficher le Modal
    var cards = document.querySelectorAll(".artist-card");
    cards.forEach(card => {
        // Evénement clic du bouton "VOIR"
        card.querySelector(".voir-zone").addEventListener("click", function(){
            // Ajout des données de l'artiste dans le modal
            document.querySelector("#group-image").src = card.querySelector(".artist-image").src;
            document.querySelector("#modal-name").innerHTML = card.querySelector(".artist-name").innerHTML;
            document.querySelector("#modal-creation_date").innerHTML = card.querySelector(".artist-year").innerHTML;
            document.querySelector("#modal-members").innerHTML = card.querySelector(".artist-members").innerHTML;
            document.querySelector("#modal-first_album").innerHTML = card.querySelector(".artist-firstAlbum").innerHTML;

            // Affichange du modal
            openModal();
        });
    });
    
    document.addEventListener("click", function(event){
        if (event.target.matches("#modal-close") || event.target.matches("#modal")) {
            closeModal();
        }
    });


    // Fonctions
    // ---------------------------------------------------------------------------
    function resizeModalInfos(e) {
        if (e.matches) {
            modalInfos.style.height = modalImage.offsetHeight+"px";
        }
    }

    function openModal() {
        document.querySelector("#modal").style.visibility = "visible";
        document.querySelector("#modal").style.opacity = "1";
        document.querySelector("#modal").style.backdropFilter = "blur(6px)";
    }

    function closeModal() {
        document.querySelector("#modal").style.visibility = "hidden";
        document.querySelector("#modal").style.opacity = "0";
        document.querySelector("#modal").style.backdropFilter = "blur(0px)";
    }
    // ---------------------------------------------------------------------------
 
 });
 