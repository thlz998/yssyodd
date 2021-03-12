<?php

class RomanToInteger
{
    const ROMAN_ONE = 'I';
    const ROMAN_FIVE = 'V';
    const ROMAN_TEN = 'X';
    const ROMAN_FIFTY = 'L';
    const ROMAN_HUNDRED = 'C';
    const ROMAN_FIVE_HUNDRED = 'D';
    const ROMAN_THOUSAND = 'M';

    const ROMAN_MAP = [
        self::ROMAN_ONE => 1,
        self::ROMAN_FIVE => 5,
        self::ROMAN_TEN => 10,
        self::ROMAN_FIFTY => 50,
        self::ROMAN_HUNDRED => 100,
        self::ROMAN_FIVE_HUNDRED => 500,
        self::ROMAN_THOUSAND => 1000,
    ];

    /**
     * @param string $romans
     * @return int
     */
    public function romansToInteger(string $romans = ''): int
    {
        if (! $romans) {
            return 0;
        }

        $romanArr = str_split($romans);

        $sum = 0;
        $flag = 0;
        foreach ($romanArr as $key => $roman) {
            if ($flag) {
                $flag = 0;
                continue;
            }
            if (isset($romanArr[$key+1]) && $romanArr[$key] >= $romanArr[$key+1]) {
                $sum += self::ROMAN_MAP[$roman];
            } elseif (isset($romanArr[$key+1]) && $roman[$key] < $romanArr[$key+1]) {
                $sum = $sum + self::ROMAN_MAP[$romanArr[$key+1]] - self::ROMAN_MAP[$romanArr[$key]];
                $flag = 1;
            }
        }

        return $sum;
    }
}