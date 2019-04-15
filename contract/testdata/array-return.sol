pragma solidity 0.4.24;

contract B {
    function B() {}

    function bar(uint x) constant returns(uint) {
        return x;
    }
    function barstring(string y)returns(string){
        return y;
    }
}
