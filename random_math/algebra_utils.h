#ifndef UTILS_F
#define UTILS_F


typedef struct{
    uint32_t cols;
    uint32_t rows;
    double **data;
} matrix;

matrix* init_matrix(uint32_t cols, uint32_t rows);
void printArray(matrix *m);
void fillArray(matrix *m, int32_t data);
matrix* applyArray(matrix *m);
matrix* copyMatrix(matrix *m);



#endif