function jsonResume(){

    var rawResume = document.getElementById("Name").value;

    var data = new FormData();
    data.append("myForm", rawResume);

    var xhr = new XMLHttpRequest();
    xhr.withCredentials = true;

    xhr.addEventListener("readystatechange", function() {
    if(this.readyState === 4) {
        console.log(this.responseText);
    }
    });

    xhr.open("POST", "http://127.0.0.1:5030/resume");

    xhr.send(data);

}
