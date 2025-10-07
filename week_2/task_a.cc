#include <iostream>
#include <vector>

int main() {
  int n = 0;
  std::cin >> n;

  switch (n) {
    case 1:
      std::cout << 1;
      break;
    case 2:
      std::cout << 2;
      break;
    case 3:
      std::cout << 4;
      break;
    default:
      std::vector<int> a(n + 1, 0);
      a[n] = 1;
      a[n - 1] = 2;
      a[n - 2] = 4;

      for (int i = n - 3; i > 0; --i) {
        a[i] = a[i + 1] + a[i + 2] + a[i + 3];
      }

      std::cout << a[1] << std::endl;
  }

  return 0;
}