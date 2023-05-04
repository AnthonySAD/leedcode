<?php
function simplifyPath($path) {
    $pathArr = explode('/', $path);
    $finalPathArr = [];
    foreach($pathArr as $value){
        if($value == '..'){
            if($finalPathArr){
                array_pop($finalPathArr);
            }
        }elseif($value && $value != '.'){
            $finalPathArr[] = $value;
        }

    }

    return '/'.implode('/', $finalPathArr);
}

var_dump(simplifyPath('/a/./b/../../c/'));