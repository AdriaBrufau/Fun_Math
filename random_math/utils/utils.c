
#include "utils.h"
#include <stdio.h>
#include <stdlib.h>


float function(float x){
    return 4 * x * x + 3 * x + 4;
}

float derivative_floats(float d, float h){
    return ((d+h - d)/h);
}
        

double** Create2DArray(int dimA, int dimB, float disipation){ 
    float entropy_rand = rand();

    int i;
    double *x = calloc(dimA*dimB, sizeof(double));
    double **y = malloc(dimB * sizeof(double*));


    for(i = 0; i < dimA; i++){
        x[i] = disipation+=entropy_rand;
    }
    return x;

}

void Destroy2DArray(float** arr){
    free(*arr);
    free(arr);
}