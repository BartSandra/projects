#include "s21_cat.h"

int main(int argc, char **argv) {
  if (argc == 1) {
    const int bufferSize = 4096;
    char buffer[bufferSize];

    while (fgets(buffer, bufferSize, stdin)) {
      int length = strlen(buffer);
      buffer[length - 1] = '\0';
      fprintf(stdout, "%s\n", buffer);
    }
  } else if (parser_options(argc, argv, &options) == 0) {
    while (optind < argc) {
      reader(argv, &options);
      optind++;
    }
  } else {
    fprintf(stderr,
            "cat: illegal option -- T\n"
            "usage: cat [-benstuv] [file ...]\n");
  }
  return 0;
}

int parser_options(int argc, char **argv,
                   opt *options) {  //указатель на структуру

  static struct option long_options[] = {
      {"number-nonblank", 0, 0, 'b'},
      {"number", 0, 0, 'n'},
      {"squeeze-blank", 0, 0, 's'},
      {NULL, 0, NULL, 0}  // где заканчивается массив
  };

  int error = 0;
  int opt;               // ссылка на результат функции
  int option_index = 0;  // для длинных аргументов
  while ((opt = getopt_long(argc, argv, "+benstvTE", long_options,
                            &option_index)) != -1) {
    switch (opt) {
      case 'b':
        options->b = 1;
        break;

      case 'e':
        options->e = 1;
        options->v = 1;
        break;

      case 'E':
        options->e = 1;
        break;

      case 'n':
        options->n = 1;
        break;

      case 's':
        options->s = 1;
        break;

      case 't':
        options->t = 1;
        options->v = 1;
        break;

      case 'T':
        options->t = 1;
        break;

      case 'v':
        options->v = 1;
        break;

      default:
        fprintf(stderr,
                "cat: illegal option -- T\n"
                "usege: cat [-benstvTE] [file ...]\n");
        error = 1;
        break;
    }
  }
  return error;
}

int reader(char **argv, opt *options) {
  int error = 0;

  FILE *fp = fopen(argv[optind], "r");

  if (fp == NULL) {
    error = 2;
    fprintf(stderr, "s21_cat: %s: No such file or directory\n", *argv);
  } else {
    int count_line = 1;   //счетчик строк
    int empty_lines = 0;  // пустые строки
    char current = '\0';  //текущий символ
    char last = '\n';
    int counter_for_b = 0;

    while ((current = fgetc(fp)) != EOF) {
      if (options->s == 1 && current == '\n') {
        if (empty_lines >= 1) {
          continue;
        }
        empty_lines++;
      } else {
        empty_lines = -1;
      }
      if ((options->b != 1 && options->n == 1 && last == '\n') ||
          (options->b == 1 && current != '\n' && last == '\n')) {
        if (counter_for_b == 1) {
        } else {
          fprintf(stdout, "%6d\t", count_line++);
          counter_for_b++;
        }
      }

      if (options->v) {
        options_v(options, &current);
      }

      if (options->e == 1 && current == '\n') {
        fprintf(stdout, "%c", '$');
      }

      if (options->t == 1 && current == '\t') {
        fprintf(stdout, "^I");
        continue;
      }

      fprintf(stdout, "%c", current);
      last = current;
      counter_for_b = 0;
    }
  }
  fclose(fp);
  return error;
}

void options_v(opt *options, char *c) {
  if (options->v) {
    int ch = (int)*c;
    options_v_ch(&ch, c);
    if (ch != 9 && ch != 10 && ch < 32) {
      printf("^");
      *c += 64;
    } else if (ch == 127) {
      printf("^");
      *c = '?';
    } else if (ch > 127 && ch < 160) {
      printf("M-^");
      *c = ch - 64;
    } else if (ch > 160 && ch <= 255) {
      *c -= 128;
    }
  }
}

void options_v_ch(int *ch, char *c) {
  if (*c < 0) {
    *c &= 127;
    *ch = (int)*c;
    *ch += 128;
  }
}
