//password form checking
var check = function() {
    if (document.getElementById('password').value ==
      "" ){
      document.getElementById('message').style.color = 'red';
      document.getElementById('message').innerHTML = '<b>You need to fill in your Password !</b>';
    } else if (document.getElementById('email').value == "") {
      document.getElementById('message').style.color = 'red';
      document.getElementById('message').innerHTML = '<b>You need to fill in your Email !</b>';
    } else {
        document.getElementById('message').style.color = 'green';
        document.getElementById('message').innerHTML = '<b>All good !</b>';
        document.getElementById("submit").disabled = false;
    }
  }