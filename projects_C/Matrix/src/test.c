#include <check.h>
#include <stdlib.h>

#include "s21_matrix.h"

START_TEST(s21_create_matrix_test1) {
  matrix_t result;
  int matrix = s21_create_matrix(2, 2, &result);
  s21_remove_matrix(&result);
  ck_assert_int_eq(0, matrix);
}
END_TEST

START_TEST(s21_create_matrix_test2) {
  matrix_t result;
  int matrix = s21_create_matrix(1, 2, &result);
  s21_remove_matrix(&result);
  ck_assert_int_eq(0, matrix);
}
END_TEST

START_TEST(s21_create_matrix_test3) {
  matrix_t result;
  int matrix = s21_create_matrix(-2, -2, &result);
  ck_assert_int_eq(1, matrix);
}
END_TEST

START_TEST(s21_create_matrix_test4) {
  matrix_t result;
  int matrix = s21_create_matrix(-2, -2, &result);
  ck_assert_int_ne(0, matrix);
}
END_TEST

START_TEST(s21_create_matrix_test5) {
  matrix_t result;
  int matrix = s21_create_matrix(2, -2, &result);
  ck_assert_int_eq(1, matrix);
}
END_TEST

Suite* s21_create_matrix_tests(void) {
  Suite* s;
  TCase* tc;
  s = suite_create("s21_create_matrix");
  tc = tcase_create("case_create");
  tcase_add_test(tc, s21_create_matrix_test1);
  tcase_add_test(tc, s21_create_matrix_test2);
  tcase_add_test(tc, s21_create_matrix_test3);
  tcase_add_test(tc, s21_create_matrix_test4);
  tcase_add_test(tc, s21_create_matrix_test5);

  suite_add_tcase(s, tc);
  return s;
}

START_TEST(s21_eq_matrix_test1) {
  matrix_t matrix1;
  matrix_t matrix2;

  s21_create_matrix(2, 2, &matrix1);
  matrix1.matrix[0][0] = 2;
  matrix1.matrix[0][1] = 4;
  matrix1.matrix[1][0] = 0;
  matrix1.matrix[1][1] = 9;

  s21_create_matrix(2, 2, &matrix2);
  matrix2.matrix[0][0] = 2;
  matrix2.matrix[0][1] = 4;
  matrix2.matrix[1][0] = 0;
  matrix2.matrix[1][1] = 9;

  int result = s21_eq_matrix(&matrix1, &matrix2);
  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);

  ck_assert_int_eq(SUCCESS, result);
}
END_TEST

START_TEST(s21_eq_matrix_test2) {
  matrix_t matrix1;
  matrix_t matrix2;

  s21_create_matrix(2, 2, &matrix1);
  matrix1.matrix[0][0] = 2.99999999;
  matrix1.matrix[0][1] = 4.99999999;
  matrix1.matrix[1][0] = 0.99999999;
  matrix1.matrix[1][1] = 9.99999999;

  s21_create_matrix(2, 2, &matrix2);
  matrix2.matrix[0][0] = 2.99999999;
  matrix2.matrix[0][1] = 4.99999999;
  matrix2.matrix[1][0] = 0.99999999;
  matrix2.matrix[1][1] = 9.99999999;

  int result = s21_eq_matrix(&matrix1, &matrix2);
  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);

  ck_assert_int_eq(SUCCESS, result);
}
END_TEST

START_TEST(s21_eq_matrix_test3) {
  matrix_t matrix1;
  matrix_t matrix2;

  s21_create_matrix(2, 2, &matrix1);
  matrix1.matrix[0][0] = 2.99999999;
  matrix1.matrix[0][1] = 4.99999999;
  matrix1.matrix[1][0] = 0.99999999;
  matrix1.matrix[1][1] = 9.99999999;

  s21_create_matrix(2, 2, &matrix2);
  matrix2.matrix[0][0] = 2.99909999;
  matrix2.matrix[0][1] = 4.99909999;
  matrix2.matrix[1][0] = 0.99909999;
  matrix2.matrix[1][1] = 9.99909999;

  int result = s21_eq_matrix(&matrix1, &matrix2);
  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);

  ck_assert_int_eq(FAILURE, result);
}
END_TEST

