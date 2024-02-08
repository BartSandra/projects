#include "gtest/gtest.h"
#include "s21_matrix_oop.h"

TEST(Matrix, EqMatrix1) {
  S21Matrix a(2, 2);
  S21Matrix b(2, 2);
  a(0, 0) = 5.3;
  a(0, 1) = 5;
  a(1, 0) = 99.9;
  a(1, 1) = -0.5;
  b(0, 0) = 5.3;
  b(0, 1) = 5;
  b(1, 0) = 99.9;
  b(1, 1) = -0.5;
  EXPECT_TRUE(a.EqMatrix(b));
}

TEST(Matrix, EqMatrix2) {
  S21Matrix a(2, 2);
  S21Matrix b(2, 2);
  a(0, 0) = 5.3;
  a(0, 1) = 5;
  a(1, 0) = 11.3;
  a(1, 1) = -0.5;
  b(0, 0) = 5.3;
  b(0, 1) = 5;
  b(1, 0) = 10;
  b(1, 1) = -0.5;
  EXPECT_FALSE(a.EqMatrix(b));
}

TEST(Matrix, EqMatrix3) {
  S21Matrix a(2, 3);
  S21Matrix b(2, 2);
  a(0, 0) = 5.3;
  a(0, 1) = 5;
  a(0, 2) = 7;
  a(1, 0) = 99.9;
  a(1, 1) = -0.5;
  a(1, 2) = -1;
  b(0, 0) = 5.3;
  b(0, 1) = 5;
  b(1, 0) = 99.9;
  b(1, 1) = -0.5;
  EXPECT_FALSE(a.EqMatrix(b));
}

TEST(Matrix, SumMatrix1) {
  S21Matrix a(2, 3);
  S21Matrix b(2, 2);
  a(0, 0) = 4.3;
  a(0, 1) = 3;
  a(0, 2) = 7;
  a(1, 0) = 11.3;
  a(1, 1) = -1.3;
  a(1, 2) = -1;
  b(0, 0) = 4.3;
  b(0, 1) = 3;
  b(1, 0) = 11.3;
  b(1, 1) = -1.3;
  EXPECT_THROW(a.SumMatrix(b), out_of_range);
}

TEST(Matrix, SumMatrix2) {
  S21Matrix a(2, 2);
  S21Matrix b(2, 2);
  a(0, 0) = 4.3;
  a(0, 1) = 3;
  a(1, 0) = 1;
  a(1, 1) = -1;
  b(0, 0) = 9;
  b(0, 1) = 0;
  b(1, 0) = -5;
  b(1, 1) = 6;
  S21Matrix res(2, 2);
  res(0, 0) = 13;
  res(0, 1) = 3;
  res(1, 0) = -4;
  res(1, 1) = 5;
  a.SumMatrix(b);
  EXPECT_FALSE(a.EqMatrix(res));
}

TEST(Matrix, SumMatrix3) {
  S21Matrix a(2, 2);
  S21Matrix b(2, 2);
  a(0, 0) = 4.3;
  a(0, 1) = 3;
  a(1, 0) = 0;
  a(1, 1) = -1;
  b(0, 0) = 9;
  b(0, 1) = 0;
  b(1, 0) = -5;
  b(1, 1) = 6;
  S21Matrix res(2, 2);
  res(0, 0) = 13;
  res(0, 1) = 3;
  res(1, 0) = -4;
  res(1, 1) = 5;
  a.SumMatrix(b);
  EXPECT_FALSE(a.EqMatrix(res));
}

