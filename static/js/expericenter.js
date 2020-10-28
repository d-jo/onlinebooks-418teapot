
var current_screen = ".identification";
var user_identifier = null;
var session_identifier = null;
var sample_id = null;
var phrases_from_server = null;

var current_phrase_set = 0;
var current_phrase = 0;
var free_response = false;
var free_response_remaining = 0;

var buffer_submit_timer = null;
var current_buffer = false;
var keyevent_buffer_true = [];
var keyevent_buffer_false = [];

function buildRequest(buff) {
  req_body = {
    user_id: user_identifier,
    session_id: session_identifier,
    sample_id:  sample_id,
    keystroke_events: buff
  }
  return req_body;
}

function postRequest(buff) {
  req_body = buildRequest(buff);
  console.log("Posting buffer");
  //console.log(req_body);
  $.ajax({
    type: "POST",
    url: "/keystroke_data",
    data: JSON.stringify(req_body),
    success: () => {
      console.log("Success")
    },
    dataType: "json"
  })
}

function go() {
  //console.log(current_screen)
  if (current_screen === ".identification") {
    checkIdentification();
  } else if (current_screen === ".instructions") {
    acceptInstructions();
  } else if (current_screen === ".prep_phrase") {
    showTypeScreen();
  } else if (current_screen === ".type_phrase") {
    submitPhrase();
  } else if (current_screen === ".dual_phrase") {
    submitPhrase();
  } else if (current_screen === ".free_response") {
    finishResponses();
  }
}

function enterListener(e) {
  original_event = e.originalEvent;
  if (original_event.key === "Enter") {
    go();
  }
  
}

function keyListener(e) {
  original_event = e.originalEvent;
  kc = original_event.keyCode;
  if (!e.shiftKey && kc >= 65 && kc <= 122) {
    kc += 32
  }
  //console.log(original_event);
  if (current_buffer) {
    keyevent_buffer_true.push({
      "time": new Date().getTime(), 
      "keycode": kc, 
      "key": original_event.key, 
      "event_type": original_event.type
    })

  } else {
    keyevent_buffer_false.push({
      "time": new Date().getTime(), 
      "keycode": kc, 
      "key": original_event.key, 
      "event_type": original_event.type
    })
  }
}

function postAllBuffers() {
  false_buff_copy = keyevent_buffer_false.slice();
  true_buff_copy = keyevent_buffer_true.slice();
  // send the data
  if (false_buff_copy.length > 0) {
    postRequest(false_buff_copy);
  }
  if (true_buff_copy.length > 0) {
    postRequest(true_buff_copy);
  }
  // clear the buffer
  false_buff_copy = [];
  true_buff_copy = [];
  keyevent_buffer_false = [];
  keyevent_buffer_true = [];
}

function switchBufferAndSend() {
  // switch what buffer we are using
  buff_len = null;
  if (current_buffer) {
    buff_len = keyevent_buffer_true.length;
  } else {
    buff_len = keyevent_buffer_false.length;
  }
  if (buff_len > 0) {
    current_buffer = !current_buffer;
    if (current_buffer) {
      // if thee current buffer is true, we
      // should empty the false buffer
      // the keystroke data doesnt need to get there in order

      // first make a copy of the buffer
      false_buff_copy = keyevent_buffer_false.slice();
      // send the data
      postRequest(false_buff_copy);
      // clear the buffer
      false_buff_copy = [];
      keyevent_buffer_false = [];
    } else {
      // if thee current buffer is false, we
      // should empty the true buffer
      // first make a copy of the buffer
      true_buff_copy = keyevent_buffer_true.slice();
      // send the data
      postRequest(true_buff_copy);
      // clear the buffer
      true_buff_copy = [];
      keyevent_buffer_true = [];
    }
  }
}

function addFullListener(target, callback) {
  $(target).keyup(callback)
  $(target).keydown(callback)
}

