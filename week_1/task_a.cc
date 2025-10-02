#include <climits>
#include <iostream>
#include <vector>

int main() {
  int vasya_sum = 0, masha_sum = 0;
  int masha_max = 0, vasya_min = INT_MAX;

  int n = 0;
  std::cin >> n;
  std::vector<int> a(n);

  for (int i = 0; i < n; ++i) {
    std::cin >> a[i];
  }

  for (int i = 0; i < n; ++i) {
    if (i % 2 == 0) {
      vasya_sum += a[i];
      vasya_min = std::min(vasya_min, a[i]);
    } else {
      masha_sum += a[i];
      masha_max = std::max(masha_max, a[i]);
    }
  }

  if (vasya_min < masha_max) {
    vasya_sum += (masha_max - vasya_min);
    masha_sum += (vasya_min - masha_max);
  }

  std::cout << vasya_sum - masha_sum;

  return 0;
}