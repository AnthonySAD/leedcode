<?php

function trap($height){
    $leftHeight = [];
    $res = 0;
    foreach($height as $k => $h){
        //还未求出左侧最高位
        if(count($leftHeight) < 2){
            if(!$leftHeight || $h >= $height[$leftHeight[0]]){
                $leftHeight = [$k];
            }else{
                array_unshift($leftHeight, $k);
            }
        }elseif(count($leftHeight) >= 2){
            if($h < $height[$leftHeight[0]]){
                array_unshift($leftHeight, $k);
            }elseif($h == $height[$leftHeight[0]]){
                $leftHeight[0] = $k;
            }elseif($h > $height[$leftHeight[0]]){
                //变成右侧了
                
                while(count($leftHeight) >= 2){
                    // if($k == 7){
                    //     var_dump($leftHeight);
                    // }
                    if($height[$leftHeight[0]] < $h){
                        $res += ($k - $leftHeight[1] - 1) * (min($h, $height[$leftHeight[1]]) - $height[$leftHeight[0]]);
                        array_shift($leftHeight);
                    }else{
                        break;
                    }
                    
                }

                if($h == $height[$leftHeight[0]]){
                    $leftHeight[0] = $k;
                }elseif($h < $height[$leftHeight[0]]){
                    array_unshift($leftHeight, $k);
                }else{
                    $leftHeight = [$k];
                }
            
            }
        }

    }
    return $res;
}

var_dump(test([9,6,8,8,5,6,3]));
