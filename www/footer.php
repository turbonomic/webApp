<?php

    echo "<hr>";
    echo "<p><center>";
    echo "<br> version.11 <br>\n";
    echo "Host: ". $_SERVER['HTTP_HOST'];
    echo "IP: ". $_SERVER['SERVER_ADDR'];
    echo "<br>";
    echo "Container: ". gethostname() ;
    echo "<br>";
    echo "<a href='/index.html'>Home</a>";
    echo "</center></p>";
?>
