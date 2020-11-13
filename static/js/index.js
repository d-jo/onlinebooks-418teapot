
function loadListings(data) {
  console.log("success")
  console.log(data)
  for (let i = 0; i < data.length; i++) {

    var listing_tmpl = `<div class='single-container'>
      <div class='single-header'>
        <a href='/listing/${data[i]['id']}' class='lst-title'>${data[i]['title']}</a>
        <p class='lst-isbn'>${data[i]['isbn']}</p>
        <p class='lst-price'>${data[i]['price']}</p>
      </div>
      <div class='single-body'>
        <p class='lst-desc'>${data[i]['description']}</p>
      </div>
    </div>`
    let lst = $(listing_tmpl);
    $('#listings').append(lst);
    
  }

}

$(() => {
  console.log("Doc Loaded")
  $.ajax({
    type: "GET",
    url: "/active",
    success: loadListings,
    error: (err) => {
      console.log("error")
      $('#listings').append("Error: " + err.message);
    }
  })
})