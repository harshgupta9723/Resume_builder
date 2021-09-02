var name = document.getElementById("Name").value;
var phoneNumber = document.getElementById("phone_number").value;
var email = document.getElementById("email").value;
var option1 = document.getElementById("option_1").value;
var option2 = document.getElementById("option_2").value;
var instruction = document.getElementById("instruction").value;

const obj = {userName: name,
            phone: phoneNumber,
            email: email, 
            option1: option1, 
            option2: option2, 
            instruction: instruction};

const myJSON = JSON.stringify(obj);

var data = new FormData();
data.append("myForm", myJSON);

var xhr = new XMLHttpRequest();
xhr.withCredentials = true;

xhr.addEventListener("readystatechange", function() {
if(this.readyState == 4) {
    console.log(this.responseText);
}
});

xhr.open("POST", "http://127.0.0.1:5030/resume");

xhr.send(data);