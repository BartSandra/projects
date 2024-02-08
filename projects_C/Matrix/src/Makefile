TESTFLAGS = -lcheck
GCCFLAGS = -Wall -Werror -Wextra
GCOVFLAGS = -fprofile-arcs -ftest-coverage

all: clean clean_screen test gcov_report

s21_matrix.a:
	@gcc $(GCCFLAGS) -std=c11 -c s21_matrix.c s21_fun.c
	ar rc s21_matrix.a *.o

test: s21_matrix.a
	@gcc $(GCCFLAGS) $(GCOVFLAGS) test.c $(TESTFLAGS) s21_fun.c s21_matrix.c -o test
	@./test
	
gcov_report: test
	@lcov -t "./gcov" -o report.info --no-external -c -d .
	@genhtml -o report report.info
	@open ./report/index.html

clean:
	@rm -rf runme
	@rm -rf *.o *.a *.out *.gcno *.gcda *.info *.gch test report gcov_report
	@clear

clean_screen:
	@clear

rebuild: clean all

style:
	clang-format --style=google -n *.c *.h
	rm -rf .clang-format
