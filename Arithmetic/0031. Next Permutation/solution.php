<?php

function nextPermutation($nums){
    $len = count($nums);
    $i = $len - 2;
    while($i >= 0 && $nums[$i] >= $nums[$i + 1]){
        $i --;
    }

    if($i >= 0){
        $j = $len - 1;
        while($j >= 0 && $nums[$i] >= $nums[$j]){
            $j --;
        }
        $tmp = $nums[$i];
        $nums[$i] = $nums[$j];
        $nums[$j] = $tmp;
    }
    
    //交换
    $b = $len - 1;
    for($a = $i + 1; $a < $i + ($len - $i)/2; $a ++){
        $tmp = $nums[$a];
        $nums[$a] = $nums[$b];
        $nums[$b] = $tmp;
        $b --;
    }

    return $nums;
}


var_dump(nextPermutation([1,3,2]));