START_TEST(s21_eq_matrix_test4) {
  matrix_t matrix1;
  matrix_t matrix2;

  s21_create_matrix(2, 2, &matrix1);
  matrix1.matrix[0][0] = 2.99999999;
  matrix1.matrix[1][0] = 0.99999999;

  s21_create_matrix(2, 2, &matrix2);
  matrix2.matrix[0][0] = 2.99909999;
  matrix2.matrix[0][1] = 4.99909999;
  matrix2.matrix[1][0] = 0.99909999;
  matrix2.matrix[1][1] = 9.99909999;

  int result = s21_eq_matrix(&matrix1, &matrix2);
  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);
  ck_assert_int_eq(FAILURE, result);
}
END_TEST

Suite* s21_eq_matrix_tests(void) {
  Suite* s;
  TCase* tc;
  s = suite_create("s21_eq_matrix");
  tc = tcase_create("case_eq");
  tcase_add_test(tc, s21_eq_matrix_test1);
  tcase_add_test(tc, s21_eq_matrix_test2);
  tcase_add_test(tc, s21_eq_matrix_test3);
  tcase_add_test(tc, s21_eq_matrix_test4);

  suite_add_tcase(s, tc);
  return s;
}

START_TEST(s21_sum_matrix_test1) {
  matrix_t matrix1;
  matrix_t matrix2;
  matrix_t matrix1_2;

  s21_create_matrix(2, 2, &matrix1);
  matrix1.matrix[0][0] = 2;
  matrix1.matrix[0][1] = 4;
  matrix1.matrix[1][0] = 0;
  matrix1.matrix[1][1] = 9;

  s21_create_matrix(2, 2, &matrix2);
  matrix2.matrix[0][0] = 2;
  matrix2.matrix[0][1] = 4;
  matrix2.matrix[1][0] = 0;
  matrix2.matrix[1][1] = 9;

  int result = s21_sum_matrix(&matrix1, &matrix2, &matrix1_2);

  ck_assert_int_eq(4, matrix1_2.matrix[0][0]);
  ck_assert_int_eq(8, matrix1_2.matrix[0][1]);
  ck_assert_int_eq(0, matrix1_2.matrix[1][0]);
  ck_assert_int_eq(18, matrix1_2.matrix[1][1]);
  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);
  s21_remove_matrix(&matrix1_2);
}
END_TEST

START_TEST(s21_sum_matrix_test2) {
  matrix_t matrix1;
  matrix_t matrix2;
  matrix_t matrix1_2;

  s21_create_matrix(4, 2, &matrix1);
  s21_create_matrix(2, 4, &matrix2);

  int result = s21_sum_matrix(&matrix1, &matrix2, &matrix1_2);

  ck_assert_int_eq(result, ERR_CALC);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);
}
END_TEST

Suite* s21_sum_matrix_tests(void) {
  Suite* s;
  TCase* tc;
  s = suite_create("s21_sum_matrix");
  tc = tcase_create("case_sum");
  tcase_add_test(tc, s21_sum_matrix_test1);
  tcase_add_test(tc, s21_sum_matrix_test2);

  suite_add_tcase(s, tc);
  return s;
}

