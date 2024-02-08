#include <check.h>

#include "s21_SmartCalc.h"

START_TEST(s21_smartCalc_test_1) {
  double res = 0.0;
  char *str = "200*7-323+50/0.2+2*(123-112)";
  int error = 0;
  res = s21_parser(str, &error, 0);
  ck_assert_int_eq(res, 1349);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_3) {
  double res = 0.0;
  char *str = "log(99)/0.2+ln(35)/sqrt(81)";
  int error = 0;
  res = s21_parser(str, &error, 0);
  ck_assert_int_eq(res, 10.3732146);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_4) {
  double res = 0.0;
  char *str = "100-2^(2^3)";
  int error = 0;
  res = s21_parser(str, &error, 0);
  ck_assert_int_eq(res, 100 - pow(2, pow(2, 3)));
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_5) {
  char *str = "(*)|(*))";
  int error = 0;
  s21_parser(str, &error, 0);
  ck_assert_int_eq(error, 1);
}
END_TEST

START_TEST(s21_smartCalc_test_6) {
  double res = 0.0;
  char *str = "10.+2^(2^(3))";
  int error = 0;
  res = s21_parser(str, &error, 0);

  ck_assert_double_eq(res, 10 + pow(2, pow(2, 3)));
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_7) {
  double res = 0.0;
  char *str = "cos(99.)/2^(sin(0.5))*10";
  int error = 0;
  res = s21_parser(str, &error, 0);

  ck_assert_double_eq(res, cos(99.) / pow(2, sin(0.5)) * 10);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_8) {
  double res = 0.0;
  char *str = "7.2^(sin(55.)-9^(cos(3.6))/7)";
  int error = 0;
  res = s21_parser(str, &error, 0);

  ck_assert_double_eq(res, pow(7.2, (sin(55.) - pow(9, cos(3.6)) / 7)));
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_9) {
  double res = 0.0;
  char *str = "tan(2)-cos(45)+2^(sin(52))*5";
  int error = 0;
  res = s21_parser(str, &error, 0);

  ck_assert_double_eq(res, tan(2) - cos(45) + pow(2, sin(52)) * 5);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_10) {
  double res = 0.0;
  char *str = "ln(140)*log(49)";
  int error = 0;
  res = s21_parser(str, &error, 0);

  ck_assert_double_eq(res, log(140) * log10(49));
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_11) {
  double res = 0.0;
  char *str = "sin(0.25)-cos(22.5)+tan(0.5)";
  int error = 0;
  res = s21_parser(str, &error, 0);
  ck_assert_double_eq_tol(res, sin(0.25) - cos(22.5) + tan(0.5), 1e-7);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_12) {
  double res = 0.0;
  char *str = "-1*2.5/sin(360)";
  int error = 0;
  res = s21_parser(str, &error, 0);

  ck_assert_double_eq(res, -1 * 2.5 / sin(360));
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_13) {
  double res = 0.0;
  char *str = "-tan(45)*cos(45)/2^(sin(33))*(-5)";
  int error = 0;
  res = s21_parser(str, &error, 0);

  ck_assert_double_eq(res, -tan(45) * cos(45) / pow(2, (sin(33))) * -5);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_14) {
  double res = 0.0;
  char *str = "0.5/0";
  int error = 0;
  res = s21_parser(str, &error, 0);

  ck_assert_double_infinite(res);
  ck_assert_int_eq(error, 1);
}
END_TEST

START_TEST(s21_smartCalc_test_15) {
  char *str = "(5)+()+(^)";
  int error = 0;
  s21_parser(str, &error, 0);
  ck_assert_int_eq(error, 1);
}
END_TEST

