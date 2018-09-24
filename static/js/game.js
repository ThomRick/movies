$(function () {
    setInterval(function () {
        $.getJSON("game/current-movie", function (data) {
            $("#movie").text(data.movie);
        });
    }, 1000);

    setInterval(function () {
        $.getJSON("game/players", function (data) {
            var playerList = "<ul>";
            for (var i in data.players) {
                var player = data.players[i];
                playerList += "<li>" + player.Name + "(" + player.Score + ")</li>"
            }
            playerList += "</ul>";
            $("#players").html(playerList);
        });
    }, 1000);
});
