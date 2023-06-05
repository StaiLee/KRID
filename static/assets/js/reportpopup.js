var openn = 1
function openform(i) {
    document.getElementById(i).style.display = "block";
    if (openn == 1) {
        exitreport(i)
    }
    if (openn == 0) {
        openn = 1
    } else {
        openn = 0
    }
}

function exitreport(i) {
    document.getElementById(i).style.display = "none";
}