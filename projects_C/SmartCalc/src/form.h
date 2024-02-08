#ifndef FORM_H
#define FORM_H

#include <QWidget>

extern "C" {
#include "s21_SmartCalc.h"
}

namespace Ui {
class Form;
}

class Form : public QWidget {
  Q_OBJECT

 public:
  explicit Form(QWidget *parent = nullptr);
  ~Form();

 private slots:
  void on_pushButton_clicked();
  void on_radioButton_clicked();
  void on_radioButton_2_clicked();

 private:
  Ui::Form *ui;

 public:
  QString globalStr;
  int globalInt = 1;
};

#endif  // FORM_H
