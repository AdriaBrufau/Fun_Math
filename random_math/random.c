#include <stdio.h>
#include <stdint.h>
#include "random.h"
#include "utils.h"

#define X_CONSOLE_SIZE  100
#define Y_CONSOLE_SIZE  100
#define derivation 0.0001
#define MAX_SIZE 10
#define X_SIZE MAX_SIZE-5

int main(int argc, char *argv[]){
    
    int size_x, size_y;
    size_x = 5;
    size_y = 10;
    float test = 3.0;
    float derivative;
    test = function(test);
    derivative = derivative_floats(test, derivation);
    //CreateCanvas(X_CONSOLE_SIZE, Y_CONSOLE_SIZE);
    double** arr = Create2DArray(X_SIZE, MAX_SIZE, 0.5);
    for(int i = 0; i < X_SIZE; i++){
        printf("[");
        for(int y = 0; y < MAX_SIZE;y++){
            printf(", %.6f", arr[i][y]);
        }
        printf("]\n");
    }

    printf("\n");
   

    return 0;
}

