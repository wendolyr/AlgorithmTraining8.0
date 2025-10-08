#include <algorithm>
#include <climits>
#include <iostream>
#include <stack>
#include <vector>

struct Shop {
  int price = 0, sale_amount = 0, discount_price = 0, avail_amount = 0;
};

int cost(Shop& p, int need) {
  if (need > p.avail_amount) {
    return INT_MAX;
  }

  int result = 0;
  if (need < p.sale_amount) {
    result = p.price * need;
  } else {
    result = p.discount_price * need;
  }

  return result;
}

int main() {
  int n = 0, l = 0;
  std::cin >> n >> l;

  if (l == 0) {
    std::cout << 0 << std::endl;
    for (int i = 1; i <= n; ++i) {
      std::cout << 0 << ' ';
    }
    return 0;
  }

  std::vector<Shop> p(n + 1);
  int max_sale_amount = 0;
  for (int i = 1; i <= n; ++i) {
    std::cin >> p[i].price >> p[i].sale_amount >> p[i].discount_price >>
        p[i].avail_amount;

    max_sale_amount = std::max(max_sale_amount, p[i].sale_amount);
  }

  int max_cols = l + max_sale_amount;

  std::vector<std::vector<long long>> dp(
      n + 1, std::vector<long long>(max_cols + 1, INT_MAX));

  std::vector<std::vector<int>> parent(n + 1,
                                       std::vector<int>(max_cols + 1, 0));

  for (int i = 0; i <= n; ++i) {
    dp[i][0] = 0;
  }

  for (int i = 0; i <= max_cols; ++i) {
    long long c = cost(p[1], i);
    if (c != INT_MAX) {
      dp[1][i] = c;
      parent[1][i] = i;
    }
  }

  for (int i = 2; i <= n; ++i) {
    for (int j = 1; j <= max_cols; ++j) {
      if (dp[i - 1][j] < dp[i][j]) {
        dp[i][j] = dp[i - 1][j];
        parent[i][j] = 0;
      }

      for (int k = 1; k <= j; ++k) {
        long long c = cost(p[i], k);
        if (c != INT_MAX && dp[i - 1][j - k] != INT_MAX) {
          long long new_c = dp[i - 1][j - k] + c;
          if (new_c < dp[i][j]) {
            dp[i][j] = new_c;
            parent[i][j] = k;
          }
        }
      }
    }
  }

  long long result = INT_MAX;
  int best = -1;
  for (int i = l; i <= max_cols; ++i) {
    if (dp[n][i] < result) {
      result = dp[n][i];
      best = i;
    }
  }

  if (result == INT_MAX) {
    std::cout << -1;
    return 0;
  }

  std::cout << result << std::endl;

  std::stack<long long> ans_recover;
  int current = best;
  for (int i = n; i >= 1; --i) {
    ans_recover.push(parent[i][current]);
    current -= parent[i][current];
  }

  while (!ans_recover.empty()) {
    std::cout << ans_recover.top() << ' ';
    ans_recover.pop();
  }

  return 0;
}