
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include "algebra_utils.h"



matrix* createZerosArray(uint32_t x, uint32_t y){
    int i, k;
    matrix *mat = calloc(1, sizeof(matrix));
    mat->cols = x;
    mat->rows = y;
    mat->data = calloc(mat->rows, sizeof(*mat->data));
    for (i = 0; i < x; i++){
        mat->data[i] = calloc(1, sizeof(matrix));
    }

    return mat;
}

void printArray(matrix *mat){
    int col, row;

    for(col = 0; col < mat->cols; col++){
        printf("[");
        for(row = 0; row < mat->rows; row++){
            printf("%2.f", mat->data[col][row]);
        }
        printf("]\n");
    }
}

matrix* fillArray(matrix *mat, double data_vals){
    // we need to copy mat to a new m
    int col, row;
    for(col = 0; col < mat->cols; col++){
        for(row = 0; row < mat->rows; row++){
            mat->data[col][row] = data_vals;
        }
    }

    return mat;
}