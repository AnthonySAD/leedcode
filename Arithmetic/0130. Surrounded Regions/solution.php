<?php

//想简单了，该方法不行。只能把相邻的格子算正确。还是得用深度搜索算法。
function solve(&$board) {
    $height = count($board);
    $width  = count($board[0]);
    $mark = 1;
    $badPoints = [];
    for($h = 0; $h < $height; $h ++){
        for($w = 0; $w < $width; $w ++){
            if($board[$h][$w] != "X"){
                if($board[$h][$w] == 'O'){
                    if($h > 0 && $board[$h - 1][$w] != 'X'){
                        $board[$h][$w] = $board[$h - 1][$w];
                    }elseif($w > 0 && $board[$h][$w - 1] !='X'){
                        $board[$h][$w] = $board[$h][$w - 1];
                    }elseif($h + 1 < $height && $board[$h + 1][$w] != 'O' && $board[$h + 1][$w] != 'X'){
                        $board[$h][$w] = $board[$h + 1][$w];
                    }elseif($w + 1 < $width && $board[$h][$w + 1] != 'O' && $board[$h][$w + 1] != 'X'){
                        $board[$h][$w] = $board[$h][$w + 1];
                    }{
                        $board[$h][$w] = $mark;
                        $mark ++;
                    }

                    if($h + 1 < $height && $board[$h + 1][$w] == 'O'){
                        $board[$h + 1][$w] = $board[$h][$w];
                    }
                    if($w + 1 < $width && $board[$h][$w + 1] == 'O'){
                        $board[$h][$w + 1] = $board[$h][$w];
                    }
                }
                
                if($h == 0 || $w == 0 || $h + 1 == $height || $w + 1 == $width){
                    $badPoints[$board[$h][$w]] = 1;
                }
            }
        }
    }


    for($h = 0; $h < $height; $h ++){
        for($w = 0; $w < $width; $w ++){
            if($board[$h][$w] !== 'X'){

                if(isset($badPoints[$board[$h][$w]])){
                    $board[$h][$w] = 'O';
                }else{
                    $board[$h][$w] = 'X';
                }
            }
        }
    }

    return $board;
}

//就从四周开始遍历即可
class Solution {

    public $board = [];
    public $height = 0;
    public $width = 0;

    function solve(&$board) {
        $this->board = $board;
        $this->height = count($board);
        $this->width  = count($board[0]);
        
        for($w = 0; $w < $this->width; $w ++){
            if($this->board[0][$w] == 'O'){
                $this->board[0][$w] = 1;
                $this->do(0, $w);
            }
            if($this->board[$this->height - 1][$w] == 'O'){
                $this->board[$this->height - 1][$w] = 1;
                $this->do($this->height - 1, $w);
            }
            
        }
    
        for($h = 0; $h < $this->height; $h ++){
            if($this->board[$h][0] == 'O'){
                $this->board[$h][0] = 1;
                $this->do($h, 0);
            }

            if($this->board[$h][$w - 1] == 'O'){
                $this->board[$h][$w - 1] = 1;
                $this->do($h, $w - 1);
            }
        }
        // var_dump($this->board);
        
        for($h = 0; $h < $this->height; $h ++){
            for($w = 0; $w < $this->width; $w ++){
                if($this->board[$h][$w] == 1){
                    $this->board[$h][$w] = 'O';
                }elseif($this->board[$h][$w] == 'O'){
                    $this->board[$h][$w] = 'X';
                }
            }
        }
        $board = $this->board;
    }

    function do($h, $w){
        if($h > 0 && $this->board[$h - 1][$w] == 'O'){
            $this->board[$h - 1][$w] = 1;
            $this->do($h - 1, $w);
        }
        
        if($w > 0 && $this->board[$h][$w - 1] =='O'){
            $this->board[$h][$w - 1] = 1;
            $this->do($h, $w - 1);
        }
        
        if($h + 1 < $this->height && $this->board[$h + 1][$w] == 'O'){
            $this->board[$h + 1][$w] = 1;
            $this->do($h + 1, $w);
        }
        
        if($w + 1 < $this->width && $this->board[$h][$w + 1] == 'O'){
            $this->board[$h][$w + 1] = 1;
            $this->do($h, $w + 1);
        }
    }
}



$board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]];

$obj = new Solution;
var_dump($obj->solve($board));