module.exports={fnRequest};
//Функция fnRequest() принимает URL адрес и отправляет HTTP запрос 
function fnRequest(url) {
    drawPreloader();
    var req = new XMLHttpRequest();
    req.addEventListener("load", renderResponse);
    req.open("GET", url);
    req.send();
}
var canvas = document.getElementById('preloader');
var ctx = canvas.getContext('2d');
var angle = 0;
//Вспомогательная функция для fnRequest(), парсинг полученной страницы формата JSON и получение необходимых параметров
function renderResponse() {
    var resp = JSON.parse(this.response);
    var ul = document.getElementById('result');
    //Очистка тела UL таблицы как способ обновления страницы для отображения содержимого директорий при переходе
    ul.innerHTML = '';
    //
    resp.VFSNode_struct.forEach(function (element) {
        var li = document.createElement("li");
        li.setAttribute("vfs_path", element.path);
        if (element.stat == "dir") {
            li.innerHTML = '<span ><div class="results"><img src="/static/img/folder.png" width="1%">' + element.path + '</div></span>';
        }
        if (element.stat == "file") {
            li.innerHTML = '<span ><div class="results"><img src="/static/img/file.png" width="1%">' + element.path + '</div></span>';
        }
        //Записываем новый путь перехода по директории в переменную newUrl
        var newUrl = window.location.href + 'flag?root=' + element.path + '/';
        var root = resp.root;
        curPath(root);
        if (element.stat == "dir") {
            li.addEventListener("click", function () {
                fnRequest(newUrl);
            }, false);
        }
        ul.appendChild(li);
        backURL = removeLastDirectoryPartOf(root);
        backURL = removeLastDirectoryPartOf(backURL);
        backURL = 'flag?root=' + backURL + '/';
    });
}
//removeLastDirectoryPartOf() - Вспомогательная функция для удаления последней директории из адреса
function removeLastDirectoryPartOf(the_url) {
    var the_arr = the_url.split('/');
    the_arr.pop();
    return (the_arr.join('/'));
}
//curPath() - функция для отображения на Web-странице текущего местоположения в директории 
function curPath(param) {
    var ul = document.getElementById("current_path");
    var li = document.createElement("li");
    ul.innerHTML = '<span><div class="results"><img src="/static/img/folder.png" width="1%">' + param + '</div></span>';
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
    requestAnimationFrame(drawPreloader);
}
var backURL;

//Обрабатываем логику возврата в предыдущую директорию
var goBack = document.getElementById("goback");
if (goBack) {
    goBack.addEventListener("click", function () {
        //Если возврат ссылается на директорию выше корня, то кнопка "Назад" не сработает
        if (backURL != 'flag?root=/') {
            fnRequest(backURL);
        }
    }, false);
}