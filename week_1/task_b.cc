#include <algorithm>
#include <climits>
#include <iomanip>
#include <iostream>

int main() {
  int a = 0, b = 0, c = 0;
  double v0 = 0, v1 = 0, v2 = 0;

  std::cin >> a >> b >> c >> v0 >> v1 >> v2;

  double t1 = 0, t2 = 0, t3 = 0, t4 = 0;

  t1 = t3 = c / v1;
  t2 = (a + b) * (1 / v0 + 1 / v1);

  if (a < b) {
    t1 += (b / v0 + a / v2);
  } else {
    t1 += (a / v0 + b / v2);
  }

  if (a + c < b) {
    t3 += a / v2 + (c + a) / v0;
    t4 = (2 * a + c) * (1 / v0 + 1 / v1);
  } else {
    t3 += b / v2 + (b + c) / v0;
    t4 = (2 * b + c) * (1 / v0 + 1 / v1);
  }

  std::cout << std::min({t1, t2, t3, t4});

  return 0;
}