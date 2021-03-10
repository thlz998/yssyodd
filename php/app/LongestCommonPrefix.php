<?php

class LongestCommonPrefixClass
{
    private $inputArray;

    public function __construct(array $inputArray = null)
    {
        $this->inputArray = $inputArray;
    }

    public function longestCommonPrefix(array $strs)
    {
        if (! $strs) {
            return '';
        } elseif (count($strs) == 1) {
            return $strs[0];
        }
        $arr = [];
        foreach ($strs as $str) {
            $arr[] = str_split($str);
        }
        $count = count($strs);
        $length = count($arr[0]);
        $s = [];
        $similar = true;
        $temp = '';
        for ($i = 0; $i < $length; $i++) {
            for ($j = 0; $j < $count - 1; $j++) {
                $temp = $arr[$j][$i];
                
                if ($arr[$j][$i] != $arr[$j+1][$i]) {
                    $similar = false;
                    break;
                }
            }
            if (! $similar) {
                break;
            }
            $s[] = $temp;
            $temp = '';
            $similar = true;
        }
        return implode('', $s);
    }
}