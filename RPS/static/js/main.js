let imgArray = [
    "static/img/hand_rock.png",
    "static/img/hand_paper.png",
    "static/img/hand_scissors.png",
];
function choose(x){
    fetch("/Play?c=" + x)
    .then(response => response.json())
    .then(data => {
        switch (x){
            case 0:
                document.getElementById("player_choice").innerHTML = "EL jugador eligio PIEDRA.";
                break
            case 1:
                document.getElementById("player_choice").innerHTML = "EL jugador eligio PAPEL.";
                break
            case 2:
                document.getElementById("player_choice").innerHTML = "EL jugador eligio TIJERA.";
                break
        }
        document.getElementById("player_score").innerHTML = data.player_score;
        document.getElementById("computer_score").innerHTML = data.computer_score;
        document.getElementById("computer_choice").innerHTML = data.computer_choise;
        document.getElementById("round_result").innerHTML = data.round_result;
        document.getElementById("round_message").innerHTML = data.message;
        
        var imgElement = document.getElementById("img_computer");
        imgElement.src = imgArray[data.computer_choise_int];
    })
}