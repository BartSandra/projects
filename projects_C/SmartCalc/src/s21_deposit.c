#include "s21_SmartCalc.h"

int s21_deposit(double *replenishment1, double *replenishment2,
                int count_replenishment, double *withdrawals1,
                double *withdrawals2, int count_withdrawals, double all_sum,
                double duration_of_credit, double interest_rate, double tax,
                int period_s, int capitalization, double *all_interest,
                double *all_tax, double *total_sum) {
  int error = 0;
  *all_interest = 0.0;
  *total_sum = 0.0;
  *all_tax = 0.0;
  interest_rate = interest_rate / 100;
  tax = tax / 100;
  if (isnan(*all_interest) != 0 || isinf(*all_interest) != 0 ||
      isnan(*total_sum) != 0 || isinf(*total_sum) != 0 ||
      isnan(*all_tax) != 0 || isinf(*all_tax) != 0) {
    error = 1;
  }
  if (capitalization == 0) {
    for (int i = 1; i <= duration_of_credit; i++) {
      all_sum = all_sum + s21_replenishment_list(replenishment1, replenishment2,
                                                 count_replenishment, i);
      all_sum = all_sum - s21_withdrawals_list(withdrawals1, withdrawals2,
                                               count_withdrawals, i, all_sum);
      *all_interest += interest_rate / 12 * all_sum;
    }
    *total_sum = all_sum;
  } else {
    int p = 0;
    for (int i = 1; i <= duration_of_credit; i++) {
      all_sum = all_sum + s21_replenishment_list(replenishment1, replenishment2,
                                                 count_replenishment, i);
      all_sum = all_sum - s21_withdrawals_list(withdrawals1, withdrawals2,
                                               count_withdrawals, i, all_sum);
      if (i % period_s == 0) {
        *all_interest +=
            (interest_rate / 12.0) * period_s * (all_sum + *all_interest);
        p = i;
      } else if (i == duration_of_credit) {
        int z = i - p;
        *all_interest += (interest_rate / 12.0) * z * (all_sum + *all_interest);
      }
    }
    *total_sum = *all_interest + all_sum;
  }
  *all_tax = ((*all_interest / duration_of_credit) - (95000 / 12.0)) *
             duration_of_credit * tax;

  *all_tax = *all_tax < 0 ? 0.0 : *all_tax;
  return error;
}

double s21_replenishment_list(double *replenishment1,
                              double const *replenishment2,
                              int count_replenishment, int count_month) {
  double buf = 0.0;
  for (int i = 0; i < count_replenishment; i++) {
    int z = (int)replenishment1[i];
    if (z == count_month) {
      buf += replenishment2[i];
    }
  }
  return buf;
}

double s21_withdrawals_list(double *withdrawals1, double const *withdrawals2,
                            int count_withdrawals, int count_month,
                            double all_sum) {
  double buf = 0.0;
  for (int i = 0; i < count_withdrawals; i++) {
    int z = (int)withdrawals1[i];
    if (z == count_month) {
      buf += withdrawals2[i];
    }
  }
  if (buf > all_sum) buf = all_sum;
  return buf;
}