function fadeTo(scr_class, intermediate_callback, final_callback) {
  $(current_screen).fadeOut("fast", () => {
    if (intermediate_callback != null) {
      intermediate_callback()
    }
    $(scr_class).fadeIn("fast", final_callback ? final_callback : ()=>{console.log("skip final callback")});
  })
  current_screen = scr_class
}

// moves thes state to instructions
function toInstructions() {
  // verify state is valid to move to instructions
  if (current_screen === ".identification" && user_identifier != null && session_identifier != null) {
    // valid, move
    fadeTo(".instructions", null, () => {
      $("#identifier_confirm").focus();
    })
    $(".instructions").css("display", "flex");
    //addFullListener("#identifier_confirm", keyListener)
    addFullListener("#phrase_input", keyListener);
    sample_id=-1;
    buffer_submit_timer = setInterval(switchBufferAndSend, 6282);
    //$("#identifier_confirm").focus();
    //$("#identifier_confirm").keyup(keyListener)
    //$("#identifier_confirm").keydown(keyListener)
  } else {
    // invalid, do not move
    alert("Invalid state. refresh page");
  }
}

function checkIdentification() {
  let ident = $("#identifier").val().trim()
  if (ident.length < 3) {
    // fail, needs to be > 3
    alert("Identifier must be at least 3 characters");
  } else {
    // success, set local variable
    user_identifier = ident;
    session_identifier = new Date().getTime() + "_" + user_identifier;
    // advance screen
    toInstructions();
  }
}

function acceptInstructions() {
  let ident = $("#identifier_confirm").val().trim()
  if (ident.length < 3) {
    // fail, needs to be > 3
    alert("Name must be at least 3 characters");
  } else {
    if (ident === user_identifier) {
      // move to main loop
      showPhraseScreen()
    }
  }
}

function showPhraseScreen() {
  phr = phrases_from_server[current_phrase_set][current_phrase]
  // set display phrase, phrase_id
  sample_id = phr["id"];
  $("#displayed_phrase").text(phr["phrase"]);
  // show display phrase screen
  fadeTo(".prep_phrase", null, null);
  //$(".prep_phrase").css("display", "flex");
}

function showTypeScreen() {
  $("#phrase_input").val("");
  fadeTo(".type_phrase", null, () => {
    $("#phrase_input").focus();
  });
  //$(".type_phrase").css("display", "flex");
}

function showDualScreen() {
  f = () => {
    $("#phrase_input_dual").val('');
    phr = phrases_from_server[current_phrase_set][current_phrase] 
    sample_id = phr['id']
    $("#displayed_phrase_dual").text(phr["phrase"]);
  }
  fadeTo(".dual_phrase", f, () => {
    $("#phrase_input_dual").focus();
  });
  //$(".dual_phrase").css("display", "flex");
  $("#phrase_input_dual").focus();
}

function showFreeResponse() {
  sample_id = -2
  fadeTo(".free_response", null, () => {
    $("#phrase_input_free").focus();
  });
  //$(".free_response").css("display", "flex")
}

function submitPhrase() {
  current_phrase += 1;
  if (current_phrase === phrases_from_server[current_phrase_set].length) {
    // move to dual
    current_phrase_set += 1;
    current_phrase = 0;
  }
  if (current_phrase_set === 0) {
    showPhraseScreen();
  } else if (current_phrase_set === 1) {
    showDualScreen();
  } else {
    showFreeResponse();
  }
}

function finishResponses() {
  postAllBuffers();
  clearInterval(buffer_submit_timer);
  console.log("Timer cleared");
  fadeTo(".debrief", null, null)
}

$(() => {
  console.log("Document ready");
  $(document).keypress(enterListener);
  $("#identifier").focus();
  $.ajax({
    type: "GET",
    url: "/getphrases",
    success: (data) => {
      console.log("Data from server")
      console.log(data)
      phrases_from_server = data
    },
    error: (err) => {
      console.log('err')
      console.log(err)
    },
    dataType: "json"
  })
})