<?php

class Solution {
    public $dp = [];

    function findMinStep($board, $hand) {
        $hand = str_split($hand, 1);
        sort($hand);
        $hand = implode('',$hand);

        $res = $this->dfs($board, $hand);
        return $res == 6 ? -1 : $res;
    }

    function dfs($board, $hand){
        if(strlen($board) == 0){
            return 0;
        }

        $key = $board . ' ' . $hand;
        if(isset($this->dp[$key])){
            return $this->dp[$key];
        }

        $res = 6;
        for($i = 0; $i <= strlen($board); $i ++){
            for($j = 0; $j < strlen($hand); $j ++){
                //如果手中的球和上一个一样
                if($j > 0 && $hand[$j - 1] == $hand[$j]){
                    continue;
                }

                //如果插入球与前球颜色相同（因为已经计算过了）
                if($i > 0 && $board[$i - 1] == $hand[$j]){
                    continue;
                }

                $mark = false;

                //如果插入球与当前球颜色相同可能产生更优解
                //打断连续球可能产生更优解（这个我是没想到的）
                //插入颜色与前后不同的球，完全没软用，只可能增加乱序度

                if($i < strlen($board) && $board[$i] == $hand[$j]){
                    $mark = true;
                }elseif($i > 0 && $i < strlen($board) && $board[$i - 1] == $board[$i] && $board[$i] != $hand[$j]){
                    $mark = true;
                }

                if($mark){
                    $newBoard = substr($board, 0, $i) . $hand[$j] . substr($board, $i);
                    $newHand = substr($hand, 0, $j) .  substr($hand, $j + 1);

                    //消除
                    while(true){
                        $tmpBoard = preg_replace('/(.)\1{2,}/', '', $newBoard);
                        if(strlen($tmpBoard) == strlen($newBoard)){
                            break;
                        }
                        $newBoard = $tmpBoard;
                    }
                    $res = min($res, $this->dfs($newBoard, $newHand) + 1);
                }

            }

        }

        $this->dp[$key] = $res;
        return $res;
    }
}

$obj = new Solution();
$res = $obj->findMinStep("RYYRRYYRYRYYRYYR","RRRRR");
var_dump($res);
// RYYRRY(R)YRYRYYRYYR
// RYY(R)RRYRYRYRYYRYYR RRYRYRYYRYYR
// RRYRYR(R)YYRYYR
// RRYRYRR(R)YYRYYR RRYRRYYR
// RRYRR(R)YYR OK 