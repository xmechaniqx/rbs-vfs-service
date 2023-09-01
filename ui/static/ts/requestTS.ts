// module.exports={fnRequest};
export{fnRequest};
//Функция fnRequest() принимает URL адрес и отправляет HTTP запрос 
function fnRequest(url) {
    var req = new XMLHttpRequest();
    req.addEventListener("load", renderResponse);
    req.open("GET", url);
    req.send();
}

// document.canvas.innerHTML = '';
const canvas :HTMLElement | any= document.getElementById('preloader');
const ctx = canvas.getContext('2d');
var angle = 0;
var goBackUrl
//Вспомогательная функция для fnRequest(), парсинг полученной страницы формата JSON и получение необходимых параметров
function renderResponse(){
    const resp = JSON.parse(this.response);
    const ul :HTMLElement | any= document.getElementById('result');
    var root = resp.root; 
    curPath(root);
    //Очистка тела UL таблицы как способ обновления страницы для отображения содержимого директорий при переходе
    ul.innerHTML = '';
    //Задаем переменную goBackUrl для фиксации предыдущей директории относительно текущей
    goBackUrl = removeLastDirectoryPartOf(window.location.href + 'flag?root=' + removeLastDirectoryPartOf(root)) + '/';
    //Цикл для массива VFSNode
    resp.VFSNodeStruct.forEach(function (element) {
        var li = document.createElement("li");
        li.setAttribute("vfsPath", element.path);
        if (element.stat == "dir") {
            li.innerHTML = '<span ><div class="results"><img src="/static/img/folder.png" width="1%">' + element.path + '</div></span>';
        }
        if (element.stat == "file") {
            li.innerHTML = '<span ><div class="results"><img src="/static/img/file.png" width="1%">' + element.path + '</div></span>';
        }
        //Записываем новый путь перехода по директории в переменную newUrl
        var newUrl = window.location.href + 'flag?root=' + element.path + '/';
       //Ограничиваем попытку использовать файл как директорию
        if (element.stat != "file") {
            
            li.addEventListener("click", function () {
                
                fnRequest(newUrl);
            }, false);
        }
        ul.appendChild(li);
    });
}
//removeLastDirectoryPartOf() - Вспомогательная функция для удаления последней директории из адреса
function removeLastDirectoryPartOf(takeUrl) {
    var dirArray = takeUrl.split('/');
    dirArray.pop();
    return (dirArray.join('/'));
}
//curPath() - функция для отображения на Web-странице текущего местоположения в директории 
function curPath(param:string): void {
    const ul:HTMLElement | any= document.getElementById("current_path")
    ul.innerHTML = '<span><div>' + param + '</div></span>';
}
//drawPreloader - функция анимации загрузки 
function drawPreloader() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.beginPath();
    ctx.arc(50, 50, 40, angle, angle + Math.PI / 4);
    ctx.lineWidth = 10;
    ctx.strokeStyle = '#000000';
    ctx.stroke();
    angle += 0.1;
    requestAnimationFrame(drawPreloader)
}
//
//Обрабатываем логику возврата в предыдущую директорию
const goBack=document.getElementById("goback");
    if (goBack) {
        goBack.addEventListener("click", function () {
        //Если возврат ссылается на директорию выше корня, то кнопка "Назад" не сработает
            if (goBackUrl != 'flag?root=/') {
               const empty :HTMLElement | any= document.getElementsByClassName('body-block');
                empty.innerHTML="";
                fnRequest(goBackUrl);
            }
        }, false);
    }
//
window.onload=drawPreloader