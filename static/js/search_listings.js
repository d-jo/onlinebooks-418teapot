
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
      console.log("succ")
      console.log(o);
      if (Number.isInteger(o)) {
        window.location.href = "/search/" + o;
      } else {
        // error
        alert('error');
      }
    },
    error: (err) => {
      console.log("err")
      console.log(err)
    },
    dataType: "json"
  })
   
}