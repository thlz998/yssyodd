<?php

define('__ROOT__', dirname(dirname(__FILE__)));
require_once(__ROOT__.'/app/LongestCommonPrefix.php');

use PHPUnit\Framework\TestCase;

class LongestCommonPrefixTest extends TestCase
{
    public function testLongestCommonPrefix()
    {
        $array = ["flower", "flow", "flight"];
        $result = "fl";

        $sumFinder = new LongestCommonPrefixClass();

        $this->assertSame($result, $sumFinder->longestCommonPrefix($array));
        $this->assertSame("f", $sumFinder->longestCommonPrefix(["fxlower", "flow", "flight"]));
        $this->assertSame("", $sumFinder->longestCommonPrefix(["afxlower", "flow", "flight"]));
        $this->assertSame("", $sumFinder->longestCommonPrefix([]));
        $this->assertSame("dsa", $sumFinder->longestCommonPrefix(["dsa"]));
    }

}
