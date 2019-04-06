<?php

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */

/**
 * @param ListNode $l1
 * @param ListNode $l2
 * @return ListNode
 */
function addTwoNumbers($l1, $l2) {
    
    $num1Str = '';
    $num2Str = '';
    
    while($l1){
        $num1Str = $l1->val . $num1Str ;
        $l1 = $l1->next;
    }
    
    while($l2){
        $num2Str = $l2->val . $num2Str;
        $l2 = $l2->next;
    }
    
    $sum = bcadd($num1Str, $num2Str);
    // $sum = $num1Str + $num2Str;
    
    $sumStr = strrev($sum);
    
    $sumCount = strlen($sumStr);
    
    for($i = 0; $i < $sumCount; $i ++){
        $sumList[$i] = new ListNode($sumStr[$i]);
        if($sum > 0){
            $sumList[$i - 1]->next = $sumList[$i];
        }
    }
    return $sumList[0];
}