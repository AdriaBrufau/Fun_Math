#ifndef UTILS_H
#define UTILS_H


void applyFunc(void(*f) (float));
float function(float x);
double** Create2DArray(int start, int end, float z);
void Destroy2DArray(float** arr);




#endif