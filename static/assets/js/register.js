//password form checking
var check = function() {
    document.getElementById("submit").disabled = true;
    if (document.getElementById('password').value !=
        document.getElementById('confirm_password').value) {
        document.getElementById('message').style.color = 'red';
        document.getElementById('message').innerHTML = '<b>Passwords do not match !</b>';
    } else if (document.getElementById('password').value.length < 5) {
        document.getElementById('message').style.color = 'red';
        document.getElementById('message').innerHTML = '<b>You need a longer Password !</b>';
    } else if (document.getElementById('email').value.length < 1) {
        document.getElementById('message').style.color = 'red';
        document.getElementById('message').innerHTML = '<b>You need a valid email !</b>';
    } else if (document.getElementById('identifiant').value.length < 1) {
        document.getElementById('message').style.color = 'red';
        document.getElementById('message').innerHTML = '<b>You need a valid username !</b>';
    } else {
        document.getElementById('message').style.color = 'green';
        document.getElementById('message').innerHTML = '<b>All good !</b>';
        document.getElementById("submit").disabled = false;
    }
}