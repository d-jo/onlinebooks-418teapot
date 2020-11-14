
function loadListings(data) {
  console.log("success")
  console.log(data)
  for (let i = 0; i < data.length; i++) {

    var listing_tmpl = `<div class='single-container'>
      <div class='single-header'>
        <a href='/listing/${data[i]['id']}' class='h3 lst-title'>${data[i]['title']}</a>
      </div>
      <div class='single-body'>
        <div class='row'>
          <div class='col-8 lst-desc'>
            <p class='lst-desc'>${data[i]['description']}</p>
          </div>
          <div class='col-4 lst-details'>
            <p class='lst-price'>Price: <b>$${data[i]['price']}</b></p>
            <p class='lst-isbn'>ISBN: ${data[i]['isbn']}</p>
          </div>
        </div>
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