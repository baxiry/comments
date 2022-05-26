// js comments application

// post comment data
function postData() {
    fmt.Pritnln()

    let area = document.querySelector("textarea")
    if (area.value == "") {
        console.log("comment is empty")
        return
    }
   
    let elem = document.getElementById("post")
    elem.addEventListener("click", postData);
	

    let formData = new FormData();

    formData.append('userid', 1);
    formData.append('parentid', 0);
    formData.append('comment', comment.value);


    fetch('http://localhost:1323/api', {
        method: "POST",
        body: formData,
    })
    .then(res => console.log(res.status))
    .catch(error => console.log("Error:", error));
    area.value = '';
}
