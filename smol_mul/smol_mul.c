/*2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20? */
#include <stdbool.h>
#include <stdio.h>
#include "smol_mul.h"

#define print_bool(s) (s ? "true" : "false")

int calculate_multiples(){
  // let's try with double loop

  int number = 0;

  // discard numbers....

  /* 
    hay que crear algo que descomponga en numeros primos

    maxim comun multiplo(1,2,.....,20) = ??? 
    se multiplican los factores comunes de estos numeros y el resultado seria el maximo comun multiplo

  */
  bool is_prime = false;
  void is_factor_prime(int n)
  {
    if (n == 2)
    {
      is_prime = true;
      return;
    }
    for (int i = 2; i < n; i++)
    {
      printf("num i %d", i);
      if (n % i == 0)
      {
        return;
      }
      else
      {
        is_prime = true;
        break;
      }
    }
  }
  is_factor_prime(321251);
  printf("not prime -> %s\n", print_bool(is_prime));
  return 1;
}