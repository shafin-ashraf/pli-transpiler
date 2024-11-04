## Overview
Basic POC to transpile `PL/I` code to `Javascript`

## How to run
`go run ./... test-pli-codes/test_3.pli > test.js`

## Input Pl/I
```
MAIN: PROCEDURE;
    DECLARE x FIXED = 5;
    DECLARE y FIXED = 10;
    DECLARE result FIXED;

    result = x + y;
END;

HELPER: PROCEDURE;
    DECLARE temp FIXED = 100;
    DECLARE factor FIXED = 2;
    temp = temp * factor;
END;
```
## Output javascript
```javascript
(function() {
function MAIN() {
  let x = 5;
  let y = 10;
  let result;
  result = x + y;
}
function HELPER() {
  let temp = 100;
  let factor = 2;
  temp = temp * factor;
}
})();
```
