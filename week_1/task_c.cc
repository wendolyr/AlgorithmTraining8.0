#include <iostream>
#include <unordered_map>

int main() {
  std::string s;
  std::cin >> s;

  long long l = s.size();

  std::unordered_map<char, int> letters;

  for (long long i = 0; i < l; ++i) {
    ++letters[s[i]];
  }

  long long total = l * (l - 1) / 2;
  long long repeat = 0;

  for (auto [c, freq] : letters) {
    repeat += 1LL * freq * (freq - 1) / 2;
  }

  std::cout << 1LL + (total - repeat);

  return 0;
}