console.log("JS loaded")

var url = window.location.href + 'flag?root=/home/username/Загрузки';
// console.log("url is",url)
var currentLocation = window.location.href;


function renderResponse() {
  const resp = JSON.parse(this.response)
  const ul = document.getElementById("result")
  console.log(ul)
  resp.VFSNode_struct.forEach(element => {
    console.log(element);
    const li = document.createElement("li")
    console.log(li)
    //li.innerHTML = element.path
    li.innerHTML = '<a class="#"><div class="results">'+element.path +'</div></a>'
    
    ul.appendChild(li);
  });
}

// console.log(Resp,"2")
// document.querySelector('.results').innerHTML = a;

// console.log(time)
var req = new XMLHttpRequest();
req.addEventListener("load", renderResponse);
req.open("GET", url);
req.send();

function circle(){
  let i =0
  for (i=0;i<=length(resp.VFSNode_struct);i++){
  }
}

// document.body.appendChild(document.createTextNode(GetJSON()));
// function produceMessage(){
//   var msg= 'Hello<br />';
//   return msg;
// }
function render(){
  dirList=[]



}