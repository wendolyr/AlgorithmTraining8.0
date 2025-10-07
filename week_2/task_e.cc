#include <iostream>
#include <stack>
#include <vector>

int main() {
  int N = 0, K = 0;
  std::cin >> N >> K;
  std::vector<int> pillars(N + 1, 0);

  for (int i = 1; i <= N; ++i) {
    std::cin >> pillars[i];
  }

  std::vector<int> pref_sum(N + 1, 0);

  for (int i = 1; i <= N; ++i) {
    pref_sum[i] = pref_sum[i - 1] + pillars[i];
  }

  std::vector<int> dp(N + 1, 0);
  std::vector<int> prev(N + 1, -1);
  std::vector<bool> is_tower_end(N + 1, false);

  for (int i = 1; i <= N; ++i) {
    dp[i] = dp[i - 1];
    prev[i] = i - 1;

    if (i >= K) {
      int sum = pref_sum[i] - pref_sum[i - K];
      int min_h = pillars[i];
      for (int j = i - K + 1; j < i; ++j) {
        min_h = std::min(min_h, pillars[j]);
      }

      int power = sum * min_h + dp[i - K];
      if (power > dp[i]) {
        dp[i] = power;
        prev[i] = i - K;
        is_tower_end[i] = true;
      }
    }
  }

  std::stack<int> res;
  for (int i = N; i > 0;) {
    if (is_tower_end[i]) {
      res.push(i - K + 1);
    }
    i = prev[i];
  }

  std::cout << res.size() << std::endl;
  while (!res.empty()) {
    std::cout << res.top() << ' ';
    res.pop();
  }

  return 0;
}