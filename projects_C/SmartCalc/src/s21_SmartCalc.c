#include "s21_SmartCalc.h"

void s21_push(Stack** top, double number, int priority, int lexeme) {
  Stack* temp = (Stack*)malloc(sizeof(Stack));
  // Stack* tmp = calloc(1, sizeof(Stack));
  if (temp == NULL) {
    exit(1);
  }
  temp->number = number;
  temp->lexeme = lexeme;
  temp->priority = priority;
  temp->next = (*top);
  (*top) = temp;
}

double s21_pop(Stack** top) {
  Stack* previous = NULL;
  double num = 0.0;
  if (top != NULL) {
    previous = *top;
    num = previous->number;
    *top = (*top)->next;
    free(previous);
  }
  return num;
}

int s21_check_str(const char* src1, char* src2) {
  int j = 0, dot = 0, error = 0;
  char buff[256] = "";
  for (int i = 0; i < (int)strlen(src1); i++) {
    error = 0;
    if (src1[i] == '.') {
      if (i == 0 || i == ((int)strlen(src1) - 1)) {  // dot = 0;
        error = 1;
      } else {
        dot++;
        if (dot > 1) {  // dot = 0;
          error = 1;
        }
        if ((dot == 1) &&
            ((!strchr("0123456789", src1[i - 1])) ||
             (!strchr("0123456789.", src1[i + 1])))) {  // dot = 0;
          error = 1;
        }
      }
    }
    if (!strchr("0123456789.", src1[i]) && dot > 0) {
      dot = 0;
    }
    if (!error) {
      if (src1[i] == '.')
        buff[j] = '.';
      else
        buff[j] = src1[i];
      j++;
    }
  }
  buff[j + 1] = '\0';
  j = 0;
  for (int i = 0; i < (int)strlen(buff); i++) {
    error = 0;
    if (buff[i] == '.') {
      if ((!strchr("0123456789", buff[i - 1])) ||
          (!strchr("0123456789.", buff[i + 1])))
        error = 1;
      if (i == ((int)strlen(buff) - 1)) {
        error = 1;
      }
    }
    if (!error) {
      src2[j] = buff[i];
      j++;
    }
  }
  src2[j + 1] = '\0';
  return error;
}

int s21_check(char* str) {
  int flag_check = 1;
  for (int i = 0; i < (int)strlen(str) && flag_check; i++) {
    if (!strchr("^x1234567890.()+-*/^acstl", str[i])) {
      if (i == 0 || i == (int)strlen(str) - 1) {
        flag_check = 0;
      }
    }
  }
  if (flag_check) {
    flag_check *= s21_check_bracket(str);
  }
  if (flag_check) {
    flag_check *= s21_check_sign(str);
  }
  if (flag_check) {
    for (int i = 0; i < (int)strlen(str); i++) {
      if (strchr("acstl", str[i])) {
        flag_check *= s21_check_functions(str, &i, str[i]);
      }
    }
  }
  return flag_check;
}

/*int s21_check_bracket(char* str) {
  int i = 0;
  int bracket = 0;
  int err = 0;
  int amount_of_numbers = 0;

  if (str[i] == '(') {
    bracket++;
    i++;
    for (; str[i] && bracket != 0; i++) {
      if (str[i] >= '0' && str[i] <= '9') {
        amount_of_numbers++;
      } else if (str[i] == '(') {
        bracket++;
      } else if (str[i] == ')') {
        bracket--;
      }
    }
  } else {
    err = 1;
  }
  if (amount_of_numbers == 0 || bracket != 0) {
    err = 1;
  }
  return err;
}*/

int s21_check_sign(char* str) {
  int flag_check = 1;
  for (size_t i = 0; i < strlen(str) && flag_check; i++) {
    if (strchr("^+-*/.^", str[i])) {
      if (i == 0 &&
          (str[i] == '*' || str[i] == '/' || str[i] == '^' || str[i] == '.')) {
        flag_check = 0;
      }
      if ((str[i] != '+' && str[i] != '-') && str[i] == str[i + 1]) {
        flag_check = 0;
      }
      if (i == strlen(str) - 1) {
        flag_check = 0;
      }
      if (str[i - 1] == '(' && str[i + 1] == ')') {
        flag_check = 0;
      }
      //    if (str[i] == '^' && str[i + 1] != '(') {
      //    flag_check = 0;
      //   }
      if (strchr("^+-*/.^", str[i + 1])) {
        flag_check = 0;
      }
    }
  }
  return flag_check;
}

