self.addEventListener("fetch", function(event) {
  if (event.request.url.includes("hola")) {
    event.respondWith(
      fetch("semestrei.png")
    );
  }else{
    event.respondWith(
      fetch(event.request.url).then(response => {
        if(response.status === 404){              
          return fetch("library.jpg");
        }return response;  
      })
    )
  }
});
