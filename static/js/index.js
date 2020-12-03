
var data = []
var full_data = []

function loadListings() {
  $("#listings").empty();
  console.log("success");
  console.log(data);
  for (let i = 0; i < data.length; i++) {

    var listing_tmpl = `<div class='single-container'>
      <div class='single-header'>
        <a href='/listing/${data[i]['id']}' class='h3 lst-title'>${data[i]['title']}</a>
        <b class='h4 lst-status ${data[i]['status']}'>${data[i]['status']}</b>
      </div>
      <div class='single-body'>
        <div class='row'>
          <div class='col-8 lst-desc'>
            <p class='lst-desc'>${data[i]['description']}</p>
          </div>
          <div class='col-4 lst-details'>
            <p class='lst-price'>Price: <b>$${data[i]['price'].toFixed(2)}</b></p>
            <p class='lst-isbn'>ISBN: ${data[i]['isbn']}</p>
          </div>
        </div>
      </div>
    </div>`
    let lst = $(listing_tmpl);
    $('#listings').append(lst);
  }
}

function viewActiveOnly() {
  let new_data = []
  for (let i = 0; i < data.length; i++) {
    if (data[i]['status'] == "active") {
      new_data.push(data[i]);
    }
  }
  data = new_data;
  loadListings();
}

function viewAll() {
  data = full_data;
  loadListings();
}

function loadAllActive() {
  $.ajax({
    type: "GET",
    url: "/active",
    success: (o) => {
      data = o;
      full_data = o;
      loadListings();
    },
    error: (err) => {
      console.log("error");
      $('#listings').append("Error: " + err.message);
    }
  })
}

$(() => {
  console.log("Doc Loaded")
  loadAllActive();
})