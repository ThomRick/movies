$(function () {
    setInterval(function () {
        $.getJSON("game/current-movie", function (data) {
            $("#movie").text(data.movie);
        });
    }, 1000);
});
