#include <algorithm>
#include <iostream>
#include <vector>

int main() {
  int n = 0;
  std::cin >> n;

  std::vector<std::vector<char>> field(n + 1, std::vector<char>(3, ' '));
  for (int i = 1; i <= n; ++i) {
    for (int j = 0; j < 3; ++j) {
      std::cin >> field[i][j];
    }
  }

  std::vector<std::vector<int>> dp(n + 1, std::vector<int>(3, 0));

  int result = 0;
  for (int i = 1; i <= n; ++i) {
    bool can_move = false;

    if (field[i][0] != 'W') {
      dp[i][0] = std::max({dp[i - 1][0], dp[i - 1][1]});
      if (field[i - 1][0] != 'W' || field[i - 1][1] != 'W') {
        dp[i][0] += field[i][0] == 'C' ? 1 : 0;
        can_move = true;
      } else {
        field[i][0] = 'W';
      }
    }

    if (field[i][1] != 'W') {
      dp[i][1] = std::max({dp[i - 1][0], dp[i - 1][1], dp[i - 1][2]});
      if (field[i - 1][0] != 'W' || field[i - 1][1] != 'W' ||
          field[i - 1][2] != 'W') {
        dp[i][1] += field[i][1] == 'C' ? 1 : 0;
        can_move = true;
      } else {
        field[i][1] = 'W';
      }
    }

    if (field[i][2] != 'W') {
      dp[i][2] = std::max({dp[i - 1][1], dp[i - 1][2]});
      if (field[i - 1][1] != 'W' || field[i - 1][2] != 'W') {
        dp[i][2] += field[i][2] == 'C' ? 1 : 0;
        can_move = true;
      } else {
        field[i][2] = 'W';
      }
    }

    result = std::max({result, dp[i][0], dp[i][1], dp[i][2]});
    if (!can_move) {
      break;
    }
  }

  std::cout << result;

  return 0;
}