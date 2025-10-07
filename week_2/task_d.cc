#include <iostream>
#include <stack>
#include <string>
#include <unordered_map>
#include <vector>

int main() {
  std::string s;
  int n = 0;

  std::cin >> s;
  std::cin >> n;
  std::unordered_map<std::string, int> dict;
  for (int i = 0; i < n; ++i) {
    std::string temp;
    std::cin >> temp;
    ++dict[temp];
  }

  int l = s.size();
  std::vector<bool> dp(l + 1, false);
  std::vector<int> prev(l + 1, -1);
  dp[0] = true;

  for (int i = 1; i <= l; ++i) {
    for (int j = 0; j < i; ++j) {
      if (dp[j]) {
        std::string cur;
        for (int k = j; k < i; ++k) {
          cur += s[k];
        }
        if (dict.contains(cur)) {
          dp[i] = true;
          prev[i] = j;
          break;
        }
      }
    }
  }

  int pos = l;
  std::stack<std::string> res;
  while (pos > 0) {
    int beg = prev[pos];
    std::string cur;
    for (int i = beg; i < pos; ++i) {
      cur += s[i];
    }
    res.push(cur);
    pos = beg;
  }

  while (!res.empty()) {
    std::cout << res.top() << ' ';
    res.pop();
  }

  return 0;
}