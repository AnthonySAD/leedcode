<?php

//36 ms	14.9 MB
function solution(String $s) : int
{
    $len = strlen($s);
    $start = 0;
    $maxLen = 0;

    for($end = 1, $start = 0; $end < $len; $end ++){
        for($j = $start; $j < $end; $j ++){
            if($s[$end] == $s[$j]){
                $start = $j + 1;
                break;
            }
            if($j == $end - 1){
                $maxLen = max($maxLen, $end - $start + 1);
            }
            
        }
    }

    return $maxLen;
}

$s = file_get_contents('string.txt');

echo solution($s);