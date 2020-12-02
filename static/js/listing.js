
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
      console.log("success");
      console.log(o);
      alert("The listing has been purchased");
      location.reload();
    },
    error: (err) => {
      console.log("err")
      console.log(err)
      alert("Unexptected server error")
    },
    dataType: "json"
  })

}

function get_private_details() {
  let listing_password = $("#listing_password").val();

  // Create object
  let dataObj = {
    listing_password: listing_password
  } 

  $.ajax({
    type: "POST",
    url: "/listing/private_details/" + listing_id,
    data: JSON.stringify(dataObj),
    success: (o) => {
      console.log("succ");
      console.log(o);
      let buyer = `<p>Buyer: ${o['buyer']}</p>`
      let billingInfo = `<p>Billing Info: ${o['billing_info']}</p>`
      let shippingInfo = `<p>Shipping Info: ${o['shipping_info']}</p>`
      $("#details-div").append(buyer);
      $("#details-div").append(billingInfo);
      $("#details-div").append(shippingInfo);
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

  if (status == "purchased") {
    $("#buyer_name").remove();
    $("#buyer_billing").remove();
    $("#buyer_shipping").remove();
    $("#purchase-btn").remove();

  }
})

