<!DOCTYPE html>
<html>
  <head>
    <link rel="stylesheet" href="//cdn.jsdelivr.net/chartist.js/latest/chartist.min.css">
    <script src="//cdn.jsdelivr.net/chartist.js/latest/chartist.min.js"></script>
    <title>VFS Line graph</title>
  </head>
  <body>
    <div style="width: 100%; height: 30px; text-align: center">  
        <span style="float: left">Размер объекта, Мбайт</span>
        Время загрузки, мс
    </div>
    <div class="ct-chart ct-golden-section" id="chart1"></div>
    <div>

<?php
            //Реализуем логику чтения из БД используя PHP
            //header('Content-Type: application/json');
    include('put_stat.php');
    $conn = new mysqli($host, $user, $pass, $db_name);
        try {
            $letReadSQL="SELECT * FROM $db_table"; 
            $readResult = $conn->query($letReadSQL);
        }catch(Exception $e){
            echo $e;
        }
    $resElapsed["Elapsed_time"] = array();
    $resSize["Size"] = array();
        if ($readResult->num_rows > 0) {
            while($row = $readResult->fetch_assoc()) {
                $resResult[$row['Size']]=$row["Elapsed_time"];
            };
        }
?>
    </div>
    <div>
    <script>
        var options = {
            width: 1600,
            height: 800
        };
            //Объявляем массив и заполняем данными через декодирование из JSON
        var phpToJsArr = <?php echo json_encode($resResult); ?>;
        var labels = new Array();
        var series = new Array();
            //Разбиваем полученный массив на два для сопоставления данных на графике
        for (var key in phpToJsArr) {
            labels.push(key);
            series.push(phpToJsArr[key]/1024/1024);
        }
        new Chartist.Line('#chart1', {
            labels: [labels],
            series: [series],
        },options);
    </script>
    </div>
    </body>
</html>