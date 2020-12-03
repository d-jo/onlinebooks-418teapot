function fill() {
    document.getElementById('title').value="moo"
    document.getElementById('seller_name').value="moo"
    document.getElementById('description').value="moo"
    document.getElementById('isbn').value="moo"
    document.getElementById('price').value="moo"
    document.getElementById('category').value="moo"
    document.getElementById('listing_password').value="moo"

    let listing_password = $("#listing_password").val();
    
    // Create object
    let dataObj = {
        listing_password: listing_password
    } 
    
    $.ajax({
        type: "POST",
        url: "/update",
        data: JSON.stringify(dataObj),
        success: (o) => {
        console.log("succ");
        console.log(o);
        let buyer = `<p>Buyer: ${o['buyer']}</p>`
        // let billingInfo = `<p>Billing Info: ${o['billing_info']}</p>`
        // let shippingInfo = `<p>Shipping Info: ${o['shipping_info']}</p>`
        // $("#details-div").append(buyer);
        // $("#details-div").append(billingInfo);
        // $("#details-div").append(shippingInfo);

        document.getElementById('title').value=`${o['title']}`
        document.getElementById('seller_name').value= buyer
        
        },
        error: (err) => {
        console.log("err")
        console.log(err)
        alert("fail")
        },
        dataType: "json"
    })

    
}

function validate_isbn(sampleISBN) {
  return isbn_re.test(sampleISBN);
}

function update_form() {
    // get form data
    let title = $("#title").val();
    let description = $("#description").val();
    let isbn = $("#isbn").val();
    let price = $("#price").val();
    let category = $("#category").val();
    let seller_name = $("#seller_name").val();
    
  
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
      seller_name: seller_name
    }
  
    console.log(dataObj)
    $.ajax({
      type: "POST",
      url: "/update",
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
  
  