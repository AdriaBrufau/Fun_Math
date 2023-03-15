#include <stdio.h>
#include <stdint.h>
#include "random.h"
#include "algebra_utils.h"

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
    double vals[3] = {2.2,3.3,1.1};
    float derivative;
    matrix *mat = init_matrix(4, 4);
    mat = fillArray(mat, 4);
    printArray(mat);

    printf("\n");
   

    return 0;
}

