<?php

//2792 ms	15.3 MB
function twoSum($nums, $target)
{
    $len = count($nums);
    for($i = 0; $i < $len - 1; $i ++){
        for($j = $i + 1; $j < $len; $j ++){
            if($nums[$i] + $nums[$j] == $target){
                return [$i, $j];
            }
            // 使用以下方法优化后运行时间减少到2164 ms
            // $temp = $target - $nums[$i];
            // if($nums[$j] == $temp){
            //     return [$i, $j];
            // }
        }
    }
    return [];
}

//36 ms	16.2 MB
function twoSum2($nums, $target) {
    $map = array_flip($nums);
    $len = count($nums);
    for($i = 0; $i < $len; $i++) 
    {
        $temp = $target-$nums[$i];
        if(!empty($map[$temp])) 
        {
            $j = $map[$temp];
            if($i != $j){
                return [$i, $j];
            }
        }
    }
    return [];
}