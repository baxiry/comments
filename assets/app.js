// js comments application




function postData() {
    let elem = document.getElementById("post")
    elem.addEventListener("click", postData);

    let comment = document.querySelector("textarea")

    let formData = new FormData();

    formData.append('userid', 1);
    formData.append('parentid', 0);
    formData.append('comment', comment.value);


    fetch('http://localhost:1323/api', {
        method: "POST",
        body: formData,
    })
    .then(res => console.log(res.status))
    .catch(function(error) {
          console.log("Error getting document:", error);
    });
}



