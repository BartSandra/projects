CC = gcc
FLAGS = -Wall -Werror -Wextra -std=c11

all: s21_cat

cat.o: s21_cat.c s21_cat.h
	$(CC) $(FLAGS) -c s21_cat.c -o s21_cat.o

s21_cat: cat.o 
	$(CC) $(FLAGS) *.o  -o s21_cat

rebuild:
	make clean
	make s21_cat

clean:
	rm -rf *.o
	rm -rf s21_cat 