START_TEST(s21_smartCalc_test_16) {
  double res = 0.0;
  char *str = "-99mod8-58mod23/88^(32mod5)";
  int error = 0;
  res = s21_parser(str, &error, 0);
  ck_assert_double_eq(res, -99 % 8 - 58 % 23 / pow(88, (32 % 5)));
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_17) {
  double res = 0.0;
  char *str = "-81/(-9)-(-23)-(-25/0.5)*(-123-112)";
  int error = 0;
  res = s21_parser(str, &error, 0);
  ck_assert_int_eq(res, -11718);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_18) {
  double res = 0.0;
  char *str = "sqrt(25)-log(5)*100+ln(5)";
  int error = 0;
  res = s21_parser(str, &error, 0);
  ck_assert_int_eq(res, -63.2875625);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_19) {
  double res = 0.0;
  char *str = "acos(10mod3)-asin(1)+atan(1)";
  int error = 0;
  res = s21_parser(str, &error, 0);
  ck_assert_int_eq(res, -0.7853981);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_20) {
  double res = 0.0;
  char *str = "-10mod3";
  int error = 0;
  res = s21_parser(str, &error, 0);
  ck_assert_int_eq(res, -1);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_21) {
  double res = 0.0;
  char *str =
      "15/(7-(1+1))*3-(2+(1+1))*15/(7-(200+1))*3-(2+(1+1))*(15/"
      "(7-(1+1))*3-(2+(1+1))+15/(7-(1+1))*3-(2+(1+1)))";
  int error = 0;
  res = s21_parser(str, &error, 0);
  ck_assert_int_eq(res, -30.0721649485);
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_23) {
  double res = 0.0;
  char *str = "-ln(140)*log(49)";
  int error = 0;
  res = s21_parser(str, &error, 0);

  ck_assert_double_eq(res, -log(140) * log10(49));
  ck_assert_int_eq(error, 0);
}
END_TEST

START_TEST(s21_smartCalc_test_24) {
  double res = 0.0;
  char *str = "2^2";
  int error = 0;
  res = s21_parser(str, &error, 0);

  ck_assert_double_eq(res, pow(2, 2));
  ck_assert_int_eq(error, 0);
}
END_TEST

Suite *s21_SmartCalc_suite() {
  Suite *s;

  s = suite_create("s21_SmartCalc");
  TCase *tc = tcase_create("Core");

  tcase_add_test(tc, s21_smartCalc_test_1);
  tcase_add_test(tc, s21_smartCalc_test_3);
  tcase_add_test(tc, s21_smartCalc_test_4);
  tcase_add_test(tc, s21_smartCalc_test_5);
  tcase_add_test(tc, s21_smartCalc_test_6);
  tcase_add_test(tc, s21_smartCalc_test_7);
  tcase_add_test(tc, s21_smartCalc_test_8);
  tcase_add_test(tc, s21_smartCalc_test_9);
  tcase_add_test(tc, s21_smartCalc_test_10);
  tcase_add_test(tc, s21_smartCalc_test_11);
  tcase_add_test(tc, s21_smartCalc_test_12);
  tcase_add_test(tc, s21_smartCalc_test_13);
  tcase_add_test(tc, s21_smartCalc_test_14);
  tcase_add_test(tc, s21_smartCalc_test_15);
  tcase_add_test(tc, s21_smartCalc_test_16);
  tcase_add_test(tc, s21_smartCalc_test_17);
  tcase_add_test(tc, s21_smartCalc_test_18);
  tcase_add_test(tc, s21_smartCalc_test_19);
  tcase_add_test(tc, s21_smartCalc_test_20);
  tcase_add_test(tc, s21_smartCalc_test_21);
  tcase_add_test(tc, s21_smartCalc_test_23);
  tcase_add_test(tc, s21_smartCalc_test_24);

  suite_add_tcase(s, tc);

  return s;
}

int main() {
  int failed = 0;
  Suite *suite = s21_SmartCalc_suite();
  SRunner *suite_runner = srunner_create(suite);

  srunner_set_fork_status(suite_runner, CK_NOFORK);
  srunner_run_all(suite_runner, CK_VERBOSE);

  failed = srunner_ntests_failed(suite_runner);
  srunner_free(suite_runner);

  return failed == 0 ? EXIT_SUCCESS : EXIT_FAILURE;
}
