#include <climits>
#include <iostream>
#include <vector>

int findMaxRowSum(std::vector<std::vector<char>>& table,
                  int& index_max_row_sum) {
  int m = table.size(), n = table[0].size();
  int max_row_sum = INT_MIN;
  for (int i = 0; i < m; ++i) {
    int current_sum = 0;
    for (int j = 0; j < n; ++j) {
      if (table[i][j] == '+' || table[i][j] == '?') {
        current_sum += 1;
      } else {
        current_sum -= 1;
      }
    }

    if (current_sum > max_row_sum) {
      max_row_sum = current_sum;
      index_max_row_sum = i;
    }
  }

  return max_row_sum;
}

int findMinColSum(std::vector<std::vector<char>>& table,
                  int& index_min_col_sum) {
  int m = table.size(), n = table[0].size();
  int min_col_sum = INT_MAX;
  for (int i = 0; i < n; ++i) {
    int current_sum = 0;
    for (int j = 0; j < m; ++j) {
      if (table[j][i] == '-' || table[j][i] == '?') {
        current_sum -= 1;
      } else {
        current_sum += 1;
      }
    }

    if (current_sum < min_col_sum) {
      min_col_sum = current_sum;
      index_min_col_sum = i;
    }
  }

  return min_col_sum;
}

int calculate(std::vector<std::vector<char>> table, int option) {
  int m = table.size(), n = table[0].size();

  int max_row_sum = 0, min_col_sum = 0;
  if (option == 1) {
    int index_max_row_sum = 0;
    int index_min_col_sum = 0;

    max_row_sum = findMaxRowSum(table, index_max_row_sum);

    for (int i = 0; i < n; ++i) {
      if (table[index_max_row_sum][i] == '?') {
        table[index_max_row_sum][i] = '+';
      }
    }

    min_col_sum = findMinColSum(table, index_min_col_sum);
  } else if (option == 2) {
    int index_min_col_sum = 0;
    int index_max_row_sum = 0;

    min_col_sum = findMinColSum(table, index_min_col_sum);

    for (int i = 0; i < m; ++i) {
      if (table[i][index_min_col_sum] == '?') {
        table[i][index_min_col_sum] = '-';
      }
    }

    max_row_sum = findMaxRowSum(table, index_max_row_sum);
  }

  return max_row_sum - min_col_sum;
}

int main() {
  int m = 0, n = 0;

  std::cin >> m >> n;

  std::vector<std::vector<char>> table(m, std::vector<char>(n));

  for (int i = 0; i < m; ++i) {
    for (int j = 0; j < n; ++j) {
      std::cin >> table[i][j];
    }
  }

  std::cout << std::max(calculate(table, 1), calculate(table, 2));

  return 0;
}