#include <algorithm>
#include <iostream>
#include <vector>

struct Elem {
  int x, y, val;
};

int main() {
  int n = 0, m = 0;
  std::cin >> n >> m;

  std::vector<std::vector<int>> table(n, std::vector<int>(m, 0));
  std::vector<Elem> elems;

  for (int i = 0; i < n; ++i) {
    for (int j = 0; j < m; ++j) {
      std::cin >> table[i][j];
      elems.emplace_back(j, i, table[i][j]);
    }
  }

  std::sort(elems.begin(), elems.end(),
            [](auto first, auto second) { return first.val < second.val; });

  int result = 1;

  int dx[4] = {0, 0, -1, 1};
  int dy[4] = {-1, 1, 0, 0};
  std::vector<std::vector<int>> dp(n, std::vector<int>(m, 1));
  int l = elems.size();
  for (int i = 0; i < l; ++i) {
    int x = elems[i].x;
    int y = elems[i].y;

    for (int j = 0; j < 4; ++j) {
      int nx = x + dx[j];
      int ny = y + dy[j];

      if (nx >= 0 && nx < m && ny >= 0 && ny < n) {
        if (table[ny][nx] == table[y][x] - 1) {
          dp[y][x] = std::max(dp[y][x], dp[ny][nx] + 1);
        }
      }
    }

    result = std::max(result, dp[y][x]);
  }

  std::cout << result;

  return 0;
}