int s21_check_functions(char* str, int* i, char ch) {
  int flag_check = 1, j = 0;
  char functions[6] = "";
  if (str[*i] == 'a' || (str[*i] == 's' && str[*i + 1] == 'q')) {
    for (j = 0; j < 5; j++) {
      functions[j] = str[*i];
      *i += 1;
    }
    functions[j + 1] = '\0';
  } else {
    if (ch == 'c' || ch == 's' || ch == 't' ||
        (ch == 'l' && str[*i + 1] == 'o')) {
      for (int j = 0; j < 4; j++) {
        functions[j] = str[*i];
        *i += 1;
      }
    } else if (ch == 'l' && str[*i + 1] == 'n') {
      for (int j = 0; j < 3; j++) {
        functions[j] = str[*i];
        *i += 1;
      }
    }
  }
  if (strlen(functions) == 3) {
    if (strcmp(functions, "ln(") != 0) {
      flag_check = 0;
    }
  }
  if (strlen(functions) == 4) {
    if (strcmp(functions, "cos(") != 0 && strcmp(functions, "sin(") != 0 &&
        strcmp(functions, "tan(") != 0 && strcmp(functions, "log(") != 0) {
      flag_check = 0;
    }
  }
  if (strlen(functions) == 5) {
    if (strcmp(functions, "acos(") != 0 && strcmp(functions, "asin(") != 0 &&
        strcmp(functions, "atan(") != 0 && strcmp(functions, "sqrt(") != 0) {
      flag_check = 0;
    }
  }
  return flag_check;
}

int s21_convert_to_double(int* i, char* str, double* number) {
  char buff[256] = {'\0'};
  int flag_check = 1, j = 0;
  *number = 0;
  while (strchr("1234567890.", str[*i]) && flag_check) {
    buff[j] = str[*i];
    *i += 1;
    j++;
    if (*i >= (int)strlen(str)) {
      flag_check = 0;
    }
  }
  buff[j + 1] = '\0';
  *number = atof(buff);
  return flag_check;
}

void s21_priority_and_lexeme(int* operator, int priority, int lexeme) {
  operator[0] = priority;
  operator[1] = lexeme;
}

int s21_functions(int* operator, int * i, char* str) {
  int flag = 1, j = 0;
  char function[5] = "";
  if (str[*i] == 'a' || (str[*i] == 's' && str[*i + 1] == 'q')) {
    for (j = 0; j < 4; j++) {
      function[j] = str[*i];
      *i += 1;
    }
    function[j + 1] = '\0';
  } else {
    if (str[*i] == 'c' || str[*i] == 's' || str[*i] == 't' || str[*i] == 'm' ||
        (str[*i] == 'l' && str[*i + 1] == 'o')) {
      for (int j = 0; j < 3; j++) {
        function[j] = str[*i];
        *i += 1;
      }
    } else if (str[*i] == 'l' && str[*i + 1] == 'n') {
      for (int j = 0; j < 2; j++) {
        function[j] = str[*i];
        *i += 1;
      }
    } else
      flag = 0;
  }
  if (strcmp(function, "ln") == 0) {
    s21_priority_and_lexeme(operator, 4, LN);
  }
  if (strcmp(function, "log") == 0) {
    s21_priority_and_lexeme(operator, 4, LOG);
  }
  if (strcmp(function, "cos") == 0) {
    s21_priority_and_lexeme(operator, 4, COS);
  }
  if (strcmp(function, "sin") == 0) {
    s21_priority_and_lexeme(operator, 4, SIN);
  }
  if (strcmp(function, "tan") == 0) {
    s21_priority_and_lexeme(operator, 4, TAN);
  }
  if (strcmp(function, "mod") == 0) {
    s21_priority_and_lexeme(operator, 2, MOD);
  }
  if (strcmp(function, "acos") == 0) {
    s21_priority_and_lexeme(operator, 4, ACOS);
  }
  if (strcmp(function, "asin") == 0) {
    s21_priority_and_lexeme(operator, 4, ASIN);
  }
  if (strcmp(function, "atan") == 0) {
    s21_priority_and_lexeme(operator, 4, ATAN);
  }
  if (strcmp(function, "sqrt") == 0) {
    s21_priority_and_lexeme(operator, 4, SQRT);
  }
  if (*i >= (int)strlen(str)) {
    flag = 0;
  }
  return flag;
}

