<?php
    $num = intval($_GET['value']);
    $memnum = intval($_GET['memory']);

    if ($num < 1) {
        $num = intval($_POST['value']);
    }
    
    if ($memnum < 1) {
        $memnum = intval($_POST['memory']);
    }

    echo "<html>";
    echo "<head> <title> Memory Test </title> </head>";
    echo "<body> <p> <center>";
    echo "<h2> Memory Test by constructing a big string </h2>";

    if($num < 0) {
        $num = 0;
    }

    if($memnum < 1) {
        $memnum = 1;
    }
    
    $maxnum = 2048;
    if($memnum > $maxnum) {
        echo "memory usage limit: [1--$maxnum]MB. <br>";
        $memnum = $maxnum;
    }

    $rnum = $num;
    if($num > 1) {
        $rnum = rand($num, $num+3);
    }
    $rmemnum = $memnum *1024*1024 + rand(100, 1024);


    echo "sleep $rnum ms, with $memnum MB memory. <br>";

    $basexx = str_repeat("hello", 100);
    $tmp = str_repeat($basexx, $rmemnum/500);
    $usedMem = round(memory_get_usage()/(1024*1024.0), 2);
    echo "memory used 1: <b> $usedMem MB </b>.";

    usleep(1000*$rnum);

    unset($tmp);
    echo "<br><br>free memory...<br><br>";
    $usedMem = round(memory_get_usage()/(1024*1024.0), 2);
    echo "memory used 2: <b> $usedMem MB </b>.";

    echo "</center> </p>";
    include 'footer.php';
    echo "</body></html>";
?>
