
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

$(() => {
  console.log(listing_price)
  console.log(listing_price)
  console.log(listing_price)

  $("#lst-price").text(formatPrice(listing_price));
})