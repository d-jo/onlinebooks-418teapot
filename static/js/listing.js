
function formatPrice(price) {
  return "$" + price.toFixed(2);
}

$(() => {
  console.log(listing_price)
  console.log(listing_price)
  console.log(listing_price)

  $("#lst-price").text(formatPrice(listing_price));
})