START_TEST(s21_sub_matrix_test1) {
  matrix_t matrix1;
  matrix_t matrix2;
  matrix_t matrix1_2;

  s21_create_matrix(2, 2, &matrix1);
  matrix1.matrix[0][0] = 7;
  matrix1.matrix[0][1] = 5;
  matrix1.matrix[1][0] = 0;
  matrix1.matrix[1][1] = 9;

  s21_create_matrix(2, 2, &matrix2);
  matrix2.matrix[0][0] = 2;
  matrix2.matrix[0][1] = 4;
  matrix2.matrix[1][0] = 0;
  matrix2.matrix[1][1] = 20;

  int result = s21_sub_matrix(&matrix1, &matrix2, &matrix1_2);

  ck_assert_int_eq(5, matrix1_2.matrix[0][0]);
  ck_assert_int_eq(1, matrix1_2.matrix[0][1]);
  ck_assert_int_eq(0, matrix1_2.matrix[1][0]);
  ck_assert_int_eq(-11, matrix1_2.matrix[1][1]);
  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);
  s21_remove_matrix(&matrix1_2);
}
END_TEST

START_TEST(s21_sub_matrix_test2) {
  matrix_t matrix1;
  matrix_t matrix2;
  matrix_t matrix1_2;

  s21_create_matrix(4, 4, &matrix1);
  s21_create_matrix(2, 4, &matrix2);

  int result = s21_sub_matrix(&matrix1, &matrix2, &matrix1_2);

  ck_assert_int_eq(result, ERR_CALC);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);
}
END_TEST

Suite* s21_sub_matrix_tests(void) {
  Suite* s;
  TCase* tc;
  s = suite_create("s21_sub_matrix");
  tc = tcase_create("case_sub");
  tcase_add_test(tc, s21_sub_matrix_test1);
  tcase_add_test(tc, s21_sub_matrix_test2);

  suite_add_tcase(s, tc);
  return s;
}

START_TEST(s21_mult_number_test1) {
  matrix_t matrix1;
  matrix_t matrix_res;

  s21_create_matrix(2, 2, &matrix1);
  matrix1.matrix[0][0] = 7;
  matrix1.matrix[0][1] = 5;
  matrix1.matrix[1][0] = 0;
  matrix1.matrix[1][1] = 9;

  int result = s21_mult_number(&matrix1, 11, &matrix_res);

  ck_assert_int_eq(77, matrix_res.matrix[0][0]);
  ck_assert_int_eq(55, matrix_res.matrix[0][1]);
  ck_assert_int_eq(0, matrix_res.matrix[1][0]);
  ck_assert_int_eq(99, matrix_res.matrix[1][1]);
  ck_assert_int_eq(matrix_res.rows, 2);
  ck_assert_int_eq(matrix_res.columns, 2);
  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix_res);
}
END_TEST

Suite* s21_mult_number_tests(void) {
  Suite* s;
  TCase* tc;
  s = suite_create("s21_mult_number");
  tc = tcase_create("case_mult_number");
  tcase_add_test(tc, s21_mult_number_test1);

  suite_add_tcase(s, tc);
  return s;
}

START_TEST(s21_mult_matrix_test2) {
  matrix_t matrix1;
  matrix_t matrix2;
  matrix_t matrix1_2;

  s21_create_matrix(2, 2, &matrix1);
  matrix1.matrix[0][0] = 3;
  matrix1.matrix[0][1] = 5;
  matrix1.matrix[1][0] = 2;
  matrix1.matrix[1][1] = 1;

  s21_create_matrix(2, 3, &matrix2);
  matrix2.matrix[0][0] = 8;
  matrix2.matrix[0][1] = 2;
  matrix2.matrix[0][2] = 3;
  matrix2.matrix[1][0] = 1;
  matrix2.matrix[1][1] = 7;
  matrix2.matrix[1][2] = 2;

  int result = s21_mult_matrix(&matrix1, &matrix2, &matrix1_2);

  ck_assert_int_eq(29, matrix1_2.matrix[0][0]);
  ck_assert_int_eq(41, matrix1_2.matrix[0][1]);
  ck_assert_int_eq(19, matrix1_2.matrix[0][2]);
  ck_assert_int_eq(17, matrix1_2.matrix[1][0]);
  ck_assert_int_eq(11, matrix1_2.matrix[1][1]);
  ck_assert_int_eq(8, matrix1_2.matrix[1][2]);
  ck_assert_int_eq(matrix1_2.rows, 2);
  ck_assert_int_eq(matrix1_2.columns, 3);
  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);
  s21_remove_matrix(&matrix1_2);
}
END_TEST

