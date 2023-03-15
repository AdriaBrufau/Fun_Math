#ifndef UTILS_F
#define UTILS_F


typedef struct{
    uint32_t cols;
    uint32_t rows;
    double **data;
} matrix;

matrix* createZerosArray(uint32_t cols, uint32_t rows);
void printArray(matrix *m);
matrix* fillArray(matrix *m, double data);
matrix* applyArray(matrix *m);



#endif