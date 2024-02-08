#ifndef S21_SMARTCALC_H_
#define S21_SMARTCALC_H_

#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

enum {
  NUM = 1,
  PLUS,
  MINUS,
  MUL,
  DIV,
  POW,
  SQRT,
  COS,
  SIN,
  TAN,
  LN,
  LOG,
  ACOS,
  ASIN,
  ATAN,
  MOD,
  BRACKET_OPEN,
  BRACKET_CLOSE
};

typedef struct Stack {
  double number;
  int priority;
  int lexeme;
  // int *error;
  struct Stack *next;
} Stack;

void s21_push(Stack **top, double number, int prioritet, int lexeme);
double s21_pop(Stack **top);
int s21_check_str(const char *src1, char *src2);
int s21_check(char *str);
int s21_check_bracket(char *str);
int s21_check_sign(char *str);
int s21_check_functions(char *str, int *i, char ch);
int s21_convert_to_double(int *i, char *str, double *number);
int s21_get_operator(int *operat, int *i, char *str);
void s21_priority_and_lexeme(int *operat, int priority, int lexeme);
int s21_functions(int *operat, int *i, char *str);
void s21_polish_notation(Stack **numbers, Stack **operators, int *operat);
void s21_operation_in_stack(Stack **numbers, Stack **operators);
void s21_binary_operations(Stack **numbers, Stack **operators,
                           double *result_of_operation);
void s21_unary_operations(Stack **numbers, Stack **operators,
                          double *result_of_operation);
double s21_calculate(Stack **numbers, Stack **operators, int *error);
double s21_parser(const char *str, int *error, double x);
// void zeroing_stack(Stack **data);
// void s21_clearStack(Stack *stack);
void s21_zeroing_info(Stack *stack);
int s21_credit_annuity(double sum_of_credit, double duration_of_credit,
                       double annual_interest_rate, double *month_pay,
                       double *over_pay, double *all_sum_of_pay);
int s21_credit_differentiated(double sum_of_credit, double duration_of_credit,
                              double annual_interest_rate,
                              double *month_pay_first, double *month_pay_last,
                              double *over_pay, double *all_sum_of_pay);
double s21_replenishment_list(double *replenishment1,
                              double const *replenishment2,
                              int count_replenishment, int count_month);
double s21_withdrawals_list(double *withdrawals1, double const *withdrawals2,
                            int count_withdrawals, int count_month,
                            double all_sum);
int s21_deposit(double *replenishment1, double *replenishment2,
                int count_replenishment, double *withdrawals1,
                double *withdrawals2, int count_withdrawals, double all_sum,
                double duration_of_credit, double interest_rate, double tax,
                int period_s, int capitalization, double *all_interest,
                double *all_tax, double *total_sum);

#endif  // S21_SMARTCALC_H_
