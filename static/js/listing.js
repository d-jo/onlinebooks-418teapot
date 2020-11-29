
function formatPrice(price) {
  return "$" + price.toFixed(2);
}


function purchase() {
  // get buyer info
  let name = $("#buyer_name").val();
  let billing = $("#buyer_billing").val();
  let shipping = $("#buyer_shipping").val();
  dataObj = {
    "buyer": name,
    "billing_info": billing,
    "shipping_info":  shipping
  }

  $.ajax({
    type: "POST",
    url: "/listing/purchase/" + listing_id,
    data: JSON.stringify(dataObj),
    success: (o) => {
      console.log("succ")
      console.log(o);
      alert("success")
    },
    error: (err) => {
      console.log("err")
      console.log(err)
      alert("fail")
    },
    dataType: "json"
  })

}

//Process user input (Password) for delete listing
function submit_form_password() {
  // get form data
 // console.log("hello im in the submit_form_password func\n")
  let listing_password = $("#listing_password").val();

  // Create object
  let dataObj = {
    listing_password: listing_password
  } 
  
  console.log(dataObj)
  $.ajax({
    type: "POST",
    url: "/listing/delete/" + listing_id,
    data: JSON.stringify(dataObj),
    success: (o) => {
      console.log("succ")
      console.log(o);
      //alert("sucess in getting a response. got to line 40 in listing.js.\n") //when password wrong, it executes this line
      if (o == false) {
        alert("Wrong password. Please try again.") //seems to execute this line when wrong password
      }
      else {
        alert("You have successfully deleted your listing.")
        window.location.href =  '/';
      }
    },
    error: (err) => {
      console.log("err")
      console.log(err)
      alert("fail")
    },
    dataType: "json"
  })
}

$(() => {
  console.log(listing_price)
  console.log(listing_price)
  console.log(listing_price)

  $("#lst-price").text(formatPrice(listing_price));
})

