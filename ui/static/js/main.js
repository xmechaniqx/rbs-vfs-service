console.log("JS loaded")
var url = window.location.href + 'flag?root=C:/workspace/';
// var url = window.location.href + 'flag?root=/home/username/Загрузки';
console.log("url is",url)
var currentLocation = window.location.href;

 function FnRequest(){
  var req = new XMLHttpRequest();
req.addEventListener("load", renderResponse);
req.open("GET", url);
req.send();
}
FnRequest()


function renderResponse() {
  const resp = JSON.parse(this.response)  
  const ul = document.getElementById("result")
  console.log(ul)
  resp.VFSNode_struct.forEach(element => {
    console.log("element",element.path);
    const li = document.createElement("li")
    console.log(li)

    // console.log(element.stat)
// FnRequest()
    // if  (String(element.stat).valueOf() == String("file").valueOf()) {
     
      // li.innerHTML = '<a class="#"><div class="results"><img src="/static/img/file.png'+element.path +'</div></a>' 
    // } else if (String(element.stat).valueOf() == String("dir").valueOf()){
      let newUrl=element.path
      
      
      console.log(newUrl.replace('ForwardSlash',/\//g)+'/')


      li.innerHTML = '<a class="#" href='+newUrl+'><div class="results">'+element.path +'</div></a>' ;
    // }
    

    // renderResponse()
    // li.innerHTML = '<a class="#"><div class="results">'+element.path +'</div></a>'  
    ul.appendChild(li);

  });
}

// renderResponse()
function click(){
  const ul = document.getElementById("click")
  document.addEventListener("mouseup",
                ()=>console.log("mouseup"));
                url= "123"
                console.log(url)
                FnRequest(url)
                renderResponse(url)
}