#include "mainwindow.h"

#include <QLabel>
#include <QPixmap>
#include <QString>

#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent), ui(new Ui::MainWindow) {
  ui->setupUi(this);
  ui->lineEdit_x_max->setText("10");
  ui->lineEdit_x_min->setText("-10");
  ui->lineEdit_y_max->setText("10");
  ui->lineEdit_y_min->setText("-10");
  ui->label_X->setPlaceholderText("0");
  connect(ui->pushButton_0, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_1, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_2, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_3, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_4, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_5, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_6, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_7, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_8, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_9, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_open, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_close, SIGNAL(clicked()), this, SLOT(s21_numbers()));
  connect(ui->pushButton_X, SIGNAL(clicked()), this, SLOT(s21_numbers()));

  connect(ui->pushButton_plus, SIGNAL(clicked()), this, SLOT(s21_operations()));
  connect(ui->pushButton_minus, SIGNAL(clicked()), this,
          SLOT(s21_operations()));
  connect(ui->pushButton_MUL, SIGNAL(clicked()), this, SLOT(s21_operations()));
  connect(ui->pushButton_DIV, SIGNAL(clicked()), this, SLOT(s21_operations()));
  connect(ui->pushButton_MOD, SIGNAL(clicked()), this, SLOT(s21_operations()));
  connect(ui->pushButton_X2, SIGNAL(clicked()), this, SLOT(s21_operations()));

  connect(ui->pushButton_dot, SIGNAL(clicked()), this, SLOT(s21_dot()));
  connect(ui->pushButton_delete, SIGNAL(clicked()), this,
          SLOT(s21_delete_simbol()));
  connect(ui->pushButton_C, SIGNAL(clicked()), this, SLOT(s21_delete()));

  connect(ui->pushButton_log, SIGNAL(clicked()), this, SLOT(s21_functions()));
  connect(ui->pushButton_ln, SIGNAL(clicked()), this, SLOT(s21_functions()));
  connect(ui->pushButton_sin, SIGNAL(clicked()), this, SLOT(s21_functions()));
  connect(ui->pushButton_cos, SIGNAL(clicked()), this, SLOT(s21_functions()));
  connect(ui->pushButton_tan, SIGNAL(clicked()), this, SLOT(s21_functions()));
  connect(ui->pushButton_asin, SIGNAL(clicked()), this, SLOT(s21_functions()));
  connect(ui->pushButton_acos, SIGNAL(clicked()), this, SLOT(s21_functions()));
  connect(ui->pushButton_atan, SIGNAL(clicked()), this, SLOT(s21_functions()));
  connect(ui->pushButton_sqrt, SIGNAL(clicked()), this, SLOT(s21_functions()));

  connect(ui->pushButton_equals, SIGNAL(clicked()), this, SLOT(s21_result()));

  connect(ui->pushButton_graph, SIGNAL(clicked()), this, SLOT(s21_graphic()));
}

MainWindow::~MainWindow() { delete ui; }

void MainWindow::s21_numbers() {
  numberflag = 1;
  flag = 1;
  if (flag == 0) ui->result_show->setText("");
  QPushButton *button = (QPushButton *)sender();
  ui->result_show->setText(ui->result_show->text() + button->text());
  flag = 1;
}

void MainWindow::s21_operations() {
  numberflag = 0;
  operation = 1;
  if (flag == 0) {
    ui->result_show->setText("");
    operation = 1;
  }
  QPushButton *button = (QPushButton *)sender();

  QString new_label;

  new_label = ui->result_show->text();

  if (button->text() == "÷" && !(ui->result_show->text().endsWith("÷") ||
                                 ui->result_show->text().endsWith("×") ||
                                 ui->result_show->text().endsWith("-") ||
                                 ui->result_show->text().endsWith("^") ||
                                 ui->result_show->text().endsWith("mod") ||
                                 ui->result_show->text().endsWith("+"))) {
    new_label += "/";
  } else if (button->text() == "×" &&
             !(ui->result_show->text().endsWith("÷") ||
               ui->result_show->text().endsWith("×") ||
               ui->result_show->text().endsWith("-") ||
               ui->result_show->text().endsWith("^") ||
               ui->result_show->text().endsWith("mod") ||
               ui->result_show->text().endsWith("+"))) {
    new_label += "*";
  } else if (button->text() == "+" &&
             !(ui->result_show->text().endsWith("÷") ||
               ui->result_show->text().endsWith("×") ||
               ui->result_show->text().endsWith("-") ||
               ui->result_show->text().endsWith("^") ||
               ui->result_show->text().endsWith("mod") ||
               ui->result_show->text().endsWith("+"))) {
    new_label += "+";
  } else if (button->text() == "-" &&
             !(ui->result_show->text().endsWith("÷") ||
               ui->result_show->text().endsWith("×") ||
               ui->result_show->text().endsWith("-") ||
               ui->result_show->text().endsWith("^") ||
               ui->result_show->text().endsWith("mod") ||
               ui->result_show->text().endsWith("+"))) {
    new_label += "-";
  } else if (button->text() == "^" &&
             !(ui->result_show->text().endsWith("÷") ||
               ui->result_show->text().endsWith("×") ||
               ui->result_show->text().endsWith("-") ||
               ui->result_show->text().endsWith("^") ||
               ui->result_show->text().endsWith("mod") ||
               ui->result_show->text().endsWith("+"))) {
    new_label += "^";
  } else if (button->text() == "mod" &&
             !(ui->result_show->text().endsWith("÷") ||
               ui->result_show->text().endsWith("×") ||
               ui->result_show->text().endsWith("-") ||
               ui->result_show->text().endsWith("^") ||
               ui->result_show->text().endsWith("mod") ||
               ui->result_show->text().endsWith("+"))) {
    new_label += "mod";
  } else if (button->text() == "(") {
    new_label += "(";
  } else if (button->text() == ")") {
    new_label += ")";
  }

  ui->result_show->setText(new_label);
  flag = 1;
}

