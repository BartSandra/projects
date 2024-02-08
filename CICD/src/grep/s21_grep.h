#ifndef S21_GREP_H
#define S21_GREP_H

#include <getopt.h>
#include <regex.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>  //optarg, optind, read()

#define bufferSize 4096

typedef struct options {
  int e;
  int i;
  int v;
  int c;
  int l;
  int n;
  int h;
  int s;
  int f;
  int o;
  int empty_lines;
} opt;

opt options = {0};

void no_argc(int fd);
void parser_options(int argc, char **argv, opt *options,
                    char s21_pattern[bufferSize]);
int opt_option_f(char *s21_pattern, char *file);
void reader(int argc, char **argv, opt *options, char *s21_pattern);
void regex_h(int argc, char **argv, int current_file, char *buff, opt *options,
             int count_n, int *count_c, int *count_vc, char *s21_pattern);

#endif  // S21_GREP_H
