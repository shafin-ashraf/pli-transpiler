## Overview
Basic POC to transpile `PL/I` code to `C#`

## How to run
`go run ./... test-pli-codes/test_3.pli > test.cs`

## Input Pl/I
```
MATHUTILS: PROCEDURE;
   DECLARE NUM FIXED;
   DECLARE RESULT FIXED;

   NUM = 15;

   IF NUM > 10 THEN
      CALL CALCULATE_SQUARE;
   ELSE
      CALL CALCULATE_CUBE;
   END;

   CALL PRINT_SERIES;
END;

CALCULATE_SQUARE: PROCEDURE;
   DECLARE TEMP FIXED;
   TEMP = NUM * NUM;
   RESULT = TEMP;
END;

CALCULATE_CUBE: PROCEDURE;
   DECLARE TEMP FIXED;
   TEMP = NUM * NUM * NUM;
   RESULT = TEMP;
END;

PRINT_SERIES: PROCEDURE;
   DECLARE I FIXED;
   DECLARE SUM FIXED;

   SUM = 0;

   DO I = 1 TO RESULT BY 2;
      IF I < 10 THEN
         SUM = SUM + I;
      ELSE
         IF I < 20 THEN
            SUM = SUM + I * 2;
         ELSE
            SUM = SUM + I * 3;
         END;
      END;
   END;

   RESULT = SUM;
END;
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
            // Entry point
            var program = new Program();
            program.MATHUTILS();
        }

        public void MATHUTILS()
        {
            int NUM;
            int RESULT;
            NUM = 15;
            if (NUM > 10)
            {
                CALL = ;
                CALCULATE_SQUARE = ;
            }
            else
            {
                CALL = ;
                CALCULATE_CUBE = ;
            }
            CALL = ;
            PRINT_SERIES = ;
        }
        public void CALCULATE_SQUARE()
        {
            int TEMP;
            TEMP = NUM * NUM;
            RESULT = TEMP;
        }
        public void CALCULATE_CUBE()
        {
            int TEMP;
            TEMP = NUM * NUM * NUM;
            RESULT = TEMP;
        }
        public void PRINT_SERIES()
        {
            int I;
            int SUM;
            SUM = 0;
            for (int I = 1; I <= RESULT; I += 2) {
                if (I < 10)
                {
                    SUM = SUM + I;
                }
                else
                {
                    if (I < 20)
                    {
                        SUM = SUM + I * 2;
                    }
                    else
                    {
                        SUM = SUM + I * 3;
                    }
                }
            }
            RESULT = SUM;
        }
    }
}
```
