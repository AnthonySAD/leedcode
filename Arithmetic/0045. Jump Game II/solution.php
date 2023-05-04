<?php

function jump($nums) {
    $count = count($nums);
    $arr[0] = 0;
    so($nums, $count - 1, $arr);

    return $arr[$count - 1];
}

function so(&$nums, $maxN, &$tmp){
    if(isset($tmp[$maxN])){
        return $tmp[$maxN];
    }

    if($maxN <= $nums[0]){
        $tmp[$maxN] = 1;
        return $tmp[$maxN];
    }

    $min = 20000;

    for($i = 0; $i < $maxN; $i ++){
        if($i + $nums[$i] >= $maxN){
            $tmpMin = so($nums, $i, $tmp) + 1;
            if($tmpMin < $min){
                $min = $tmpMin;
            }
        }
    }
    $tmp[$maxN] = $min;
    return $tmp[$maxN];
}

//这思路6啊
function jump2($nums) {
    $n = count($nums);
    $start = 0;
    $end = 0;
    $times = 0;

    while($end < $n - 1){
        $tmpEnd = $end;
        for($i = $start; $i <= $end; $i ++){
            $tmpEnd = max($tmpEnd, $i + $nums[$i]);
        }
        $start = $end;
        $end = $tmpEnd;
        $times ++;
    }
    return $times;

}

var_dump(jump2([2,3,0,1,4]));