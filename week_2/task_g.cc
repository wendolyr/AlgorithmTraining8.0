#include <iostream>
#include <vector>

int main() {
  int n = 0;
  std::cin >> n;

  std::vector<std::vector<long long>> dp(n + 1,
                                         std::vector<long long>(n + 1, 0));

  for (int j = 0; j <= n; ++j) {
    dp[0][j] = 1;
  }

  for (int i = 1; i <= n; ++i) {
    for (int j = 1; j <= n; ++j) {
      if (i >= j) {
        dp[i][j] += dp[i - j][j - 1];
      }

      dp[i][j] += dp[i][j - 1];
    }
  }

  std::cout << dp[n][n];

  return 0;
}