#include <iostream>

int main() {
  long long n = 0, k = 0;
  std::cin >> n >> k;

  if (k > 0) {
    n += n % 10;
    --k;
  }

  if (n % 10 == 0 || k == 0) {
    std::cout << n;
    return 0;
  }

  long long cycles = k / 4;

  n += cycles * 20;

  int rem = k - cycles * 4;

  while (rem > 0) {
    n += n % 10;
    --rem;
  }

  std::cout << n;

  return 0;
}

// 1 -> 02 04 08 16 22 24 28 36 42 44 48 56 // 14
// 2 -> 04 08 16 22 24 28 36 42 // 18
// 3 -> 06 12 14 18 26 32 34 38 // 12
// 4 -> 8 16 22 24 28 // 16

// 6 -> 12 14 18 26 32 34 38 46 // 14
// 7 -> 14 18 26 32 34 38 // 18
// 8 -> 16 22 24 28 36 // 12
// 9 -> 18 26 32 34 38 // 16