#include "s21_matrix_oop.h"

bool S21Matrix::EqMatrix(const S21Matrix &other) const {
  bool flag = true;
  if (this->cols_ != other.cols_ || this->rows_ != other.rows_) {
    flag = false;
  } else {
    for (int i = 0; i < other.rows_; ++i) {
      for (int j = 0; j < other.cols_; ++j) {
        if (fabs(other.matrix_[i][j] - this->matrix_[i][j]) > EPS) {
          flag = false;
        }
      }
    }
  }
  return flag;
}

void S21Matrix::SumMatrix(const S21Matrix &other) {
  if (this->cols_ != other.cols_ || this->rows_ != other.rows_) {
    throw out_of_range("Different matrix dimensions");
  }
  for (int i = 0; i < other.rows_; ++i) {
    for (int j = 0; j < other.cols_; ++j) {
      this->matrix_[i][j] += other.matrix_[i][j];
    }
  }
}

void S21Matrix::SubMatrix(const S21Matrix &other) {
  if (this->cols_ != other.cols_ || this->rows_ != other.rows_) {
    throw out_of_range("Different matrix dimensions");
  }
  for (int i = 0; i < other.rows_; ++i) {
    for (int j = 0; j < other.cols_; ++j) {
      this->matrix_[i][j] -= other.matrix_[i][j];
    }
  }
}

void S21Matrix::MulNumber(const double num) {
  for (int i = 0; i < this->rows_; ++i) {
    for (int j = 0; j < this->cols_; ++j) {
      this->matrix_[i][j] *= num;
    }
  }
}

void S21Matrix::MulMatrix(const S21Matrix &other) {
  if (this->cols_ != other.rows_) {
    throw out_of_range(
        "The number of columns of the first matrix is not equal to the number "
        "of rows of the second matrix");
  }
  S21Matrix result(this->rows_, other.cols_);
  for (int i = 0; i < this->rows_; ++i) {
    for (int j = 0; j < other.cols_; ++j) {
      for (int k = 0; k < this->cols_; ++k) {
        result.matrix_[i][j] += this->matrix_[i][k] * other.matrix_[k][j];
      }
    }
  }
  *this = result;
}

S21Matrix S21Matrix::Transpose() {
  S21Matrix result(this->cols_, this->rows_);
  for (int i = 0; i < result.rows_; ++i) {
    for (int j = 0; j < result.cols_; ++j) {
      result.matrix_[i][j] = this->matrix_[j][i];
    }
  }
  return result;
}

void S21Matrix::Minor(int i_row, int j_column, const S21Matrix &minor) {
  int minor_i = 0;
  int minor_j = 0;
  for (int current_i = 0; current_i < this->rows_; ++current_i) {
    for (int current_j = 0; current_j < this->cols_; ++current_j) {
      if (current_j != j_column && current_i != i_row) {
        minor.matrix_[minor_i][minor_j] = this->matrix_[current_i][current_j];
        ++minor_j;
        if (minor_j == this->cols_ - 1) {
          ++minor_i;
          minor_j = 0;
        }
      }
    }
  }
}

S21Matrix S21Matrix::CalcComplements() {
  if (this->rows_ != this->cols_) {
    throw out_of_range("The matrix is not square");
  }
  S21Matrix result(this->rows_, this->cols_);
  S21Matrix minor(this->rows_ - 1, this->cols_ - 1);
  for (int i = 0; i < this->rows_; ++i) {
    for (int j = 0; j < this->cols_; ++j) {
      Minor(i, j, minor);
      result.matrix_[i][j] = pow(-1, (j + i) + 2) * minor.Determinant();
    }
  }
  return result;
}

double S21Matrix::Determinant() {
  double result = 0.0;
  if (this->rows_ != this->cols_) {
    throw out_of_range("The matrix is not square");
  }
  if (this->rows_ == 1) {
    result = this->matrix_[0][0];
  } else if (this->rows_ == 2) {
    result = this->matrix_[0][0] * this->matrix_[1][1] -
             this->matrix_[1][0] * this->matrix_[0][1];
  } else {
    int sign = pow(-1, 1 + 1);
    double determ_low;
    for (int i = 0; i < this->rows_; ++i) {
      S21Matrix matrix_low(rows_ - 1, cols_ - 1);
      Minor(i, 0, matrix_low);
      determ_low = matrix_low.Determinant();
      result += this->matrix_[i][0] * sign * determ_low;
      sign = -sign;
    }
  }
  return result;
}

S21Matrix S21Matrix::InverseMatrix() {
  double det = Determinant();
  if (det == 0) {
    throw out_of_range("Matrix determinant is 0");
  }
  if (this->rows_ != this->cols_) {
    throw out_of_range("Matrix is not square");
  }
  S21Matrix result(this->rows_, this->cols_);
  if (this->rows_ == 1) {
    result.matrix_[0][0] = 1 / det;
  } else {
    S21Matrix temp = this->CalcComplements();
    S21Matrix temp_2 = temp.Transpose();
    result = temp_2;
    result.MulNumber(1 / det);
  }
  return result;
}

