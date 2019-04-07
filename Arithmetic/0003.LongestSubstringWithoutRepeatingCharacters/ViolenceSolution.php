<?php

//暴力解法的耗时太长,本地测试花了12S
function answer1(String $s) : int
{
    $len = strlen($s);
    $start = 0;
    $end = 1;
    $maxLen = 1;

    if($len == 0){
        return 0;
    }

    for($i = 0; $i < $len - 1; $i++){
            for($j = $i + 1; $j < $len; $j ++){
                if(check($s, $i, $j)){
                    $maxLen = max($maxLen, $j - $i + 1);
                }else{
                    break;
                }
            }
    }

   return $maxLen;
}

function check($s, $start, $end)
{
    for($i = $start; $i <= $end ; $i ++)
    {
        if(isset($arr[$s[$i]])){
            return false;
        }
        $arr[$s[$i]] = true;
    }
    unset($arr);
    return true;
}

$s = file_get_contents('string.txt');

echo answer1($s);