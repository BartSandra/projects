#include "s21_matrix.h"

int s21_create_matrix(int rows, int columns, matrix_t *result) {
  null_A(result);

  if (rows <= 0 || columns <= 0) {
    null_A(result);
    return ERR_MATRIX;
  }
  result->rows = rows;
  result->columns = columns;
  result->matrix = (double **)malloc(rows * sizeof(double *));

  if (result->matrix == NULL) {
    return ERR_CALC;
  }
  for (int i = 0; i < rows; ++i) {
    result->matrix[i] = (double *)malloc(columns * sizeof(double));
    if (result->matrix[i] == NULL) {
      for (int j = 0; j < i; ++j) {
        free(result->matrix[j]);
      }
      free(result->matrix);
      return ERR_CALC;
    }
  }
  matrix_null(result);

  return 0;
}

void s21_remove_matrix(matrix_t *A) {
  if (A->matrix != NULL) {
    // matrix_null(A);
    for (int i = 0; i < A->rows; i++) {
      free(A->matrix[i]);
      A->matrix[i] = NULL;
    }
  }
  free(A->matrix);
  null_A(A);
}

int s21_eq_matrix(matrix_t *A, matrix_t *B) {
  int flag = SUCCESS;
  if (A->rows != B->rows || A->columns != B->columns) {
    flag = FAILURE;
  } else {
    for (int i = 0; i < A->rows; i++) {
      for (int j = 0; j < A->columns; j++) {
        if (fabs(A->matrix[i][j] - B->matrix[i][j]) > EPS) {
          flag = FAILURE;
          //     break;
        }
      }
    }
  }
  return flag;
}

int s21_sum_matrix(matrix_t *A, matrix_t *B, matrix_t *result) {
  int flag = OK;
  null_A(result);

  if (A->rows < 1 || A->columns < 1 || B->rows < 1 || B->columns < 1 ||
      A == NULL || B == NULL) {
    flag = ERR_MATRIX;
  } else {
    if (matrix_size_check(A, B)) {
      if (result == NULL) {
        result = (matrix_t *)malloc(sizeof(matrix_t));
      }
      s21_create_matrix(A->rows, A->columns, result);
      for (int i = 0; i < A->rows; i++) {
        for (int j = 0; j < A->columns; j++) {
          result->matrix[i][j] = A->matrix[i][j] + B->matrix[i][j];
        }
      }
    } else {
      flag = ERR_CALC;
    }
  }
  return flag;
}

int s21_sub_matrix(matrix_t *A, matrix_t *B, matrix_t *result) {
  int flag = OK;
  null_A(result);

  if (A->rows < 1 || A->columns < 1 || B->rows < 1 || B->columns < 1 ||
      A == NULL || B == NULL) {
    flag = ERR_MATRIX;
  } else {
    if (matrix_size_check(A, B)) {
      if (result == NULL) {
        result = (matrix_t *)malloc(sizeof(matrix_t));
      }
      s21_create_matrix(A->rows, A->columns, result);

      for (int i = 0; i < A->rows; i++) {
        for (int j = 0; j < A->columns; j++) {
          result->matrix[i][j] = A->matrix[i][j] - B->matrix[i][j];
        }
      }
    } else {
      flag = ERR_CALC;
    }
  }
  return flag;
}

int s21_mult_number(matrix_t *A, double number, matrix_t *result) {
  int flag = OK;
  null_A(result);

  if (A->rows == 0 || A->columns == 0 || A == NULL) {
    flag = ERR_MATRIX;
  } else {
    if (result == NULL) {
      result = (matrix_t *)malloc(sizeof(matrix_t));
    }
    s21_create_matrix(A->rows, A->columns, result);

    for (int i = 0; i < A->rows; i++) {
      for (int j = 0; j < A->columns; j++) {
        result->matrix[i][j] = A->matrix[i][j] * number;
      }
    }
  }
  return flag;
}

int s21_mult_matrix(matrix_t *A, matrix_t *B, matrix_t *result) {
  int flag = OK;
  null_A(result);

  if (A->rows == 0 || A->columns == 0 || A == NULL || B == NULL) {
    flag = ERR_MATRIX;
  } else {
    if (A->columns == B->rows) {
      if (result == NULL) {
        result = (matrix_t *)malloc(sizeof(matrix_t));
      }
      s21_create_matrix(A->rows, B->columns, result);

      for (int i = 0; i < A->rows; i++) {
        for (int j = 0; j < B->columns; j++) {
          for (int k = 0; k < A->columns; k++) {
            result->matrix[i][j] += A->matrix[i][k] * B->matrix[k][j];
          }
        }
      }
    } else {
      flag = ERR_CALC;
    }
  }
  return flag;
}

