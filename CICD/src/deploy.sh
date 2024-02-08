#!/bin/bash

scp /home/gitlab-runner/builds/mHhksKns/0/students/DO6_CICD.ID_356283/finchmar_student.21_school.ru/DO6_CICD-2/src/cat/s21_cat finchmarws2@192.168.100.10:/usr/local/bin/
scp /home/gitlab-runner/builds/mHhksKns/0/students/DO6_CICD.ID_356283/finchmar_student.21_school.ru/DO6_CICD-2/src/grep/s21_grep finchmarws2@192.168.100.10:/usr/local/bin/
ssh finchmarws2@192.168.100.10 ls -lah /usr/local/bin