## Overview
Basic POC to transpile `PL/I` code to `C#`

## How to run
`go run ./... test-pli-codes/test_3.pli > test.cs`

## Input Pl/I
```
MAIN: PROCEDURE;
	DECLARE i FIXED;
	DECLARE sum FIXED = 0;

	DO i = 1 TO 10;
		sum = sum + i;
	END;

END MAIN;
```
## Output C#
```cs
using System;

namespace PLIProgram
{
    public class Program
    {
        public static void Main(string[] args)
        {
            int i;
            int sum = 0;
            for (int i = 1; i <= 10; i += 1) {
                sum = sum + i;
            }
        }
    }
}
```
