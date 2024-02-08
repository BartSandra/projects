#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <string.h>

#include <QMainWindow>
#include <QVector>

#include "form.h"
#include "form_2.h"

extern "C" {
#include "s21_SmartCalc.h"
}

QT_BEGIN_NAMESPACE
namespace Ui {
class MainWindow;
}
QT_END_NAMESPACE

class MainWindow : public QMainWindow {
  Q_OBJECT

 public:
  MainWindow(QWidget *parent = nullptr);
  ~MainWindow();

 private slots:
  void s21_numbers();
  void s21_operations();
  void s21_result();
  void s21_delete();
  void s21_dot();
  void s21_functions();
  void s21_delete_simbol();
  void s21_graphic();

  void on_pushButton_credit_clicked();

  void on_pushButton_10_clicked();

 private:
  Ui::MainWindow *ui;
  //  graf *graph;
  int flag = 1;
  int operation = 0;
  int dot = 0;
  int numberflag = 0;
  double xBegin, xEnd, h;
  double X = 0.0;
  double Y = 0.0;
  int N;
  QVector<double> x, y;
  Form form;
  Form_2 form_2;
};
#endif  // MAINWINDOW_H
