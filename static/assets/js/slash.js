slash = "";
function myKeyPress(e) {
    var keynum;

    if (window.event) { // IE                  
        keynum = e.keyCode;
    }
    key = String.fromCharCode(keynum);
    if (key == "/" && slash.length > 1) {
        slash+=" "
        document.getElementById('displayslash').innerHTML += slash;
        slash = "/";
        document.getElementById('slash').innerHTML = slash;
    } else {
        slash += key;
    };
}