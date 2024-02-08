TESTFLAGS=-lgtest -lgcov
GCOVFLAGS=--coverage
GCCFLAGS = -Wall -Werror -Wextra
HTML=lcov -t test -o rep.info -c -d ./
CFLAGS=--std=c++17 -lstdc++ -lm

all: clean clean_screen s21_matrix_oop.a test gcov_report

s21_matrix_oop.a: clean
	@gcc $(GCOVFLAGS) -c s21_matrix_oop.cc
	@ar rcs s21_matrix_oop.a *.o
	@ranlib s21_matrix_oop.a

test: s21_matrix_oop.a
	gcc -g test.cc s21_matrix_oop.a $(CFLAGS) $(TESTFLAGS) -o test
	./test
	

gcov_report: test
	@lcov -t "./gcov" -o report.info --no-external -c -d .
	@genhtml -o report report.info
	@open ./report/index.html

clean:
	@rm -rf runme
	@rm -rf *.o *.a *.out *.gcno *.gcda *.info *.gch test report *.dSYM
	@clear

clean_screen:
	@clear

rebuild: clean all

style:
	@cp ../materials/linters/.clang-format ../src/.clang-format
	@clang-format --style=google -n *.cc *.h
	@rm -rf .clang-format
