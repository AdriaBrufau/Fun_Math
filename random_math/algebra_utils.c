
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include "algebra_utils.h"



matrix* init_matrix(uint32_t x, uint32_t y){
    int i, k;
    matrix *mat = calloc(1, sizeof(matrix));
    mat->cols = x;
    mat->rows = y;
    mat->data = malloc(x * sizeof(double*));
    for (i = 0; i < x; i++){
        mat->data[i] = malloc(y * sizeof(double));
    }

    return mat;
}

matrix* matrix_copy(matrix *mat){
    matrix *new_matrix = init_matrix(mat->rows, mat->cols);
    for(int i = 0; i < mat->rows; i++){
        for(int k = 0; k<mat->cols; k++){
            new_matrix->data[i][k] = mat->data[i][k];
        }
    }

    return new_matrix;
}

void printArray(matrix *mat){
    int col, row;

    for(col = 0; col < mat->rows; col++){
        for(row = 0; row < mat->rows; row++){
            printf("%1.2f ", mat->data[col][row]);
        }
        printf("\n");
    }
}

void* fillArray(matrix *mat, int32_t data_vals){
    for(int col = 0; col < mat->cols; col++){
        for(int row = 0; row < mat->rows; row++){
            mat->data[col][row] = data_vals;
        }
    }
}
