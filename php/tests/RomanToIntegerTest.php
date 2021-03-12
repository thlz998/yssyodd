<?php

define('__ROOT__', dirname(dirname(__FILE__)));
require_once(__ROOT__.'/app/RomanToInteger.php');

use PHPUnit\Framework\TestCase;

class RomanToIntegerTest extends TestCase
{
    public function testRomanToInteger()
    {
        $array = [
            '3'     => 'III',
            '4'     => 'IV',
            '9'     => 'IX',
            '58'    => 'LVIII',
            '0'     => '',
            '1437'  => 'MCDXXXVII',
            '1994'  => 'MCMXCIV',
        ];

        $romanToInteger = new RomanToInteger();

        foreach ($array as $key => $value) {
            $this->assertSame($key, $romanToInteger->romansToInteger($value));
        }
    }
}