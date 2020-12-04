
function error_message(message) {
    alert(message);
  }


function search() {
    var searchKey = document.getElementById("search").value
    console.log(searchKey)
    // Create object
  let dataObj = {
    keyword: searchKey
  }

  console.log(dataObj)
  $.ajax({
    type: "POST",
    url: "/search",
    data: JSON.stringify(dataObj),
    success: (o) => {
      data = o;
      loadListings();
    },
    error: (err) => {
      console.log("err")
      console.log(err)
    },
    dataType: "json"
  })
   
}