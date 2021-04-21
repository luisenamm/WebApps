
/*self.addEventListener("fetch", function(event) {
  event.respondWith(
    fetch(event.request).catch(function() {
      return new Response(
      "I sit on a man’s back choking him and making him carry me, and yet assure myself and others that I am sorry for him and wish to lighten his load by all means possible…except by getting off his back."+
      " Tolstoi"
      );
    })
  );
});*/

self.addEventListener("fetch", function(event) {
  if (event.request.url.includes("segato")) {
    event.respondWith(
      fetch("library.jpg")
    );
  }
});

/*self.addEventListener("fetch", function(event) {
  if (event.request.url.includes("segato.jpg")) {
  	 console.log("Fetch request for:", event.request.url);
    event.respondWith(
      new Response(
        "body {background: black!important;}",
        { headers: { "Content-Type": "text/css" }}
      )
    );
  }
});*/

