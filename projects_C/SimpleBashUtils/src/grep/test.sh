#!/bin/bash

./s21_grep -e qw *.txt > s21_grep.txt
grep -e qw *.txt > orig_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep S text.txt > orig_grep.txt
./s21_grep S text.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep World test.txt > orig_grep.txt
./s21_grep World test.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep -v T test.txt > orig_grep.txt
./s21_grep -v T test.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep -c W test.txt > orig_grep.txt
./s21_grep -c W test.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep -c t test.txt > orig_grep.txt
./s21_grep -c t test.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep -i T test.txt > orig_grep.txt
./s21_grep -i T test.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep -l S *.txt > orig_grep.txt
./s21_grep -l S *.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep -n W test.txt > orig_grep.txt
./s21_grep -n W test.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep -h S test.txt > orig_grep.txt
./s21_grep -h S test.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep -o S text.txt > orig_grep.txt
./s21_grep -o S text.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt

grep -f text.txt test.txt > orig_grep.txt
./s21_grep -f text.txt test.txt > s21_grep.txt
diff -s orig_grep.txt s21_grep.txt