START_TEST(s21_mult_matrix_test3) {
  matrix_t matrix1;
  matrix_t matrix2;
  matrix_t matrix1_2;

  s21_create_matrix(4, 3, &matrix1);
  s21_create_matrix(2, 3, &matrix2);

  int result = s21_mult_matrix(&matrix1, &matrix2, &matrix1_2);

  ck_assert_int_eq(result, ERR_CALC);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);
}
END_TEST

START_TEST(s21_mult_matrix_test4) {
  matrix_t matrix1;
  matrix_t matrix2;
  matrix_t matrix1_2;

  s21_create_matrix(5, 2, &matrix1);
  matrix1.matrix[0][0] = 1;
  matrix1.matrix[0][1] = 2;
  matrix1.matrix[1][0] = 3;
  matrix1.matrix[1][1] = 4;
  matrix1.matrix[2][0] = 5;
  matrix1.matrix[2][1] = 6;
  matrix1.matrix[3][0] = 7;
  matrix1.matrix[3][1] = 8;
  matrix1.matrix[4][0] = 9;
  matrix1.matrix[4][1] = 10;

  s21_create_matrix(2, 3, &matrix2);
  matrix2.matrix[0][0] = 6;
  matrix2.matrix[0][1] = 0;
  matrix2.matrix[0][2] = 8;
  matrix2.matrix[1][0] = 2;
  matrix2.matrix[1][1] = 3;
  matrix2.matrix[1][2] = 9;

  int result = s21_mult_matrix(&matrix1, &matrix2, &matrix1_2);

  ck_assert_int_eq(10, matrix1_2.matrix[0][0]);
  ck_assert_int_eq(6, matrix1_2.matrix[0][1]);
  ck_assert_int_eq(26, matrix1_2.matrix[0][2]);
  ck_assert_int_eq(26, matrix1_2.matrix[1][0]);
  ck_assert_int_eq(12, matrix1_2.matrix[1][1]);
  ck_assert_int_eq(60, matrix1_2.matrix[1][2]);
  ck_assert_int_eq(42, matrix1_2.matrix[2][0]);
  ck_assert_int_eq(18, matrix1_2.matrix[2][1]);
  ck_assert_int_eq(94, matrix1_2.matrix[2][2]);
  ck_assert_int_eq(58, matrix1_2.matrix[3][0]);
  ck_assert_int_eq(24, matrix1_2.matrix[3][1]);
  ck_assert_int_eq(128, matrix1_2.matrix[3][2]);
  ck_assert_int_eq(74, matrix1_2.matrix[4][0]);
  ck_assert_int_eq(30, matrix1_2.matrix[4][1]);
  ck_assert_int_eq(162, matrix1_2.matrix[4][2]);

  ck_assert_int_eq(matrix1_2.rows, 5);
  ck_assert_int_eq(matrix1_2.columns, 3);
  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix2);
  s21_remove_matrix(&matrix1_2);
}
END_TEST

Suite* s21_mult_matrix_tests(void) {
  Suite* s;
  TCase* tc;
  s = suite_create("s21_mult_matrix");
  tc = tcase_create("case_mult_matrix");
  tcase_add_test(tc, s21_mult_matrix_test2);
  tcase_add_test(tc, s21_mult_matrix_test3);
  tcase_add_test(tc, s21_mult_matrix_test4);

  suite_add_tcase(s, tc);
  return s;
}