int s21_transpose(matrix_t *A, matrix_t *result) {
  int flag = OK;
  null_A(result);

  if (A->rows == 0 || A->columns == 0 || A == NULL) {
    flag = ERR_MATRIX;
  } else {
    if (result == NULL) {
      result = (matrix_t *)malloc(sizeof(matrix_t));
    }
    s21_create_matrix(A->columns, A->rows, result);

    for (int i = 0; i < A->columns; i++) {
      for (int j = 0; j < A->rows; j++) {
        result->matrix[i][j] = A->matrix[j][i];
      }
    }
  }
  return flag;
}

int s21_calc_complements(matrix_t *A, matrix_t *result) {
  int flag = OK;

  if (A == NULL || A->rows <= 0 || A->columns <= 0) {
    flag = ERR_MATRIX;
  } else {
    null_A(result);

    if (A->rows != A->columns || A->columns < 1) {
      flag = ERR_CALC;
    } else {
      if (result == NULL) {
        result = (matrix_t *)malloc(sizeof(matrix_t));
        matrix_null(result);
      }
      s21_create_matrix(A->columns, A->rows, result);
      if (A->columns == 1) {
        result->matrix[0][0] = A->matrix[0][0];
      } else {
        matrix_t matrix_low = {0};
        double determ_low;
        int al = 1;
        for (int i = 0; i < A->rows; i++) {
          for (int j = 0; j < A->columns; j++) {
            s21_create_matrix(A->rows - 1, A->columns - 1, &matrix_low);
            less(A, i, j, &matrix_low);
            s21_determinant(&matrix_low, &determ_low);
            al = (i + j) % 2 == 0 ? 1 : -1;
            result->matrix[i][j] = determ_low * al;
            s21_remove_matrix(&matrix_low);
          }
        }
      }
    }
  }
  return flag;
}

int s21_determinant(matrix_t *A, double *result) {
  int flag = OK;
  *result = 0;
  if (result != NULL) {
    if (A == NULL || A->rows == 0 || A->columns == 0) {
      flag = ERR_MATRIX;
    } else {
      if (A->columns != A->rows) {
        flag = ERR_CALC;
      } else {
        if (A->columns == 1) {
          *result = A->matrix[0][0];
        } else if (A->columns == 2) {
          *result = A->matrix[0][0] * A->matrix[1][1] -
                    A->matrix[1][0] * A->matrix[0][1];
        } else {
          double determ_low = 0;
          int al = 1;
          for (int i = 0; i < A->columns; i++) {
            al = i % 2 == 0 ? 1 : -1;
            matrix_t matrix_low = {0};
            s21_create_matrix(A->rows - 1, A->columns - 1, &matrix_low);
            less(A, 0, i, &matrix_low);
            s21_determinant(&matrix_low, &determ_low);
            *result += al * determ_low * A->matrix[0][i];
            s21_remove_matrix(&matrix_low);
          }
        }
      }
    }
  }
  return flag;
}

int s21_inverse_matrix(matrix_t *A, matrix_t *result) {
  int flag = OK;

  null_A(result);
  if (A->rows == 0 || A->columns == 0 || A == NULL) {
    flag = ERR_MATRIX;
  } else {
    if (A->rows != A->columns) {
      flag = ERR_CALC;
    } else {
      if (result == NULL) {
        result = (matrix_t *)malloc(sizeof(matrix_t));
      }
      //   s21_create_matrix(A->columns, A->rows, result);
      if (A->rows == 1) {
        s21_create_matrix(A->columns, A->rows, result);
        if (fabs(A->matrix[0][0]) > EPS) {
          result->matrix[0][0] = (double)(1.0 / A->matrix[0][0]);
        } else {
          flag = ERR_CALC;
        }
      } else {
        double determ = 0;
        s21_determinant(A, &determ);
        if (determ == 0) {
          flag = ERR_CALC;
        } else {
          matrix_t alg_dop = {0};
          matrix_t dop_transp = {0};
          s21_calc_complements(A, &alg_dop);
          s21_transpose(&alg_dop, &dop_transp);
          s21_mult_number(&dop_transp, (double)(1.0 / determ), result);
          s21_remove_matrix(&alg_dop);
          s21_remove_matrix(&dop_transp);
        }
      }
    }
  }
  return flag;
}
