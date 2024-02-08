#!/bin/bash

# Test_1
./s21_cat test.txt text.txt > s21_cat_test
cat test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_2
./s21_cat -b test.txt text.txt > s21_cat_test
cat -b test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_3
./s21_cat --number-nonblank test.txt text.txt > s21_cat_test
cat -b test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_4
./s21_cat -e -v test.txt text.txt > s21_cat_test
cat -ev test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_5
./s21_cat -e test.txt text.txt > s21_cat_test
cat -e test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_6
./s21_cat -n test.txt text.txt > s21_cat_test
cat -n test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_7
./s21_cat --number test.txt text.txt > s21_cat_test
cat -n test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_8
./s21_cat -s test.txt text.txt > s21_cat_test
cat -s test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_9
./s21_cat --squeeze-blank test.txt text.txt > s21_cat_test
cat -s test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_10
./s21_cat -v -b -s -n test.txt text.txt > s21_cat_test
cat -v -b -s -n test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_11
./s21_cat -t -v test.txt text.txt > s21_cat_test
cat -t -v test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_12
./s21_cat -tve test.txt text.txt > s21_cat_test
cat -tve test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_13
./s21_cat -tse test.txt text.txt > s21_cat_test
cat -tse test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_14
./s21_cat -b -n test.txt text.txt > s21_cat_test
cat -b -n test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_15
./s21_cat -sv test.txt text.txt > s21_cat_test
cat -sv test.txt text.txt > cat
diff -s s21_cat_test cat

# Test_15
./s21_cat -v test.txt text.txt > s21_cat_test
cat -v test.txt text.txt > cat
diff -s s21_cat_test cat



