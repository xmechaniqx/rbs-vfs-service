"use strict";
console.log("JS loaded");
import {fnRequest} from "./request.js"
//Задаём корневой путь
var defaultURL = window.location.href + 'flag?root=/home/username/node_modules/';
//Объявляем переменную для последующий записи путей возврата (кнопка "Назад")
fnRequest(defaultURL);
