<?php
function maxProfit($prices) {
    $n = count($prices);

    $smallest = 0;

    for($i = 1; $i < $n; $i ++){
        if($prices[$i] <= $prices[$smallest]){
            $smallest = $i;
        }else{
            break;
        }
    }

    if($smallest == $n - 1){
        return 0;
    }
    $tmp = array_fill(0, $smallest + 1, 0);

    //从smallest开始计算
    for($i = $smallest + 1; $i < $n; $i ++){
        //计算从第j天买入后赚的钱+第j天前买入赚的最多的钱
        $money = 0;
        for($j = $smallest; $j < $i; $j++){
            $max = getMaxPrice($prices, $j + 1, $i) - $prices[$j];
            if($j > 0){
                $max = max(0, $max) + $tmp[$j - 1];
            }
            $money = max($money, $max);
        }

        $tmp[$i] = $money;
    }
    return $tmp[$n  - 1];
}

function getMaxPrice(&$prices, $start, $end)
{
    $max = 0;
    for($i = $start; $i <= $end; $i ++){
        $max = max($max, $prices[$i]);
    }
    return $max;
}

//假设每天买入卖出，则求出每天能赚的最多的钱就行。
function maxProfit2($prices) {
    $n = count($prices);

    $money = 0;

    for($i = 1; $i < $n; $i ++){
        if($prices[$i] > $prices[$i - 1]){
            $money += $prices[$i] - $prices[$i - 1];
        }
    }

    return $money;
   
}



var_dump(maxProfit2( [7,1,5,3,6,4]));