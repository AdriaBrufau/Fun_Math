
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
        printf_s("[");
        for(row = 0; row < mat->rows; row){
            printf_s(mat->data[col][row]);
        }
        printf_s("]");
    }
}