<?php
function longestArithSeqLength($nums) {
    $len = count($nums);
    if($len <= 2){
        return $len;
    }

    $minCount = [];
    for($i = 0; $i< $len - 1; $i ++){
        for($j = $i + 1; $j < $len; $j ++){
            $minus = $nums[$j] - $nums[$i];
            if(isset($minCount[$minus][$i])){
                $minCount[$minus][$j] = $minCount[$minus][$i] + 1;
                unset($minCount[$minus][$i]);
            }else{
                $minCount[$minus][$j] = 2;
            }
        }
    }
    $max = 2;
    var_dump($minCount);
    foreach($minCount as $child){
        $max = max(max($child), $max);
    }

    return $max;
}


var_dump(longestArithSeqLength([20,1,15,3,10,5,8]));