START_TEST(s21_transpose_test1) {
  matrix_t matrix1;
  matrix_t matrix_res;

  s21_create_matrix(2, 3, &matrix1);
  matrix1.matrix[0][0] = 1;
  matrix1.matrix[0][1] = -2;
  matrix1.matrix[0][2] = 4;
  matrix1.matrix[1][0] = 5;
  matrix1.matrix[1][1] = 0;
  matrix1.matrix[1][2] = 7;

  int result = s21_transpose(&matrix1, &matrix_res);

  ck_assert_int_eq(1, matrix_res.matrix[0][0]);
  ck_assert_int_eq(5, matrix_res.matrix[0][1]);
  ck_assert_int_eq(-2, matrix_res.matrix[1][0]);
  ck_assert_int_eq(0, matrix_res.matrix[1][1]);
  ck_assert_int_eq(4, matrix_res.matrix[2][0]);
  ck_assert_int_eq(7, matrix_res.matrix[2][1]);
  ck_assert_int_eq(matrix_res.rows, 3);
  ck_assert_int_eq(matrix_res.columns, 2);
  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix_res);
}
END_TEST

START_TEST(s21_transpose_test2) {
  matrix_t matrix1;
  matrix_t matrix_res;

  s21_create_matrix(3, 1, &matrix1);
  matrix1.matrix[0][0] = 1;
  matrix1.matrix[1][0] = -2;
  matrix1.matrix[2][0] = 4;

  int result = s21_transpose(&matrix1, &matrix_res);

  ck_assert_int_eq(1, matrix_res.matrix[0][0]);
  ck_assert_int_eq(-2, matrix_res.matrix[0][1]);
  ck_assert_int_eq(4, matrix_res.matrix[0][2]);
  ck_assert_int_eq(matrix_res.rows, 1);
  ck_assert_int_eq(matrix_res.columns, 3);
  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix_res);
}
END_TEST

Suite* s21_transpose_tests(void) {
  Suite* s;
  TCase* tc;
  s = suite_create("s21_transpose");
  tc = tcase_create("case_transpose");
  tcase_add_test(tc, s21_transpose_test1);
  tcase_add_test(tc, s21_transpose_test2);

  suite_add_tcase(s, tc);
  return s;
}

START_TEST(s21_calc_complements_test1) {
  matrix_t matrix1;
  matrix_t matrix_res;
  s21_create_matrix(2, 2, &matrix1);
  matrix1.matrix[0][0] = 1;
  matrix1.matrix[0][1] = 2;
  matrix1.matrix[1][0] = 3;
  matrix1.matrix[1][1] = 4;

  int result = s21_calc_complements(&matrix1, &matrix_res);

  ck_assert_int_eq(4, matrix_res.matrix[0][0]);
  ck_assert_int_eq(-3, matrix_res.matrix[0][1]);
  ck_assert_int_eq(-2, matrix_res.matrix[1][0]);
  ck_assert_int_eq(1, matrix_res.matrix[1][1]);
  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix_res);
}
END_TEST

START_TEST(s21_calc_complements_test3) {
  matrix_t matrix1;
  matrix_t matrix_res;
  s21_create_matrix(3, 2, &matrix1);

  int result = s21_calc_complements(&matrix1, &matrix_res);
  ck_assert_int_eq(result, ERR_CALC);

  s21_remove_matrix(&matrix1);
}
END_TEST

Suite* s21_calc_complements_tests(void) {
  Suite* s;
  TCase* tc;
  s = suite_create("s21_calc_complements");
  tc = tcase_create("case_calc_complements");
  tcase_add_test(tc, s21_calc_complements_test1);
  tcase_add_test(tc, s21_calc_complements_test3);

  suite_add_tcase(s, tc);
  return s;
}

START_TEST(s21_determinant_test1) {
  matrix_t matrix1;
  double matrix_res;
  s21_create_matrix(2, 2, &matrix1);
  matrix1.matrix[0][0] = 1;
  matrix1.matrix[0][1] = 0;
  matrix1.matrix[1][0] = 2;
  matrix1.matrix[1][1] = 5;

  int result = s21_determinant(&matrix1, &matrix_res);

  ck_assert_int_eq(matrix_res, 5);
  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
}
END_TEST

