<?php

function strStr1($haystack, $needle) {
    $lenh = strlen($haystack);
    $lenn = strlen($needle);
    for($i = 0; $i <= $lenh - $lenn; $i ++){
        $mark = true;
        for($j = 0; $j < $lenn; $j ++){
            if($haystack[$i + $j] != $needle[$j]){
                $mark = false;
                break;
            }
        }

        if($mark)
        {
            return $i;
        }
    }
    return -1;
}


function strStr2($haystack, $needle) {
    $lenh = strlen($haystack);
    $lenn = strlen($needle);

    $next = [];
    $next[0] = -1;
    $k = -1;
    for($i = 0; $i < $lenn; $i ++){
        while($k >= 0 && $needle[$k + 1] != $needle[$i]){
            $k = $next[$k];
        }

        if($needle[$k + 1] == $needle[$i]){
            $k ++;
        }

        $next[$i] = $k;
    }

    var_dump($next);
    die;
    $j = 0;
    for($i = 0; $i <= $lenh; $i ++){
        while($j > 0 && $needle[$j] != $haystack[$i]){
            $j = $next[$j - 1] + 1;
        }

        if($haystack[$i] == $needle[$j]){
            $j ++;
        }

        if($j == $lenn){
            return $i - $j + 1;
        }

    }
    return -1;

}

var_dump(strStr2("mississippi","issip"));