
let isbn_re = new RegExp('^[0-9][0-9]?[0-9]?\-[0-9][0-9]?[0-9]?[0-9]?[0-9]?[0-9]?[0-9]?\-[0-9][0-9]?[0-9]?[0-9]?[0-9]?[0-9]?\-[0-9][0-9]?[0-9]?[0-9]?[0-9]?[0-9]?\-?[0-9]?$');

function validate_isbn(sampleISBN) {
  return isbn_re.test(sampleISBN);
}

function error_message(message) {
  alert(message);
}

function submit_form() {
  // get form data
  let title = $("#title").val();
  let description = $("#description").val();
  let isbn = $("#isbn").val();
  let price = $("#price").val();
  let category = $("#category").val();
  let seller_name = $("#seller_name").val();
  let listing_password = $("#listing_password").val();

  let errorMessage = "";

  // VERIFY HERE
  if (title.length < 6) {
    // error
    errorMessage += "Title must be 6 characters minimum\n";
  } 
  if (description.length < 10) {
    // error
    errorMessage += "Description must be 10 characters\n";
  } 
  if (!validate_isbn(isbn)) {
    // error
    errorMessage += "Invalid ISBN, must include dashes\n";
  } 
  if (Number.isNaN(parseFloat(price))) {
    // error
    errorMessage += "Price must be a float\n";
  } 
  if (category.length < 3) {
    // error
    errorMessage += "Invalid Category\n";
  } 
  if (seller_name.length < 2) {
    // error
    errorMessage += "Seller name must be at least 2 characters\n";
  } 
  if (listing_password.length < 12) {
    // error
    errorMessage += "Password must be at least 12 characters\n"
  }

  console.log(errorMessage);

  if (errorMessage.length > 1) {
    alert(errorMessage);
    return;
  }
  
  price = parseFloat(price);

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
  $.ajax({
    type: "POST",
    url: "/create_listing",
    data: JSON.stringify(dataObj),
    success: (o) => {
      console.log("succ")
      console.log(o);
      if (Number.isInteger(o)) {
        window.location.href = "/listing/" + o;
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