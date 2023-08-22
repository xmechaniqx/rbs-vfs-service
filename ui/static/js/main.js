console.log("JS loaded")

var url = window.location.href + 'flag?root=/home/username/Загрузки';
console.log("url is",url)
var currentLocation = window.location.href;

  // function FnRequest(){
  var req = new XMLHttpRequest();
req.addEventListener("load", renderResponse);
req.open("GET", url);
req.send();
// }
function renderResponse() {
  const resp = JSON.parse(this.response)  
  const ul = document.getElementById("result")
  console.log(ul)
  resp.VFSNode_struct.forEach(element => {
    console.log(element);
    const li = document.createElement("li")
    console.log(li)
    // console.log(element.stat)

    // if  (String(element.stat).valueOf() == String("file").valueOf()) {
     
      // li.innerHTML = '<a class="#"><div class="results"><img src="/static/img/file.png'+element.path +'</div></a>' 
    // } else if (String(element.stat).valueOf() == String("dir").valueOf()){
      li.innerHTML = '<a class="#"><div class="results">'+element.path +'</div></a>' 
    // }
    // li.innerHTML = '<a class="#"><div class="results">'+element.path +'</div></a>'  
    ul.appendChild(li);
  });
}