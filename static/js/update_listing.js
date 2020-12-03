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


function update_listing() {
    let listing_password = $("#listing_password").val();
  
    // Create object
    let dataObj = {
      listing_password: listing_password
    } 
  
    $.ajax({
      type: "POST",
      url: "/listing/" + listing_id + "/update",
      data: JSON.stringify(dataObj),
      success: (o) => {
        console.log("succ");
        console.log(o);

        let title = `<p>Title: ${o['title']}</p>`
        let seller = `<p>Seller Name: ${o['seller']}</p>`
        let desc = `<p>Description: ${o['desc']}</p>`
        let isbn = `<p>ISBN: ${o['isbn']}</p>`
        let price = `<p>Price: ${o['price']}</p>`
        let category = `<p>Category: ${o['category']}</p>`
        let pass = `<p>Password: ${o['pass']}</p>`
        
        $("#details-div").append(title);
        $("#details-div").append(seller);
        $("#details-div").append(desc);
        $("#details-div").append(isbn);
        $("#details-div").append(price);
        $("#details-div").append(category);
        $("#details-div").append(pass);
      },
      error: (err) => {
        console.log("err")
        console.log(err)
        alert("fail")
      },
      dataType: "json"
    })
  
  }
  