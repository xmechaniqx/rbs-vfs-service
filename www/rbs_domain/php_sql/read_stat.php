<?php
            //Подключаем скрипт и воспользуемся уже объявленными в нём переменными чтобы не повторять их ввод для соединения с базой
    include 'put_stat.php';
    $conn = new mysqli($host, $user, $pass, $db_name);
        try {
            //Зачитываем содержимое
            $letRead="SELECT * FROM $db_table"; 
            $readResult = $conn->query($letRead);
            //Проверяем ошибки
        }catch(Exception $e){
            echo $e;
        }
        if ($readResult->num_rows > 0) {
            //Собираем таблицу используя полученные данные
            echo '<table> <tr> <th> Id </th> <th> Directory </th> <th> Size </th> <th> Elapsed time </th> <th> Date </th></tr>';
            while($row = $readResult->fetch_assoc()) {
                echo '<tr><td>' . $row["id"]. '</td>
                <td>' . $row["Directory"]. '</td>
                <td>' . $row["Size"]. '</td>
                <td>' . $row["Elapsed_time"]. '</td>
                <td>' . $row["Date"]. '</td> </tr>';
                "<br>";
            }
        } else {
            echo "0 results";
        }
    $link->close();
?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>VFS Database</title>
    <div align="right"> <a href="http://192.168.81.46/graphic.php" >График вычислений</a> </div>
    <style>
        table{
            width: 50%;
            margin: auto;
            font-family: Arial, Helvetica, sans-serif;
            font-size: 12px;
        }
        table, tr, th, td{
            border: 1px solid #d4d4d4;
            border-collapse: collapse;
            padding: 4px;
        }
        th, td{
            text-align: left;
            vertical-align: top;
        }
        tr:nth-child(even){
            background-color: #e7e9eb;
        }
    </style>
<body>