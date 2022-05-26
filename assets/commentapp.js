function $(element) {
    return document.querySelector(element)
}

let siteUrl = document.URL
console.log(siteUrl)
let url = 'http://localhost:1323/api?site='
let comentArea = document.createElement('iframe')
comentArea.setAttribute("src", url+siteUrl)
comentArea.setAttribute("style","width: 100%; height: 100%;")
document.body.center.appendChild(comentArea);  

