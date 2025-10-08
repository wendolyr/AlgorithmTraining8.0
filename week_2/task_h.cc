#include <iostream>
#include <vector>

int main() {
  int n = 0;
  std::cin >> n;
  std::vector<bool> prime(n + 1, true);
  prime[0] = prime[1] = false;

  for (int i = 2; i <= n; ++i) {
    if (prime[i]) {
      for (int j = 2 * i; j <= n; j = j + i) {
        prime[j] = false;
      }
    }
  }

  std::vector<bool> dp(n + 1, false);

  for (int i = 1; i <= n; ++i) {
    for (int j = 1; j <= 3; ++j) {
      int rem = i - j;
      if (rem == 0 || !prime[rem]) {
        if (!dp[rem]) {
          dp[i] = true;
          break;
        }
      }
    }
  }

  if (dp[n]) {
    std::cout << 1;
  } else {
    std::cout << 2;
  }

  return 0;
}