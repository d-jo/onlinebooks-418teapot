
function submit_form() {
  // get form data
  let title = $("#title").val();
  let description = $("#description").val();
  let isbn = $("#isbn").val();
  let price = $("#price").val();
  let category = $("#category").val();
  let seller_name = $("#seller_name").val();
  let listing_password = $("#listing_password").val();

  // VERIFY HERE

  // Create object
  let dataObj = {
    title: title,
    description: description,
    isbn: isbn,
    price: price,
    category: category,
    seller_name: seller_name,
    listing_password: listing_password
  }

  console.log(dataObj)

}