<?php

class ValidParentheses
{
    const LeftParenthesis = '(';
    const LeftBrackets = '[';
    const LeftBrace = '{';

    const LeftMap = [
        self::LeftParenthesis => ')',
        self::LeftBrackets => ']',
        self::LeftBrace => '}',
    ];

    /**
     * @param string $romans
     * @return int
     */
    public function valid(string $parentheses = '')
    {
        if (! $parentheses) {
            return true;
        }
        if (strlen($parentheses) & 1) {
            return false;
        }

        $parenthesesArr = str_split($parentheses);
        $arr = [];

        $ret = true;
        foreach ($parenthesesArr as $val) {
            if (in_array($val, self::LeftMap)) {
                array_push($arr, $val);
            } else {
                $temp = array_pop($arr);
                if ($val == self::LeftMap[$temp]) {
                    continue;
                } else {
                    $ret = false;
                    break;
                }
            }
        }

        return $ret;
    }
}