void MainWindow::s21_dot() {
  if (flag == 0) ui->result_show->setText("");
  if (operation == 1) dot = 0;
  if (!(ui->result_show->text().endsWith('.')) && dot == 0) {
    if (numberflag == 0)
      ui->result_show->setText(ui->result_show->text() + "0.");
    else
      ui->result_show->setText(ui->result_show->text() + ".");
    numberflag = 0;
    dot = 1;
  }
  flag = 1;
}

void MainWindow::s21_delete() {
  flag = 1;
  operation = 0;
  dot = 0;
  numberflag = 0;
  ui->result_show->setText("");
  ui->label_X->setText("");
}

void MainWindow::s21_delete_simbol() {
  if (ui->result_show->text().size() == 1) {
    ui->result_show->setText("");
  } else {
    ui->result_show->setText(
        ui->result_show->text().left(ui->result_show->text().size() - 1));
  }
}

void MainWindow::s21_functions() {
  if (flag == 0) {
    ui->result_show->setText("");
    operation = 1;
  }
  QPushButton *button = (QPushButton *)sender();

  if (button->text() == "sin") {
    ui->result_show->setText(ui->result_show->text() + "sin(");
  } else if (button->text() == "asin") {
    ui->result_show->setText(ui->result_show->text() + "asin(");
  } else if (button->text() == "cos") {
    ui->result_show->setText(ui->result_show->text() + "cos(");
  } else if (button->text() == "acos") {
    ui->result_show->setText(ui->result_show->text() + "acos(");
  } else if (button->text() == "tan") {
    ui->result_show->setText(ui->result_show->text() + "tan(");
  } else if (button->text() == "atan") {
    ui->result_show->setText(ui->result_show->text() + "atan(");
  } else if (button->text() == "ln") {
    ui->result_show->setText(ui->result_show->text() + "ln(");
  } else if (button->text() == "log") {
    ui->result_show->setText(ui->result_show->text() + "log(");
  } else if (button->text() == "√") {
    ui->result_show->setText(ui->result_show->text() + "sqrt(");
  }
  flag = 1;
}

void MainWindow::s21_result() {
  QByteArray temp = ui->result_show->text().toLocal8Bit();
  char *src = temp.data();
  double result = 0;
  int error = 0;
  QString x_num = ui->label_X->text();
  double x = x_num.toDouble();
  result = s21_parser(src, &error, x);
  if (error) {
    ui->result_show->setText("error!");
  } else {
    QString res_num = QString::number((double)result);
    ui->result_show->setText(res_num);
  }
}

void MainWindow::s21_graphic() {
  ui->widget->clearGraphs();

  QString str = ui->result_show->text();
  QByteArray ba = str.toLocal8Bit();
  char *buf = ba.data();

  h = 0.1;
  xBegin = ui->lineEdit_x_min->text().toDouble();
  xEnd = ui->lineEdit_x_max->text().toDouble();
  double y_min = ui->lineEdit_y_min->text().toDouble();
  double y_max = ui->lineEdit_y_max->text().toDouble();
  int error = 0;
  ui->widget->clearGraphs();
  ui->widget->xAxis->setRange(xBegin, xEnd);
  ui->widget->yAxis->setRange(y_min, y_max);
  if (str != "") {
    for (X = xBegin; X <= xEnd; X += h) {
      Y = s21_parser(buf, &error, X);
      if (Y < (y_min * 2) || Y > (y_max * 2)) {
        Y = qQNaN();
      }
      x.push_back(X);
      y.push_back(Y);
    }
  }
  ui->widget->addGraph();
  ui->widget->graph(0)->addData(x, y);
  ui->widget->replot();
  x.clear();
  y.clear();
}

void MainWindow::on_pushButton_credit_clicked() { form.show(); }

void MainWindow::on_pushButton_10_clicked() { form_2.show(); }