TEST(Matrix, SumMatrix4) {
  S21Matrix a(2, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  a(1, 0) = 1;
  a(1, 1) = -1;
  S21Matrix b(2, 2);
  b(0, 0) = 9;
  b(0, 1) = 0;
  b(1, 0) = -5;
  b(1, 1) = 6;

  S21Matrix res(2, 2);
  res(0, 0) = 13;
  res(0, 1) = 3;
  res(1, 0) = -4;
  res(1, 1) = 5;

  EXPECT_TRUE((a + b) == res);
}

TEST(Matrix, SumMatrix5) {
  S21Matrix a(2, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  a(1, 0) = 1;
  a(1, 1) = -1;
  S21Matrix b(2, 2);
  b(0, 0) = 9;
  b(0, 1) = 0;
  b(1, 0) = -5;
  b(1, 1) = 6;
  S21Matrix res(2, 2);
  res(0, 0) = 13;
  res(0, 1) = 3;
  res(1, 0) = -4;
  res(1, 1) = 5;
  EXPECT_TRUE((a += b) == res);
}

TEST(Matrix, SumMatrix6) {
  S21Matrix a(1, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  S21Matrix b(2, 2);
  b(0, 0) = 9;
  b(0, 1) = 0;
  b(1, 0) = -5;
  b(1, 1) = 6;
  EXPECT_THROW(a += b, out_of_range);
}

TEST(Matrix, SubMatrix1) {
  S21Matrix a(2, 3);
  S21Matrix b(2, 2);
  a(0, 0) = 4.3;
  a(0, 1) = 3;
  a(0, 2) = 7;
  a(1, 0) = 11.3;
  a(1, 1) = -1.3;
  a(1, 2) = -1;
  b(0, 0) = 4.3;
  b(0, 1) = 3;
  b(1, 0) = 11.3;
  b(1, 1) = -1.3;
  EXPECT_THROW(a.SumMatrix(b), out_of_range);
}

TEST(Matrix, SubMatrix2) {
  S21Matrix a(2, 2);
  S21Matrix b(2, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  a(1, 0) = 1;
  a(1, 1) = -1;
  b(0, 0) = 9;
  b(0, 1) = 0;
  b(1, 0) = -5;
  b(1, 1) = 6;
  S21Matrix res(2, 2);
  res(0, 0) = -5;
  res(0, 1) = 3;
  res(1, 0) = 6;
  res(1, 1) = -7;
  a.SubMatrix(b);
  EXPECT_TRUE(a.EqMatrix(res));
}

TEST(Matrix, SubMatrix3) {
  S21Matrix a(2, 2);
  S21Matrix b(2, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  a(1, 0) = 0;
  a(1, 1) = -1;
  b(0, 0) = 9;
  b(0, 1) = 0;
  b(1, 0) = -5;
  b(1, 1) = 6;
  S21Matrix res(2, 2);
  res(0, 0) = -5;
  res(0, 1) = 3;
  res(1, 0) = 6;
  res(1, 1) = -7;
  a.SubMatrix(b);
  EXPECT_FALSE(a.EqMatrix(res));
}

TEST(Matrix, SubMatrix4) {
  S21Matrix a(2, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  a(1, 0) = 1;
  a(1, 1) = -1;
  S21Matrix b(2, 2);
  b(0, 0) = 9;
  b(0, 1) = 0;
  b(1, 0) = -5;
  b(1, 1) = 6;

  S21Matrix res(2, 2);
  res(0, 0) = -5;
  res(0, 1) = 3;
  res(1, 0) = 6;
  res(1, 1) = -7;

  EXPECT_TRUE((a - b) == res);
}

TEST(Matrix, SubMatrix5) {
  S21Matrix a(2, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  a(1, 0) = 1;
  a(1, 1) = -1;
  S21Matrix b(2, 2);
  b(0, 0) = 9;
  b(0, 1) = 0;
  b(1, 0) = -5;
  b(1, 1) = 6;
  S21Matrix res(2, 2);
  res(0, 0) = -5;
  res(0, 1) = 3;
  res(1, 0) = 6;
  res(1, 1) = -7;
  EXPECT_TRUE((a -= b) == res);
}

TEST(Matrix, SubMatrix6) {
  S21Matrix a(1, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  S21Matrix b(2, 2);
  b(0, 0) = 9;
  b(0, 1) = 0;
  b(1, 0) = -5;
  b(1, 1) = 6;
  EXPECT_THROW(a -= b, out_of_range);
}

TEST(Matrix, MulNumber1) {
  S21Matrix a(2, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  a(1, 0) = 0;
  a(1, 1) = -1;
  const double num = 2.5;
  S21Matrix res(2, 2);
  res(0, 0) = 10;
  res(0, 1) = 7.5;
  res(1, 0) = 0;
  res(1, 1) = -2.5;
  a.MulNumber(num);
  EXPECT_TRUE(a.EqMatrix(res));
}

TEST(Matrix, MulNumber2) {
  S21Matrix a(2, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  a(1, 0) = 0;
  a(1, 1) = -1;
  const double num = 2.5;

  S21Matrix res(2, 2);
  res(0, 0) = 10;
  res(0, 1) = 7.5;
  res(1, 0) = 0;
  res(1, 1) = -2.5;

  EXPECT_TRUE(a * num == res);
}

TEST(Matrix, MulNumber3) {
  S21Matrix a(2, 2);
  a(0, 0) = 4;
  a(0, 1) = 3;
  a(1, 0) = 0;
  a(1, 1) = -1;
  const double num = 2.5;
  S21Matrix res(2, 2);
  res(0, 0) = 10;
  res(0, 1) = 7.5;
  res(1, 0) = 0;
  res(1, 1) = -2.5;
  EXPECT_TRUE((a *= num) == res);
}

TEST(Matrix, MulMatrix1) {
  S21Matrix a(3, 2);
  a(0, 0) = 1;
  a(0, 1) = 4;
  a(1, 0) = 2;
  a(1, 1) = 5;
  a(2, 0) = 3;
  a(2, 1) = 6;
  S21Matrix b(3, 4);
  b(0, 0) = 1;
  b(0, 1) = -1;
  b(0, 2) = 1;
  b(0, 3) = 4;
  b(1, 0) = 2;
  b(1, 1) = 3;
  b(1, 2) = 4;
  b(1, 3) = -5;
  b(2, 0) = 1;
  b(2, 1) = 1;
  b(2, 2) = 1;
  b(2, 3) = 1;

  EXPECT_THROW(a.MulMatrix(b), out_of_range);
}

TEST(Matrix, MulMatrix2) {
  S21Matrix a(3, 2);
  S21Matrix b(2, 3);
  a(0, 0) = 1;
  a(0, 1) = 4;
  a(1, 0) = 2;
  a(1, 1) = 5;
  a(2, 0) = 3;
  a(2, 1) = 6;
  b(0, 0) = 1;
  b(0, 1) = -1;
  b(0, 2) = 1;
  b(1, 0) = 2;
  b(1, 1) = 3;
  b(1, 2) = 4;
  S21Matrix res(3, 3);
  res(0, 0) = 9;
  res(0, 1) = 11;
  res(0, 2) = 17;
  res(1, 0) = 12;
  res(1, 1) = 13;
  res(1, 2) = 22;
  res(2, 0) = 15;
  res(2, 1) = 15;
  res(2, 2) = 27;
  a.MulMatrix(b);
  EXPECT_TRUE(a.EqMatrix(res));
}

TEST(Matrix, MulMatrix3) {
  S21Matrix a(3, 2);
  a(0, 0) = 1;
  a(0, 1) = 4;
  a(1, 0) = 2;
  a(1, 1) = 5;
  a(2, 0) = 3;
  a(2, 1) = 6;
  S21Matrix b(3, 4);
  b(0, 0) = 1;
  b(0, 1) = -1;
  b(0, 2) = 1;
  b(0, 3) = 4;
  b(1, 0) = 2;
  b(1, 1) = 3;
  b(1, 2) = 4;
  b(1, 3) = -5;
  b(2, 0) = 1;
  b(2, 1) = 1;
  b(2, 2) = 1;
  b(2, 3) = 1;
  EXPECT_THROW(a * b, out_of_range);
}

TEST(Matrix, MulMatrix4) {
  S21Matrix a(3, 2);
  a(0, 0) = 1;
  a(0, 1) = 4;
  a(1, 0) = 2;
  a(1, 1) = 5;
  a(2, 0) = 3;
  a(2, 1) = 6;
  S21Matrix b(2, 3);
  b(0, 0) = 1;
  b(0, 1) = -1;
  b(0, 2) = 1;
  b(1, 0) = 2;
  b(1, 1) = 3;
  b(1, 2) = 4;
  S21Matrix res(3, 3);
  res(0, 0) = 9;
  res(0, 1) = 11;
  res(0, 2) = 17;
  res(1, 0) = 12;
  res(1, 1) = 13;
  res(1, 2) = 22;
  res(2, 0) = 15;
  res(2, 1) = 15;
  res(2, 2) = 27;
  EXPECT_TRUE(a * b == res);
}

TEST(Matrix, MulMatrix5) {
  S21Matrix a(3, 2);
  a(0, 0) = 1;
  a(0, 1) = 4;
  a(1, 0) = 2;
  a(1, 1) = 5;
  a(2, 0) = 3;
  a(2, 1) = 6;
  S21Matrix b(2, 3);
  b(0, 0) = 1;
  b(0, 1) = -1;
  b(0, 2) = 1;
  b(1, 0) = 2;
  b(1, 1) = 3;
  b(1, 2) = 4;
  S21Matrix res(3, 3);
  res(0, 0) = 9;
  res(0, 1) = 11;
  res(0, 2) = 17;
  res(1, 0) = 12;
  res(1, 1) = 13;
  res(1, 2) = 22;
  res(2, 0) = 15;
  res(2, 1) = 15;
  res(2, 2) = 27;
  EXPECT_TRUE((a *= b) == res);
}

TEST(Matrix, Transponse) {
  S21Matrix a(3, 2);
  a(0, 0) = 1;
  a(0, 1) = 4;
  a(1, 0) = 2;
  a(1, 1) = 5;
  a(2, 0) = 3;
  a(2, 1) = 6;
  S21Matrix res(2, 3);
  res(0, 0) = 1;
  res(0, 1) = 2;
  res(0, 2) = 3;
  res(1, 0) = 4;
  res(1, 1) = 5;
  res(1, 2) = 6;
  S21Matrix temp(2, 3);
  temp = a.Transpose();
  EXPECT_TRUE(res.EqMatrix(temp));
}

TEST(Matrix, CalcComplements1) {
  S21Matrix a(3, 3);
  a(0, 0) = 1;
  a(0, 1) = 2;
  a(0, 2) = 3;
  a(1, 0) = 0;
  a(1, 1) = 4;
  a(1, 2) = 2;
  a(2, 0) = 5;
  a(2, 1) = 2;
  a(2, 2) = 1;
  S21Matrix res(3, 3);
  res(0, 0) = 0;
  res(0, 1) = 10;
  res(0, 2) = -20;
  res(1, 0) = 4;
  res(1, 1) = -14;
  res(1, 2) = 8;
  res(2, 0) = -8;
  res(2, 1) = -2;
  res(2, 2) = 4;
  EXPECT_TRUE(res.EqMatrix(a.CalcComplements()));
}

TEST(Matrix, CalcComplements2) {
  S21Matrix mat(1, 2);
  EXPECT_THROW(mat.CalcComplements(), out_of_range);
}

TEST(Matrix, Determinant1) {
  S21Matrix a(4, 3);
  a(0, 0) = 1;
  a(0, 1) = 2;
  a(0, 2) = 3;
  a(1, 0) = 4;
  a(1, 1) = 5;
  a(1, 2) = 6;
  a(2, 0) = 7;
  a(2, 1) = 8;
  a(2, 2) = 9;
  a(3, 0) = 1;
  a(3, 1) = 1;
  a(3, 2) = 1;
  EXPECT_THROW(a.Determinant(), out_of_range);
}

TEST(Matrix, Determinant2) {
  S21Matrix a(3, 3);
  a(0, 0) = 1;
  a(0, 1) = 2;
  a(0, 2) = 3;
  a(1, 0) = 4;
  a(1, 1) = 5;
  a(1, 2) = 6;
  a(2, 0) = 7;
  a(2, 1) = 8;
  a(2, 2) = 9;
  EXPECT_EQ(a.Determinant(), 0);
}

TEST(Matrix, Determinant3) {
  S21Matrix a(3, 3);
  a(0, 0) = 1;
  a(0, 1) = 2;
  a(0, 2) = 3;
  a(1, 0) = 4;
  a(1, 1) = 5;
  a(1, 2) = 6;
  a(2, 0) = 7;
  a(2, 1) = 8;
  a(2, 2) = 10;
  EXPECT_DOUBLE_EQ(a.Determinant(), -3);
}

TEST(Matrix, Determinant4) {
  S21Matrix mat(1, 1);
  mat(0, 0) = 5;
  EXPECT_DOUBLE_EQ(mat(0, 0), mat.Determinant());
}

TEST(Matrix, Inverse1) {
  S21Matrix a(3, 3);
  a(0, 0) = 1;
  a(0, 1) = 2;
  a(0, 2) = 3;
  a(1, 0) = 4;
  a(1, 1) = 5;
  a(1, 2) = 6;
  a(2, 0) = 7;
  a(2, 1) = 8;
  a(2, 2) = 9;
  EXPECT_THROW(a.InverseMatrix(), out_of_range);
}

TEST(Matrix, Inverse2) {
  S21Matrix a(3, 3);
  a(0, 0) = 2;
  a(0, 1) = 5;
  a(0, 2) = 7;
  a(1, 0) = 6;
  a(1, 1) = 3;
  a(1, 2) = 4;
  a(2, 0) = 5;
  a(2, 1) = -2;
  a(2, 2) = -3;
  S21Matrix res(3, 3);
  res(0, 0) = 1;
  res(0, 1) = -1;
  res(0, 2) = 1;
  res(1, 0) = -38;
  res(1, 1) = 41;
  res(1, 2) = -34;
  res(2, 0) = 27;
  res(2, 1) = -29;
  res(2, 2) = 24;
  EXPECT_TRUE(res.EqMatrix(a.InverseMatrix()));
}

TEST(Matrix, Inverse3) {
  S21Matrix mat(1, 1);
  mat(0, 0) = 4;
  S21Matrix expected(1, 1);
  expected(0, 0) = 0.25;
  EXPECT_TRUE(mat.InverseMatrix().EqMatrix(expected));
}

TEST(Matrix, Inverse4) {
  S21Matrix mat(3, 4);
  EXPECT_THROW(mat.InverseMatrix(), out_of_range);
}

TEST(Matrix, Inverse5) {
  S21Matrix mat(2, 2);
  EXPECT_THROW(mat(-1, -2), out_of_range);
  EXPECT_THROW(mat.InverseMatrix(), out_of_range);
}

TEST(Matrix, Constructors) {
  S21Matrix left;
  left(0, 0) = 11.11;
  S21Matrix right(1, 1);
  right(0, 0) = 11.11;
  EXPECT_EQ(left, right);
}

TEST(Matrix, CopyConstructor) {
  S21Matrix matrix(2, 3);
  matrix(0, 0) = 1;
  matrix(0, 1) = 2;
  matrix(0, 2) = 3.1;
  matrix(1, 0) = 4.5;
  matrix(1, 1) = 5;
  matrix(1, 2) = 6.2;
  S21Matrix copyMatrix = matrix;
  ASSERT_EQ(copyMatrix.getRows(), 2);
  ASSERT_EQ(copyMatrix.getCols(), 3);
  for (int i = 0; i < 2; ++i)
    for (int j = 0; j < 3; ++j) ASSERT_EQ(matrix(i, j), copyMatrix(i, j));
}

TEST(Matrix, MoveConstructor) {
  S21Matrix matrix(2, 3);
  matrix(0, 0) = 1;
  matrix(0, 1) = 2;
  matrix(0, 2) = 3.1;
  matrix(1, 0) = 4.5;
  matrix(1, 1) = 5;
  matrix(1, 2) = 6.2;
  S21Matrix matrixTemp(matrix);
  S21Matrix copy(move(matrix));
  ASSERT_EQ(copy.getRows(), 2);
  ASSERT_EQ(copy.getCols(), 3);
  for (int i = 0; i < 2; ++i) {
    for (int j = 0; j < 3; ++j) {
      ASSERT_EQ(matrixTemp(i, j), copy(i, j));
    }
  }
}

TEST(Matrix, SetterGetter) {
  S21Matrix right(123, 911);
  EXPECT_EQ(right.getCols(), 911);
  EXPECT_EQ(right.getRows(), 123);
  right.setRows(2);
  EXPECT_EQ(right.getRows(), 2);
  right.setCols(20);
  EXPECT_EQ(right.getCols(), 20);
}

TEST(Matrix, invalid_argument) {
  S21Matrix mat(3, 3);
  EXPECT_THROW(mat(0, 5), out_of_range);
}

int main(int argc, char** argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
