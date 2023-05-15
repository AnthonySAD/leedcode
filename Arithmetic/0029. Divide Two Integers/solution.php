<?php


function divide($dividend, $divisor){
    if($dividend == 0){
        return 0;
    }
    if($divisor == 1){
        return $dividend;
    }
    if($divisor == -1){
        if($dividend > -2147483648){
            return -$dividend;
        }else{
            return 2147483647;
        }
    }

    if(($dividend < 0 && $divisor > 0) || ($dividend > 0 && $divisor < 0)){
        $sign = -1;
    }else{
        $sign = 1;
    }

    $dividend = abs($dividend);
    $divisor = abs($divisor);

    $res = so($dividend, $divisor);
    return $sign>0?$res:-$res;
}

function so($dividend, $divisor){
    if($dividend < $divisor){
        return 0;
    }
    $count = 1;
    $tmp = $divisor;
    while(($tmp + $tmp) <= $divisor){
        $count += $count;
        $tmp += $tmp;
    }
    return $count + so($dividend - $tmp, $divisor);
}


var_dump(divide(10,3));

function divideWrong($dividend, $divisor) {
    $times = 0;
    $max = 0;

    if($dividend == 0){
        return 0;
    }

    if($divisor == 0){
        if($dividend > 0){
            return PHP_INT_MAX;
        }else{
            return PHP_INT_MIN;
        }
    }

    

    while(true){
        if(($dividend < 0 && $divisor > 0) || ($dividend > 0 && $divisor < 0)){
            $max -= $divisor;
            $times --;
        }else{
            $max += $divisor;
            $times ++;
        }
        
        if($max == $dividend){
            return $times;
        }
        if($dividend > 0){
            if($max > $dividend){
                if($divisor > 0){
                    return $times - 1;
                }else{
                    return $times + 1;
                }
            }
        }else{
            if($max < $dividend){
                if($divisor > 0){
                    return $times + 1;
                }else{
                    return $times - 1;
                }
            }
        }
        echo $times.PHP_EOL;
        
    }
}