int s21_get_operator(int* operator, int * i, char* str) {
  int flag = 1;

  if (strchr("acstlm", str[*i]))
    flag = s21_functions(operator, i, str);
  else {
    if (str[*i] == '+') {
      s21_priority_and_lexeme(operator, 1, PLUS);
    }
    if (str[*i] == '-') {
      s21_priority_and_lexeme(operator, 1, MINUS);
    }
    if (str[*i] == '*') {
      s21_priority_and_lexeme(operator, 2, MUL);
    }
    if (str[*i] == '/') {
      s21_priority_and_lexeme(operator, 2, DIV);
    }
    if (str[*i] == '^') {
      s21_priority_and_lexeme(operator, 3, POW);
    }
    if (str[*i] == '(') {
      s21_priority_and_lexeme(operator, 5, BRACKET_OPEN);
    }
    if (str[*i] == ')') {
      s21_priority_and_lexeme(operator, 5, BRACKET_CLOSE);
    }
    *i += 1;
  }
  if (*i >= (int)strlen(str)) {
    flag = 0;
  }
  return flag;
}

void s21_polish_notation(Stack** numbers, Stack** operators, int* operator) {
  int stopFlag = 1;
  while (stopFlag) {
    if ((*operators == NULL || (operator[0]>(*operators)->priority) ||
         (*operators)->lexeme == BRACKET_OPEN) &&
        (operator[1] != BRACKET_CLOSE)) {
      s21_push(operators, 0, operator[0], operator[1]);
      stopFlag = 0;
    } else {
      if (operator[1] == BRACKET_CLOSE) {
        s21_push(operators, 0, operator[0], operator[1]);
        s21_pop(operators);
        while (stopFlag) {
          if ((*operators)->lexeme == BRACKET_OPEN) {
            s21_pop(operators);
            stopFlag = 0;
          }
          if (stopFlag) {
            s21_operation_in_stack(numbers, operators);
          }
        }
      } else {
        s21_operation_in_stack(numbers, operators);
      }
    }
  }
}

void s21_operation_in_stack(Stack** numbers, Stack** operators) {
  double result_of_operation = 0;
  int operator=(*operators)->lexeme;
  if (operator== PLUS || operator== MINUS || operator== MUL || operator== DIV ||
      operator== MOD ||
      operator== POW) {
    s21_binary_operations(numbers, operators, &result_of_operation);
  }
  if (operator== COS || operator== SIN || operator== TAN || operator== LN ||
      operator== LOG ||
      operator== ACOS ||
      operator== ASIN ||
      operator== ATAN ||
      operator== SQRT) {
    s21_unary_operations(numbers, operators, &result_of_operation);
  }
  s21_push(numbers, result_of_operation, 0, NUM);
  s21_pop(operators);
}

void s21_binary_operations(Stack** numbers, Stack** operators,
                           double* result_of_operation) {
  double number1 = 0, number2 = 0;
  number2 = s21_pop(numbers);
  number1 = s21_pop(numbers);
  if ((*operators)->lexeme == PLUS) {
    *result_of_operation = number1 + number2;
  }
  if ((*operators)->lexeme == MINUS) {
    *result_of_operation = number1 - number2;
  }
  if ((*operators)->lexeme == MUL) {
    *result_of_operation = number1 * number2;
  }
  if ((*operators)->lexeme == DIV) {
    *result_of_operation = number1 / number2;
  }
  if ((*operators)->lexeme == MOD) {
    *result_of_operation = fmod(number1, number2);
  }
  if ((*operators)->lexeme == POW) {
    *result_of_operation = pow(number1, number2);
  }
}

