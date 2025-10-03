#include <iostream>

long long calculate1(int x, int y, int f, int g) {
  long long result = 0;
  int dir_x = 0;

  if (x < f) {
    dir_x = 1;
    ++x;
    while (x < f) {
      result += 3;
      ++x;
    }
  } else if (x > f) {
    dir_x = -1;
    --x;
    while (x > f) {
      result += 3;
      --x;
    }
  }

  if (y < g) {
    if (dir_x == 1) {
      result += 1;
    } else if (dir_x == -1) {
      result += 3;
    }

    ++y;

    while (y < g) {
      result += 3;
      ++y;
    }
  } else if (y > g) {
    if (dir_x == 1) {
      result += 3;
    } else if (dir_x == -1) {
      result += 1;
    }
    --y;

    while (y > g) {
      result += 3;
      --y;
    }
  }

  return result;
}

long long calculate2(int x, int y, int f, int g) {
  long long result = 0;
  int dir_y = 0;

  if (y < g) {
    dir_y = 1;
    ++y;

    while (y < g) {
      result += 3;
      ++y;
    }
  } else if (y > g) {
    dir_y = -1;
    --y;

    while (y > g) {
      result += 3;
      --y;
    }
  }

  if (x < f) {
    if (dir_y == 1) {
      result += 3;
    } else if (dir_y == -1) {
      result += 1;
    }

    ++x;
    while (x < f) {
      result += 3;
      ++x;
    }
  } else if (x > f) {
    if (dir_y == 1) {
      result += 1;
    } else if (dir_y == -1) {
      result += 3;
    }
    --x;
    while (x > f) {
      result += 3;
      --x;
    }
  }

  return result;
}

int main() {
  int x = 0, y = 0, f = 0, g = 0;

  std::cin >> x >> y;
  std::cin >> f >> g;

  std::cout << std::min(calculate1(x, y, f, g), calculate2(x, y, f, g));

  return 0;
}