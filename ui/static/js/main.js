console.log("JS loaded")
// var url = window.location.href + 'flag?root=C:/workspace/';
var defaultURL = window.location.href + 'flag?root=/home/username/Загрузки';
var currentLocation = window.location.href;

FnRequest(defaultURL)
// renderResponse()
function FnRequest(url){
  var req = new XMLHttpRequest(url);
  req.addEventListener("load", renderResponse);
  req.open("GET", url);
  req.send();
}

function renderResponse() {
  const resp = JSON.parse(this.response)  
  const ul = document.getElementById("result")
  ul.innerHTML = ''
  resp.VFSNode_struct.forEach(element => {
  const li = document.createElement("li")


  
  li.innerHTML = '<span ><div class="results"><img src="/static/img/folder.png" width="1%">'+element.path +'</div></span>';

      let newUrl=window.location.href + 'flag?root='+element.path+'/'

li.addEventListener("click", () => {
  FnRequest(newUrl)
},  false,)    
  ul.appendChild(li);
  });
}


const goBack=document.getElementById("goback").addEventListener("click", () => {
  // const li = document.createElement("li")
  // alert("puk")
  console.log("catch a click")

  
  // }
  
},  false,) 