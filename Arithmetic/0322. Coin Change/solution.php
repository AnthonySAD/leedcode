<?php


function coinChange($coins, $amount) {
    $minCount = PHP_INT_MAX;
    getQuantity($coins, $amount, 0, $minCount);
    if($minCount == PHP_INT_MAX){
        return -1;
    }else{
        return $minCount;
    }
}

function getQuantity(&$coins, $amount, $count, &$minCount){
    if($amount < 0){
        return;
    }

    if($amount == 0){
        $minCount = min($minCount, $count);
    }

    foreach($coins as $coin){
        getQuantity($coins, $amount - $coin, $count + 1, $minCount);
    }
}


function coinChange2($coins, $amount) {
    $minCounts = array_fill(0, $amount + 1, -1);
    $minCounts[0] = 0;
    foreach($coins as $coin){
        $minCounts[$coin] = 1;
    }
    getQuantity2($coins, $amount, $minCounts);
    if($minCounts[$amount] == PHP_INT_MAX){
        return -1;
    }else{
        return $minCounts[$amount];
    }
}

function getQuantity2(&$coins, $amount, &$minCounts){

    if($minCounts[$amount] != -1){
        return $minCounts[$amount];
    }

    $tmpMin = [];
    foreach($coins as $coin){
        if($coin < $amount ){
            $tmpMin[] = getQuantity2($coins, $amount - $coin, $minCounts);
        }
    }

    if($tmpMin){
        $minCounts[$amount] = min($tmpMin) + 1;
    }else{
        $minCounts[$amount] = PHP_INT_MAX;
    }
    return $minCounts[$amount];
}


function coinChange3($coins, $amount) {
    if($amount == 0){
        return 0;
    }
    $minCounts = [];
    
    for($i = 1; $i <= $amount; $i ++){
        $min = PHP_INT_MAX;
        foreach($coins as $coin){
            if($coin < $i){
                $min = min($min, $minCounts[$i-$coin] + 1);
            }elseif($coin == $i){
                $min = 1;
                break;
            }

        }
        $minCounts[$i] = $min;
    }
    if($minCounts[$amount] == PHP_INT_MAX){
        return -1;
    }else{
        return $minCounts[$amount];
    }
}


var_dump(coinChange3([1, 2, 5], 11));