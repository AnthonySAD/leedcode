<?php

class Solution {

    private $ans = [];

    function generateParenthesis($n) {
        $this->make(0, 0 , $n, '');
        return $this->ans;
    }

    function make($left, $right, $max, $str) {
        if($right == $max){
            $this->ans[] = $str;
        }else{
            if($left < $max){
                $this->make( $left + 1, $right, $max, $str . '(');
            }

            if($right < $left){
                $this->make($left, $right + 1, $max, $str . ')');
            }
        }
    }

}

$solution = new Solution;
$ans = $solution->generateParenthesis(3);
var_dump($ans);