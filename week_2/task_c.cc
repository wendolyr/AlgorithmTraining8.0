#include <algorithm>
#include <iostream>
#include <vector>

struct Interval {
  double begin, end, weight;
};

int main() {
  int n = 0;
  std::cin >> n;
  if (n == 0) {
    std::cout << 0.0;
    return 0;
  }

  std::vector<Interval> intervals(n);

  for (int i = 0; i < n; ++i) {
    std::cin >> intervals[i].begin >> intervals[i].end >> intervals[i].weight;
  }

  std::sort(intervals.begin(), intervals.end(),
            [](auto first, auto second) { return first.end < second.end; });

  std::vector<double> dp(n, 0.0);
  dp[0] = intervals[0].weight;

  for (int i = 1; i < n; ++i) {
    double current = intervals[i].weight;
    int left = 0, right = i - 1;
    int last = -1;
    while (left <= right) {
      int mid = left + (right - left) / 2;
      if (intervals[mid].end <= intervals[i].begin) {
        left = mid + 1;
        last = mid;
      } else {
        right = mid - 1;
      }
    }

    if (last != -1) {
      current += dp[last];
    }

    dp[i] = std::max(current, dp[i - 1]);
  }

  std::cout << dp[n - 1];

  return 0;
}