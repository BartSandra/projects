#ifndef S21_CAT_H
#define S21_CAT_H

#include <getopt.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

typedef struct options {
  int b;
  int e;
  int n;
  int s;
  int t;
  int v;
} opt;

opt options = {0};

int parser_options(int argc, char **argv, opt *options);
int reader(char **argv, opt *options);
void options_v(opt *options, char *c);
void options_v_ch(int *ch, char *c);

#endif  // S21_CAT_H
