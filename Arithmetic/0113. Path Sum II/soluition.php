<?php

class TreeNode {
   public $val = null;
   public $left = null;
   public $right = null;
   function __construct($val = 0, $left = null, $right = null) {
        $this->val = $val;
        $this->left = $left;
        $this->right = $right;
  }
}


function pathSum($root, $targetSum) {
    $res = [];

    dfs($root, $targetSum, $res, []);

    return $res;
}

function dfs($node, $sum, &$res, $path){
    if($node === null){
        return;
    }

    $sum -= $node->val;
    $path[] = $node->val;
    if($root->left === null && $root->right === null && $sum == 0){
        $res[] = $path;
    }
    dfs($node->left, $sum, $res, $path);
    dfs($node->right, $sum, $res, $path);

}