START_TEST(s21_determinant_test2) {
  matrix_t matrix1;
  double matrix_res;
  s21_create_matrix(3, 3, &matrix1);
  matrix1.matrix[0][0] = 1;
  matrix1.matrix[0][1] = 0;
  matrix1.matrix[0][2] = 7;
  matrix1.matrix[1][0] = 2;
  matrix1.matrix[1][1] = 5;
  matrix1.matrix[1][2] = 11;
  matrix1.matrix[2][0] = 4;
  matrix1.matrix[2][1] = 17;
  matrix1.matrix[2][2] = 11;

  int result = s21_determinant(&matrix1, &matrix_res);

  ck_assert_int_eq(matrix_res, -34);
  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
}
END_TEST

START_TEST(s21_determinant_test3) {
  matrix_t matrix1;
  double matrix_res;
  s21_create_matrix(2, 1, &matrix1);
  matrix1.matrix[0][0] = 1;
  matrix1.matrix[1][0] = 2;

  int result = s21_determinant(&matrix1, &matrix_res);

  ck_assert_int_eq(result, ERR_CALC);

  s21_remove_matrix(&matrix1);
}
END_TEST

Suite* s21_determinant_tests(void) {
  Suite* s;
  TCase* tc;
  s = suite_create("s21_determinant");
  tc = tcase_create("case_determinant");
  tcase_add_test(tc, s21_determinant_test1);
  tcase_add_test(tc, s21_determinant_test2);
  tcase_add_test(tc, s21_determinant_test3);

  suite_add_tcase(s, tc);
  return s;
}

START_TEST(s21_inverse_matrix_test1) {
  matrix_t matrix1;
  matrix_t matrix_res;
  s21_create_matrix(3, 3, &matrix1);
  matrix1.matrix[0][0] = 12;
  matrix1.matrix[0][1] = 2;
  matrix1.matrix[0][2] = 3;
  matrix1.matrix[1][0] = 4;
  matrix1.matrix[1][1] = 5;
  matrix1.matrix[1][2] = 6;
  matrix1.matrix[2][0] = 7;
  matrix1.matrix[2][1] = 8;
  matrix1.matrix[2][2] = 9;

  int result = s21_inverse_matrix(&matrix1, &matrix_res);
  ck_assert_double_eq_tol(1.0 / 11, matrix_res.matrix[0][0], 1e-7);
  ck_assert_double_eq_tol(-(2.0 / 11), matrix_res.matrix[0][1], 1e-7);
  ck_assert_double_eq_tol(1.0 / 11, matrix_res.matrix[0][2], 1e-7);
  ck_assert_double_eq_tol(-(2.0 / 11), matrix_res.matrix[1][0], 1e-7);
  ck_assert_double_eq_tol(-(29.0 / 11), matrix_res.matrix[1][1], 1e-7);
  ck_assert_double_eq_tol(20.0 / 11, matrix_res.matrix[1][2], 1e-7);
  ck_assert_double_eq_tol(1.0 / 11, matrix_res.matrix[2][0], 1e-7);
  ck_assert_double_eq_tol(82.0 / 33, matrix_res.matrix[2][1], 1e-7);
  ck_assert_double_eq_tol(-(52.0 / 33), matrix_res.matrix[2][2], 1e-7);

  ck_assert_int_eq(result, OK);

  s21_remove_matrix(&matrix1);
  s21_remove_matrix(&matrix_res);
}
END_TEST