void s21_unary_operations(Stack** numbers, Stack** operators,
                          double* result_of_operation) {
  double number = 0;
  number = s21_pop(numbers);
  if ((*operators)->lexeme == COS) {
    *result_of_operation = cos(number);
  }
  if ((*operators)->lexeme == SIN) {
    *result_of_operation = sin(number);
  }
  if ((*operators)->lexeme == TAN) {
    *result_of_operation = tan(number);
  }
  if ((*operators)->lexeme == LN) {
    *result_of_operation = log(number);
  }
  if ((*operators)->lexeme == LOG) {
    *result_of_operation = log10(number);
  }
  if ((*operators)->lexeme == ACOS) {
    *result_of_operation = acos(number);
  }
  if ((*operators)->lexeme == ASIN) {
    *result_of_operation = asin(number);
  }
  if ((*operators)->lexeme == ATAN) {
    *result_of_operation = atan(number);
  }
  if ((*operators)->lexeme == SQRT) {
    *result_of_operation = sqrt(number);
  }
}

double s21_calculate(Stack** numbers, Stack** operators, int* error) {
  double result_of_operation = 0;
  int flag = 1;
  while (flag && *error == 0) {
    if ((*numbers == NULL) && (*operators == NULL))
      *error = 2;
    else if (*operators == NULL)
      flag = 0;
    else {
      s21_operation_in_stack(numbers, operators);
      *error = 0;
      if (operators == NULL) {
        flag = 0;
      }
    }
  }
  if (*error != 2) {
    result_of_operation = (*numbers)->number;
  }
  return result_of_operation;
}

double s21_parser(const char* str, int* error, double x) {
  double result = 0;
  *error = 0;
  Stack* top = NULL;
  Stack* operators = NULL;
  int operator[2] = {0};
  char buff[256] = {'\0'};
  int i = 0, flag = 1;
  double val = 0;
  if (s21_check_str(str, buff) != 1) {
    flag *= s21_check(buff);
    if (!flag) *error = 1;
    while (i < (int)strlen(buff) && flag) {
      if (strchr("1234567890x", buff[i])) {
        if (buff[i] == 'x') {
          s21_push(&top, x, 0, NUM);
          i++;
        } else {
          flag *= s21_convert_to_double(&i, buff, &val);
          s21_push(&top, val, 0, NUM);
        }
      } else {
        if ((i == 0 && buff[i] == '-') ||
            (buff[i] == '-' && buff[i - 1] == '(')) {
          s21_push(&top, 0, 0, NUM);
        }
        flag = s21_get_operator(operator, & i, buff);
        s21_polish_notation(&top, &operators, operator);
      }
    }
    if (*error == 0) {
      result = s21_calculate(&top, &operators, error);
    }
    if ((result == INFINITY) || (result == NAN)) {
      *error = 1;
    }
    if (top != NULL) {
      s21_zeroing_info(top);
      free(top);
    }
    if (operators != NULL) {
      s21_zeroing_info(operators);
      free(operators);
    }
  } else {
    printf("error!");
  }
  return result;
}

void s21_zeroing_info(Stack* stack) {
  if (stack) {
    stack->number = 0;
    stack->priority = 0;
    stack->lexeme = 0;
    stack->next = NULL;
  }
}

void s21_pair_bracket(char* str, int* result_flag, int* index) {
  int flag = 0;
  *index += 1;
  while (*index < (int)strlen(str) && !flag && *result_flag == 1) {
    if (str[*index] == ')') {
      flag = 1;
    }
    if (*index == (int)strlen(str) - 1 && !flag) *result_flag = 0;
    if (str[*index] == '(' && result_flag && !flag) {
      s21_pair_bracket(str, result_flag, index);
      if (*index == (int)strlen(str)) *result_flag = 0;
    } else
      *index += 1;
  }
}

int s21_check_bracket(char* str) {
  int result_flag = 1;
  int count_1 = 0;
  int count_2 = 0;
  for (size_t i = 0; i < strlen(str) && result_flag; i++) {
    if (str[i] == '(') {
      if (str[i + 1] == ')' || i == strlen(str) - 1) result_flag = 0;
      count_1++;
    }
    if (str[i] == ')' && result_flag) {
      if (strchr("+-*/.^", str[i - 1]) || i == 0) result_flag = 0;
      count_2++;
    }
  }
  if (count_1 != count_2) result_flag = 0;
  int index = 0;
  while (index < (int)strlen(str) && result_flag) {
    if (str[index] == '(')
      s21_pair_bracket(str, &result_flag, &index);
    else
      index++;
  }
  return result_flag;
}
