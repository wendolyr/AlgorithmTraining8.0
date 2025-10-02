#include <iostream>
#include <list>
#include <unordered_map>

int main() {
  int n = 0, m = 0;
  std::string s;
  std::cin >> n >> m;
  std::cin >> s;

  std::unordered_map<std::string, std::list<int>> parts(m);
  for (int i = 1; i <= m; ++i) {
    std::string p;
    std::cin >> p;
    parts[p].push_back(i);
  }

  std::string search;
  for (int i = 0; i < n; ++i) {
    search += s[i];

    if (parts.contains(search)) {
      std::cout << parts[search].front() << ' ';
      parts[search].pop_front();
      search = "";
    }
  }

  return 0;
}