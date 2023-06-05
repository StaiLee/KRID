//check if filesize is too big (>20Mb) to inform the user if his file is valid
document.getElementById('file').onchange = function(e) {
    if (!window.FileReader) { // This is VERY unlikely, browser support is near-universal
        console.log("The file API isn't supported on this browser yet.");
        return;
    }
    var input = document.getElementById('file');
    var file = input.files[0];
    if (file.size > 2000000) {
        document.getElementById('txt').innerHTML = "Your image is too big ", file.size / 1000000, "Mb";
    } else {
        document.getElementById('txt').innerHTML = "File accepted !", file.size / 1000000, "Mb";
    }
};