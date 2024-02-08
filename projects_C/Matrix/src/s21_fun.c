#include "s21_matrix.h"

int matrix_size_check(matrix_t *A, matrix_t *B) {
  int flag = SUCCESS;
  if (A->rows != B->rows || A->columns != B->columns) {
    flag = FAILURE;
  }
  return flag;
}

void less(matrix_t *A, int row, int column, matrix_t *result) {
  for (int i = 0; i < A->rows; i++) {
    for (int j = 0; j < A->columns; j++) {
      if (i < row && j < column) {
        result->matrix[i][j] = A->matrix[i][j];
      } else if (i > row && j > column) {
        result->matrix[i - 1][j - 1] = A->matrix[i][j];
      } else if (i < row && j > column) {
        result->matrix[i][j - 1] = A->matrix[i][j];
      } else if (i > row && j < column) {
        result->matrix[i - 1][j] = A->matrix[i][j];
      }
    }
  }
}

void matrix_null(matrix_t *A) {
  for (int i = 0; i < A->rows; i++) {
    for (int j = 0; j < A->columns; j++) {
      A->matrix[i][j] = 0;
    }
  }
}

void null_A(matrix_t *A) {
  A->rows = 0;
  A->columns = 0;
  A->matrix = NULL;
}