S21Matrix::S21Matrix() : rows_(1), cols_(1), matrix_(nullptr) {
  this->Memory_allocation();
}

S21Matrix::~S21Matrix() { this->Delete_matrix(); }

S21Matrix::S21Matrix(int rows, int cols) : rows_(rows), cols_(cols) {
  this->Memory_allocation();
}

S21Matrix::S21Matrix(const S21Matrix &other)
    : S21Matrix(other.rows_, other.cols_) {
  for (int i = 0; i < rows_; ++i) {
    for (int j = 0; j < cols_; ++j) {
      this->matrix_[i][j] = other.matrix_[i][j];
    }
  }
}

S21Matrix::S21Matrix(S21Matrix &&other) {
  this->matrix_ = other.matrix_;
  this->rows_ = other.rows_;
  this->cols_ = other.cols_;
  other.matrix_ = nullptr;
  other.cols_ = 0;
  other.rows_ = 0;
}

void S21Matrix::Delete_matrix() {
  if (this->matrix_ != nullptr) {
    for (int i = 0; i < this->rows_; ++i) {
      delete[] matrix_[i];
      this->matrix_[i] = nullptr;
    }
    delete[] this->matrix_;
    this->matrix_ = nullptr;
  }
  this->cols_ = 0;
  this->rows_ = 0;
}

void S21Matrix::Memory_allocation() {
  if (this->rows_ < 1 || this->cols_ < 1) {
    throw out_of_range("Memory allocation error!");
  }
  this->matrix_ = new double *[this->rows_]();
  for (int i = 0; i < this->rows_; ++i) {
    this->matrix_[i] = new double[this->cols_]();
  }
}

void S21Matrix::Copy_matrix(const S21Matrix &other) {
  for (int i = 0; i < this->rows_; ++i) {
    for (int j = 0; j < this->cols_; ++j) {
      this->matrix_[i][j] = other.matrix_[i][j];
    }
  }
}

int S21Matrix::getRows() { return this->rows_; }

void S21Matrix::setRows(int rows) {
  S21Matrix copy(*this);
  this->Delete_matrix();
  rows_ = rows;
  cols_ = copy.cols_;
  this->Memory_allocation();
  for (int i = 0; i < this->rows_; ++i) {
    for (int j = 0; j < this->cols_; ++j) {
      this->matrix_[i][j] = copy.matrix_[i][j];
    }
  }
}

int S21Matrix::getCols() { return this->cols_; }

void S21Matrix::setCols(int cols) {
  S21Matrix copy(*this);
  this->Delete_matrix();
  rows_ = copy.rows_;
  cols_ = cols;
  this->Memory_allocation();
  for (int i = 0; i < this->rows_; ++i) {
    for (int j = 0; j < this->cols_; ++j) {
      this->matrix_[i][j] = copy.matrix_[i][j];
    }
  }
}

S21Matrix S21Matrix::operator+(const S21Matrix &other) {
  S21Matrix result(*this);
  result.SumMatrix(other);
  return result;
}

S21Matrix S21Matrix::operator-(const S21Matrix &other) {
  S21Matrix result(*this);
  result.SubMatrix(other);
  return result;
}

S21Matrix S21Matrix::operator*(const S21Matrix &other) {
  S21Matrix result(*this);
  result.MulMatrix(other);
  return result;
}

S21Matrix S21Matrix::operator*(const double num) {
  S21Matrix result(*this);
  result.MulNumber(num);
  return result;
}

bool S21Matrix::operator==(const S21Matrix &other) const {
  return this->EqMatrix(other);
}

S21Matrix &S21Matrix::operator=(const S21Matrix &other) {
  this->Delete_matrix();
  this->rows_ = other.rows_;
  this->cols_ = other.cols_;
  this->Memory_allocation();
  this->Copy_matrix(other);
  return *this;
}

S21Matrix S21Matrix::operator+=(const S21Matrix &other) {
  this->SumMatrix(other);
  return *this;
}

S21Matrix S21Matrix::operator-=(const S21Matrix &other) {
  this->SubMatrix(other);
  return *this;
}

S21Matrix S21Matrix::operator*=(const S21Matrix &other) {
  this->MulMatrix(other);
  return *this;
}

S21Matrix S21Matrix::operator*=(const double num) {
  this->MulNumber(num);
  return *this;
}

double &S21Matrix::operator()(int i, int j) {
  if (i < 0 || i > rows_ - 1) {
    throw out_of_range("Invalid argument i - rows");
  }
  if (j < 0 || j > cols_ - 1) {
    throw out_of_range("Invalid argument j - cols");
  }
  return matrix_[i][j];
}
