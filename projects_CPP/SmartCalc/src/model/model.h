#ifndef SRC_MODEL_MODEL_H
#define SRC_MODEL_MODEL_H

#include <cmath>
#include <cstring>
#include <iostream>
#include <list>
#include <vector>

#define EPS 2.71828182846

namespace s21 {

class Model {
 private:
  struct Lexeme {
   public:
    double value_;
    int operation_;
    int priority_;

    double GetValue() { return value_; }
    int GetOperation() { return operation_; }
    int GetPriority() { return priority_; }

    Lexeme(double value, int operation, int priority)
        : value_(value), operation_(operation), priority_(priority) {}
  };

 public:
  int Calculate(std::string &str, double x);
  std::list<Lexeme> Parser(const std::string &str);
  void PolisNotation(std::list<Lexeme> &List);
  int GetOperators(const char *str);
  int GetPriorities(int oper);
  void Replace(std::string &str, const std::string a, const std::string b);
  void ReplaceDot(std::string &str, const std::string a, const std::string b,
                  const std::string c);
  void ValueX(double value, std::list<Lexeme> &List);
  double Counter(std::list<Lexeme> &List);
  void UnaryOperations(int oper, std::list<double> &res);
  void Operations(int oper, std::list<double> &res);
  void Functions(int oper, std::list<double> &res);
  void Check(std::string &s);
  void CheckE(std::string &s);
  bool CheckPriority(std::list<Lexeme> &support, int i);
  int CheckFunctions(const std::string &str);
  bool CheckDot(const char *str);
  bool CheckBracket(const char *str);
  bool CheckSymbol(const char *s);
  bool Checkk(const char s);
  bool CheckSign(const char s);
  bool CheckMod(const char *str);
  bool CheckSqrt(const char *s);
  std::pair<std::vector<double>, std::vector<double>> Grafic(
      double min, double max, const std::string &str);
  void CreditAnnuity(double sum_of_credit, double duration_of_credit,
                     double annual_interest_rate, double &month_pay,
                     double &over_pay, double &all_sum_of_pay);
  void CreditDifferentiated(double sum_of_credit, double duration_of_credit,
                            double annual_interest_rate,
                            double &month_pay_first, double &month_pay_last,
                            double &over_pay, double &all_sum_of_pay);
  void Deposit(double replenishment1[], double replenishment2[],
               int count_replenishment, double withdrawals1[],
               double withdrawals2[], int count_withdrawals, double all_sum,
               double duration_of_credit, double interest_rate, double tax,
               int period_s, int capitalization, double &all_interest,
               double &all_tax, double &total_sum);
  double ReplenishmentList(double replenishment1[],
                           double const replenishment2[],
                           int count_replenishment, int count_month);
  double WithdrawalsList(double withdrawals1[], double const withdrawals2[],
                         int count_withdrawals, int count_month,
                         double all_sum);
};

}  // namespace s21
#endif  // SRC_MODEL_MODEL_H
