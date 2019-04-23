<?php
/**
 * Created by PhpStorm.
 * User: Anthony
 * Date: 2019/4/23
 * Time: 9:18
 */

class Solution
{
    public $answerLen = 1;
    public $answerStart = 0;
    public $len;
    public $str;

    public function answer($str)
    {
        $this->str = $str;
        $this->len = strlen($str);
        if ($this->len < 1){
            return 0;
        }

        for ($i = 0; $i < $this->len; $i ++)
        {
            $this->handleAnswer($i - 1, $i + 1);
            $this->handleAnswer($i, $i + 1);
        }

        return substr($this->str, $this->answerStart, $this->answerLen);
    }

    private function handleAnswer($start, $end)
    {
        while ($start > -1 && $end < $this->len && $this->str[$start] == $this->str[$end])
        {
            $tempLen = $end - $start + 1;
            if ($tempLen > $this->answerLen){
                $this->answerLen = $tempLen;
                $this->answerStart = $start;
            }
            $start --;
            $end ++;
        }
    }
}

$solution = new Solution();
echo $solution->answer('ababacdfdfsafaafafdsfafaf');