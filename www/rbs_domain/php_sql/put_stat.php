<?php
      //Делаем POST запрос
   $jsonData = file_get_contents('php://input');
      //Декодируем JSON в PHP
   $data = json_decode($jsonData, true);
      $host = 'localhost';  // Хост, у нас все локально http://192.168.81.46/
      $user = 'admin';    // Имя созданного пользователя в БД
      $pass = 'password'; // Пароль пользователя для БД
      $db_name = 't_stat';   // Имя базы данных
      $db_table ="t_stat";   // Таблица имеет такое же имя как и сама база
    
   $link = mysqli_connect($host, $user, $pass, $db_name); // Соединяемся с базой
      //Проверяем на ошибку
      if (!$link) {
         echo 'Не могу соединиться с БД. Код ошибки: ' . mysqli_connect_errno() . ', ошибка: ' . mysqli_connect_error();
      }
      //Проверяем успех декодирования
      if ($data !== null) {
      //Получаем данные и проводим нобходимые операции
      $dir=$data['root'];
      $size=$data['mainsize'];
      $elapsedTime=$data['time'];
      $dateOfProcess=$data['data'];
      //Создадим исключение для проверки на ошибку при записи данных
      try {
         $sql = "INSERT INTO t_stat (Directory, Size, Elapsed_time, Date) VALUES ('$dir', '$size', '$elapsedTime', '$dateOfProcess')";
         $result = $link->query($sql);
      }catch(Exception $e){
      echo $e;
      }
      if ($result) {
      //echo '<p>Данные успешно добавлены в таблицу.</p>';  - оставим короткий вариант обозначения успешного статуса
         echo "success";
      } else {
         echo '<p>Произошла ошибка: ' . mysqli_error($link) . '</p>';
      }
      }
?>