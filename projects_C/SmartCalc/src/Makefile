GCCFLAGS=-Wall -Wextra -Werror -g -std=c11
TESTFLAGS=-lcheck
GCOVFLAGS=-fprofile-arcs -ftest-coverage
LINUX = -lrt -lpthread -lm -lsubunit
OS=$(shell uname -s)

all: clean clean_screen gcov_report dvi

test:
ifeq ($(OS), Darwin)
	@gcc $(GCCFLAGS) test.c s21_SmartCalc.c -o test $(TESTFLAGS)
	@./test
else
	@gcc $(GCCFLAGS) test.c s21_SmartCalc.c -o test $(TESTFLAGS) $(LINUX)
	@./test
endif

gcov_report: s21_SmartCalc.c
ifeq ($(OS), Darwin)
	@gcc $(GCCFLAGS) $(GCOVFLAGS) test.c s21_SmartCalc.c -o test $(TESTFLAGS)
else
	@gcc $(GCCFLAGS) $(GCOVFLAGS) test.c s21_SmartCalc.c -o test $(TESTFLAGS) $(LINUX)
endif
	@./test
	@lcov -t "./gcov" -o report.info --no-external -c -d .
	@genhtml -o report report.info
	@make open

open:
ifeq ($(OS), Darwin)
	@open report/index.html
else
	@xdg-open report/index.html
endif

rebuild: clean all

clean:
	@rm -rf runme
	@rm -rf *.o *.a *.out *.gcno *.gcda *.info *.gch test test.dSYM report gcov_report build
	@clear
	
clean_screen:
	@clear

build:
	mkdir build
	cd ./build/ && qmake ../ && make

install: build
	@cp -rf build/calc.app $(HOME)/Desktop/calc.app
	#make clean

uninstall:
	@rm -rf $(HOME)/Desktop/calc.app

dvi:
	open description.md

dist:
	mkdir SmartCalc_v1.0/
	mkdir SmartCalc_v1.0/src
	cp Makefile *.c *.h *.pro *.cpp *.ui *.user SmartCalc_v1.0/src/
	tar cvzf SmartCalc_v1.0.tgz SmartCalc_v1.0/
	mv SmartCalc_v1.0.tgz $(HOME)/Desktop/
	rm -rf SmartCalc_v1.0/

open_app:
	@open ./build/calc.app
	
clang:
	clang-format --style=Google -n *.c *.h
	clang-format --style=Google -n *.cpp *.h
