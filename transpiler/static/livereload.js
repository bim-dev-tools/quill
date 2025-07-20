(function () {
  var ws = new WebSocket("ws://" + location.host + "/livereload");
  ws.onmessage = function (msg) {
    if (msg.data === "reload") {
      location.reload();
    }
  };
})();
