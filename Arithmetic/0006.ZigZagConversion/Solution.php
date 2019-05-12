<?php

$s = 'LEETCODEISHIRING';
$numRows = 3;
$answer = solute($s, $numRows);
echo $answer;
function solute($s, $numRows)
{
    if($numRows == 1){
        return $s;
    }
    for($i = 0; $i < $numRows; $i ++){
        $data[$i] = '';
    }
    $currentRow = 0;
    $to = 0;
    for($i = 0; $i < strlen($s); $i ++){
        $data[$currentRow] .= $s[$i];
        if($to == 0){
            if($currentRow < $numRows - 1){
                $currentRow ++;
            }else{
                $currentRow --;
                $to = 1;
            }
        }else{
            if($currentRow > 0){
                $currentRow --;
            }else{
                $currentRow ++;
                $to = 0;
            }
        }
        
    }

    $answer = '';
    for($i = 0; $i < $numRows; $i ++){
        $answer .= $data[$i];
    }

    return $answer;
}

