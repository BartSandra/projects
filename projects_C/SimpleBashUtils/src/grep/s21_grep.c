#include "s21_grep.h"

int main(int argc, char **argv) {
  char s21_pattern[bufferSize] = {0};
  if (argc == optind) {
    fprintf(stderr,
            "usage: grep [-abcDEregerrHhIiJLlmnOoqRSsUVvwxZ] [-A num] [-B num] "
            "[-C[num]]\n\t[-e pattern] [-f file] [--binary-files=value] "
            "[--color=when]\n\t[--context[=num]] [--directories=action] "
            "[--label] [--line-buffered]\n\t[--null] [pattern] [file ...]\n");
  } else if (argc < 3) {
    no_argc(argc);
  } else if (argc >= 3) {
    parser_options(argc, argv, &options, s21_pattern);
    if (!options.e && !options.f) {
      if (!*argv[optind]) argv[optind] = ".";
      strcat(s21_pattern, argv[optind]);
      optind += 1;
    }
    reader(argc, argv, &options, s21_pattern);
  }
  return 0;
}

void no_argc(int fd) {  // grep text.txt

  char buff[bufferSize];
  int byte_read;
  byte_read =
      read(fd, buff,
           bufferSize);  // read() считывает bufferSize байт из файла fd в buff
  while (byte_read) {
    byte_read = read(fd, buff, bufferSize);
  }
}

void parser_options(int argc, char **argv, opt *options,
                    char s21_pattern[bufferSize]) {
  int opt = 0;
  while ((opt = getopt(argc, argv, ":eivclnhsfo")) != -1) {
    switch (opt) {
      case 'e':
        options->e = 1;
        if (!*argv[optind]) argv[optind] = ".";
        strcat(s21_pattern, argv[optind]);
        optind += 1;
        break;
      case 'i':
        options->i = 1;
        break;
      case 'v':
        options->v = 1;
        break;
      case 'c':
        options->c = 1;
        break;
      case 'l':
        options->l = 1;
        break;
      case 'n':
        options->n = 1;
        break;
      case 'h':
        options->h = 1;
        break;
      case 's':
        options->s = 1;
        break;
      case 'f':
        options->f = 1;
        opt_option_f(s21_pattern, argv[optind]);
        break;
      case 'o':
        options->o = 1;
        break;
      default:
        fprintf(stderr, "grep: unknown --directories option\n");
        break;
    }
  }
  if (options->empty_lines)  // если пустая строка
    options->o = 0;
}

void reader(int argc, char **argv, opt *options, char *s21_pattern) {
  FILE *fp = NULL;
  char *buffer = NULL;  // указатель на строку
  int count_n = 1;
  int current_file =
      optind;  // индекс следующего обрабатываемого элемента в argv
  int count_c = 0;
  int count_cv = 0;
  while (current_file < argc) {
    buffer = calloc(bufferSize, sizeof(char));
    fp = fopen(argv[current_file], "r");
    if (fp == NULL) {
      if (!options->s) {
        fprintf(stderr, "%s: %s: No such file or directory", argv[0],
                argv[current_file]);
      }
      exit(1);
    }
    while (
        fgets(buffer, bufferSize, fp)) {  // считываем из fp, bufferSize
                                          // символов, сохраняем в строку buffer
      regex_h(argc, argv, current_file, buffer, options, count_n, &count_c,
              &count_cv, s21_pattern);
      count_n++;
    }
    if (argc - optind > 1 && ((options->c && !options->v) ||
                              (options->c && options->v))) {  // -c || -cv
      printf("%s:", argv[current_file]);
    }
    if (options->c && !options->v) {  // -c
      printf("%d\n", count_c);
    }
    if (options->c && options->v) {  // -cv
      printf("%d\n", count_cv);
    }
    if (options->l && !options->v && count_c > 0) {  // -l
      printf("%s\n", argv[current_file]);
    } else {
      if (options->l && options->v && count_cv > 0) {  // -lv
        printf("%s\n", argv[current_file]);
      }
    }
    fclose(fp);
    free(buffer);
    current_file++;
    count_n = 1;
    count_c = 0;
    count_cv = 0;
  }
}

void regex_h(int argc, char **argv, int current_file, char *buffer,
             opt *options, int count_n, int *count_c, int *count_cv,
             char *s21_pattern) {  // buffer - строка
  regex_t regex;                   //структура
  int status;

  regmatch_t rm[1];

  int matc;
  int regex_options = REG_EXTENDED;  // рег. выражение
  if (options->i) {                  // -i
    regex_options =
        REG_EXTENDED | REG_ICASE;  // не учитывать регистр  | побитовое или
  }
  status = regcomp(&regex, s21_pattern,
                   regex_options);  // компиляция регулярного выражения,
                                    // возвращает ноль при успешной компиляции
  if (status == 0) {
    matc = regexec(&regex, buffer, 1, rm,
                   0);  // сравнение строк, завершающейся нулем, с
                        // предворительно обработанным буферным шаблоном regex
                        // // возвращает ноль при совпадениях
  }
  if (matc == 0) {
    *count_c += 1;
    if (argc - optind > 1 && !options->f && !options->l && !options->h &&
        !options->v && !options->c)
      printf("%s:", argv[current_file]);
    if (options->n && !options->f && !options->l && !options->c &&
        !options->v)  // -n
      printf("%d:", count_n);

    if (!options->v && !options->o && !options->f && !options->l &&
        !options->c) {
      if (buffer[strlen(buffer) - 1] != '\n') {
        printf("%s\n", buffer);
      } else {
        printf("%s", buffer);
      }
    }

    if (options->f && !options->o) {
      if (buffer[strlen(buffer) - 1] != '\n') {
        printf("%s", buffer + 1);
      } else {
        printf("%s", buffer);
      }
    }

    if (options->o && (options->v == 0)) {
      do {
        printf("%.*s\n", (int)(rm[0].rm_eo - rm[0].rm_so),
               buffer + rm[0].rm_so);
        buffer = buffer + (int)rm[0].rm_eo;
      } while (regexec(&regex, buffer, 1, rm, 0) == 0);
    }

  } else {
    if (argc - optind > 1 && options->v && !options->c && !options->l &&
        !options->h && !options->e) {  // -v
      printf("%s:", argv[current_file]);
    }
    if (options->n && options->v && !options->l && !options->c) {  // -nv
      printf("%d:", count_n);
    }
    if (options->v && !options->c && !options->l) {  // -v
      if (buffer[strlen(buffer) - 1] != '\n') {
        printf("%s\n", buffer);
      } else {
        printf("%s", buffer);
      }
    }
    *count_cv += 1;
  }
  regfree(&regex);  // regex освободит память, отведенную шаблону во время
                    // процесса компиляции regcomp
}

int opt_option_f(char *s21_pattern, char *file) {
  FILE *fp;
  fp = fopen(file, "r");
  int i = 0;

  if (fp == NULL) {
    i = -1;
  } else {
    int ch;
    while ((ch = fgetc(fp)) != EOF) {
      if (ch == 13 || ch == 10) s21_pattern[i++] = '|';
      if (ch != 13 && ch != 10) s21_pattern[i++] = ch;
    }

    if (s21_pattern[i - 1] == '|') s21_pattern[i - 1] = '\0';
    fclose(fp);
  }
  return (i == -1) ? -1 : (i + 1);
}
