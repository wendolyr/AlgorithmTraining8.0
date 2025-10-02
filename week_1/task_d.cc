#include <iostream>
#include <unordered_map>
#include <vector>

int main() {
  int n = 0, k = 0;

  std::cin >> n >> k;

  std::unordered_map<int, int> um;
  for (int i = 0; i < n; ++i) {
    int inp = 0;
    std::cin >> inp;
    ++um[inp];
  }

  std::vector<int> ans;
  int l = um.size();
  while (k) {
    for (auto it = um.begin(); it != um.end() && k;) {
      if (it->second) {
        ans.push_back(it->first);
        --it->second;
        --k;
        if (k <= l) {
          ++it;
        }
      } else {
        ++it;
      }
    }
  }

  for (int i : ans) {
    std::cout << i << ' ';
  }

  return 0;
}