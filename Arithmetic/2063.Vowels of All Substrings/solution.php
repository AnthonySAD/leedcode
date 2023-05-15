<?php

function countVowels($word) {
    $len = strlen($word);
    $times = 0;
    for($i = 0; $i < $len; $i ++){
        if(in_array($word[$i], ['a', 'e', 'i', 'o', 'u'])){
            $times += ($len - $i) * ($i + 1);
        }
    }
    return $times;
}