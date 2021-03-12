<?php

define('__ROOT__', dirname(dirname(__FILE__)));
require_once(__ROOT__.'/app/ValidParentheses.php');

use PHPUnit\Framework\TestCase;

class ValidParenthesesTest extends TestCase
{
    public function testParentheses()
    {
        $array = [
            '()'        => true,
//            '()[]{}'    => true,
//            '(]'        => false,
//            '([)]'      => false,
//            '{[]}'      => true,
//            '{]}'       => false,
//            '{[]'       => false,
        ];

        $validParentheses = new ValidParentheses();

        foreach ($array as $key => $value) {
            $this->assertSame($value, $validParentheses->valid($key));
        }
    }
}