$(function () {
    var displayCurrentMovieName = function () {
        $.getJSON("game/current-movie", function (data) {
            $("#movie").text(data.movie.Title + " " + data.movie.MsToNextMovie + "ms");
            setTimeout(function () {
                displayCurrentMovieName();
            }, data.movie.MsToNextMovie)
        });

    };

    displayCurrentMovieName();

    setInterval(function () {
        $.getJSON("game/players", function (data) {
            var playerList = "<ul>";
            for (var i in data.players) {
                var player = data.players[i];
                playerList += "<li>" + player.Name + "(" + (player.Score) + ")</li>"
            }
            playerList += "</ul>";
            $("#players").html(playerList);
        });
    }, 1000);

    $("#answer").keypress(function (e) {
        if (e.which === 13) {
            $.post("game/answer", {title: this.value}, function (data) {
                $("#message").text(data.message);
            });
            this.value = "";
        }
    });
});
