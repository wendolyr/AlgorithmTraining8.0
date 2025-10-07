#include <iostream>
#include <string>

int main() {
  std::string river;
  std::cin >> river;

  int l_min = 0, r_min = 1;
  for (char c : river) {
    if (c == 'L') {
      r_min = std::min(r_min, l_min + 1);
      l_min = std::min(l_min + 1, r_min + 1);
    } else if (c == 'R') {
      l_min = std::min(l_min, r_min + 1);
      r_min = std::min(r_min + 1, l_min + 1);
    } else {
      ++l_min;
      ++r_min;
    }
  }

  std::cout << std::min(l_min + 1, r_min);

  return 0;
}