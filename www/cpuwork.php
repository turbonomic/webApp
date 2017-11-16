<?php
    $num = intval($_GET["cpu"]);

    if ($num < 1) {
        $num = intval($_POST['cpu']);
    }

    if($num < 1) {
        $num = 1;
    }

    //$low = $num;
    //if($low < 1) {
    //    $low = 1;
    //}

    echo "<html>";
    echo "<head> <title> CPU Test </title> </head>";
    echo "<body> <p> <center>";
    echo "<h2> CPU Test by calculating MD5 of numbers. </h2>";

    // $rnum = rand($low, $num) * 1000;
    $rnum = $num * 1000;
    echo "computation amount: $rnum <br>";

    $begin = microtime(true);
    for($i=0; $i < $rnum; $i ++) {
        md5($i);
    }
    $duration = intval((microtime(true) - $begin)*1000);

    echo "used $duration ms. \n";

    echo "</center> </p>";
    include 'footer.php';
    echo "</body></html>";
?>