START_TEST(s21_inverse_matrix_test2) {
  matrix_t matrix1;
  matrix_t matrix_res;
  s21_create_matrix(3, 3, &matrix1);
  matrix1.matrix[0][0] = 1;
  matrix1.matrix[0][1] = 2;
  matrix1.matrix[0][2] = 3;
  matrix1.matrix[1][0] = 4;
  matrix1.matrix[1][1] = 5;
  matrix1.matrix[1][2] = 6;
  matrix1.matrix[2][0] = 7;
  matrix1.matrix[2][1] = 8;
  matrix1.matrix[2][2] = 9;

  int result = s21_inverse_matrix(&matrix1, &matrix_res);
  ck_assert_int_eq(result, ERR_CALC);

  s21_remove_matrix(&matrix1);
}
END_TEST

START_TEST(s21_inverse_matrix_test3) {
  matrix_t m = {0};
  matrix_t result = {0};
  s21_create_matrix(3, 3, &m);

  int code = s21_inverse_matrix(&m, &result);
  ck_assert_int_eq(code, ERR_CALC);
  s21_remove_matrix(&m);
  s21_remove_matrix(&result);
}
END_TEST

START_TEST(s21_inverse_matrix_test4) {
  matrix_t matrix1;
  matrix_t matrix_res;
  s21_create_matrix(4, 3, &matrix1);

  int result = s21_inverse_matrix(&matrix1, &matrix_res);
  ck_assert_int_eq(result, ERR_CALC);

  s21_remove_matrix(&matrix1);
}
END_TEST

START_TEST(s21_inverse_matrix_test5) {
  matrix_t m = {0};
  matrix_t result = {0};
  s21_create_matrix(1, 4, &m);
  int code = s21_inverse_matrix(&m, &result);
  ck_assert_int_eq(code, ERR_CALC);
  s21_remove_matrix(&m);
  s21_remove_matrix(&result);
}
END_TEST

START_TEST(test_s21_inverse_matrix_3) {
  matrix_t temp, result;
  s21_create_matrix(0, 0, &temp);
  int num = 1;
  for (int i = 0; i < temp.rows; i++) {
    for (int j = 0; j < temp.columns; j++, num++) {
      temp.matrix[i][j] = num;
    }
  }

  int result_status = s21_inverse_matrix(&temp, &result);
  ck_assert_int_eq(result_status, 1);
}
END_TEST

Suite* s21_inverse_matrix_tests(void) {
  Suite* s;
  TCase* tc;
  s = suite_create("s21_inverse_matrix");
  tc = tcase_create("case_inverse_matrix");
  tcase_add_test(tc, s21_inverse_matrix_test1);
  tcase_add_test(tc, s21_inverse_matrix_test2);
  tcase_add_test(tc, s21_inverse_matrix_test3);
  tcase_add_test(tc, s21_inverse_matrix_test4);
  tcase_add_test(tc, s21_inverse_matrix_test5);
  tcase_add_test(tc, test_s21_inverse_matrix_3);

  suite_add_tcase(s, tc);
  return s;
}

int main(void) {
  int failed = 0;
  Suite* s21_matrix_tests[] = {s21_create_matrix_tests(),
                               s21_eq_matrix_tests(),
                               s21_sum_matrix_tests(),
                               s21_sub_matrix_tests(),
                               s21_mult_number_tests(),
                               s21_mult_matrix_tests(),
                               s21_transpose_tests(),
                               s21_calc_complements_tests(),
                               s21_determinant_tests(),
                               s21_inverse_matrix_tests(),
                               NULL};

  for (int i = 0; s21_matrix_tests[i] != NULL; i++) {  // (&& failed == 0)
    SRunner* sr = srunner_create(s21_matrix_tests[i]);

    srunner_set_fork_status(sr, CK_NOFORK);
    printf("\n");

    srunner_run_all(sr, CK_NORMAL);

    failed += srunner_ntests_failed(sr);
    srunner_free(sr);
    printf("\n");
  }
  printf("========= FAILED: %d =========\n", failed);

  return failed == 0 ? 0 : 1;
}
