<?php

function canJump($nums) {
    $len = count($nums);
    $max = $nums[0];
    for($i = 0; $i < $len; $i++){
        if($i > $max){
            return false;
        }

        $max = max($max, $i + $nums[$i]);
        if($max >= $len - 1){
            return true; 
        }
    }
}


var_dump(canJump([2,3,1,1,4]));