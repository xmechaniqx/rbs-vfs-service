console.log("JS loaded")
// var url = window.location.href + 'flag?root=C:/workspace/';
var defaultURL = window.location.href + 'flag?root=/home/username/Загрузки';
var currentLocation = window.location.href;
var BackURL

const goBack=document.getElementById("goback").addEventListener("click", () => {
  FnRequest(BackURL)
  console.log("catch a click")  
}, false) 

// var oldUrl=defaultURL
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
    if (element.stat=="dir") {
      li.innerHTML = '<span ><div class="results"><img src="/static/img/folder.png" width="2%">'+element.path +'</div></span>';
    }
    if (element.stat=="file") {
      li.innerHTML = '<span ><div class="results"><img src="/static/img/file.png" width="2%">'+element.path +'</div></span>';
    }
    let newUrl=window.location.href + 'flag?root='+element.path+'/'
    let root=resp.root ;
    // console.log("old" , oldUrl,"\n","new",newUrl)
    // oldUrl=window.location.href + 'flag?root='+root
    curPath(root)
    if (element.stat=="dir"){
    li.addEventListener("click", () => { 
      FnRequest(newUrl)
    },  false,)  
  }
    ul.appendChild(li);
    
    BackURL= RemoveLastDirectoryPartOf(root)
    BackURL= RemoveLastDirectoryPartOf(BackURL)
    BackURL='flag?root='+BackURL+'/'
  });
}

function RemoveLastDirectoryPartOf(the_url){
    var the_arr = the_url.split('/');
    the_arr.pop();
    return( the_arr.join('/') );
}

// function goBack(url){
//  const goBack=document.getElementById("goback").addEventListener("click", () => {
//   console.log("catch a click")  
//   FnRequest(url)
// },  false,) 
// }

function curPath(param){
const ul = document.getElementById("current_path")
  const li = document.createElement("li")
  ul.innerHTML = '<span ><div class="results"><img src="/static/img/folder.png" width="1%">'+param+'</div></span>'
}