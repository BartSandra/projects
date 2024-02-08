CC = gcc
FLAGS = -Wall -Werror -Wextra -std=c11

all: s21_grep

s21_grep.o: s21_grep.c s21_grep.h
	$(CC) $(FLAGS) -c s21_grep.c -o s21_grep.o

s21_grep: s21_grep.o 
	$(CC) $(FLAGS) *.o  -o s21_grep

rebuild:
	make clean
	make s21_grep

clean:
	rm -rf *.o
	rm -rf s21_grep 
