#include <iostream>
#include <list>
#include <memory>
#include <string>
#include <unordered_map>

class MyList {
 public:
  struct Node {
    int val = 0;
    std::shared_ptr<Node> prev = nullptr;
    std::shared_ptr<Node> next = nullptr;

    Node() {}
    Node(int v) : val(v) {}
  };

  MyList(std::list<int>& other) {
    auto shared =
        std::make_shared<std::unordered_map<int, std::shared_ptr<Node>>>();
    fast_getter_ = shared;

    std::shared_ptr<Node> prev = nullptr;
    for (int i : other) {
      std::shared_ptr<Node> new_node = std::make_shared<Node>(i);
      (*fast_getter_.lock())[++size_] = new_node;
      if (!prev) {
        head_ = new_node;
        prev = new_node;
        continue;
      }

      prev->next = new_node;
      new_node->prev = prev;

      prev = new_node;
    }

    tail_ = prev;
  }

  MyList() {}

  MyList(const MyList& other) {
    head_ = other.head_;
    tail_ = other.tail_;
    is_sublist_ = other.is_sublist_;
    fast_getter_ = other.fast_getter_;
    start_ = other.start_;
    size_ = other.size_;
  }

  MyList(MyList&& other) {
    head_ = other.head_;
    tail_ = other.tail_;
    is_sublist_ = other.is_sublist_;
    fast_getter_ = other.fast_getter_;
    start_ = other.start_;
    size_ = other.size_;
    other.head_ = nullptr;
    other.tail_ = nullptr;
  }

  MyList& operator=(const MyList& other) {
    head_ = other.head_;
    tail_ = other.tail_;
    is_sublist_ = other.is_sublist_;
    fast_getter_ = other.fast_getter_;
    start_ = other.start_;
    size_ = other.size_;

    return *this;
  }

  MyList(const MyList& other, int from, int to) {
    is_sublist_ = true;
    fast_getter_ = other.fast_getter_;
    start_ = other.start_ + from - 1;
    head_ = other.fast_getter_.lock()->at(start_);
    tail_ = other.fast_getter_.lock()->at(start_ + to - from);
  }

  void Set(int i, int v) {
    int pos = start_ + i - 1;
    (*fast_getter_.lock())[pos] -> val = v;
  }

  int Get(int i) {
    int pos = start_ + i - 1;
    return (*fast_getter_.lock())[pos] -> val;
  }

  void Add(int x) {
    if (is_sublist_) {
      return;
    }

    std::shared_ptr<Node> new_node = std::make_shared<Node>(x);
    tail_->next = new_node;
    new_node->prev = tail_;
    tail_ = new_node;

    (*fast_getter_.lock())[++size_] = new_node;
  }

 private:
  std::shared_ptr<Node> head_ = nullptr;
  std::shared_ptr<Node> tail_ = nullptr;
  bool is_sublist_ = false;
  std::weak_ptr<std::unordered_map<int, std::shared_ptr<Node>>> fast_getter_;
  int start_ = 1;
  size_t size_ = 0;
};

int main() {
  std::string command;

  std::unordered_map<std::string, MyList> lists;
  int n = 0;
  std::cin >> n;

  for (int k = 0; k <= n; ++k) {
    std::getline(std::cin, command);
    size_t pos;

    if ((pos = command.find("new List")) != std::string::npos) {
      std::string name;
      for (size_t i = 5; command[i] != ' '; ++i) {
        name += command[i];
      }

      std::list<int> temp;

      pos += 9;

      std::string num;
      for (size_t i = pos;; ++i) {
        if (command[i] >= '0' && command[i] <= '9') {
          num += command[i];
          continue;
        }

        temp.push_back(std::stoi(num));
        num = "";
        if (command[i] == ')') {
          break;
        }
      }

      lists[name] = MyList(temp);
    } else if ((pos = command.find("subList")) != std::string::npos) {
      std::string name;
      for (size_t i = 5; command[i] != ' '; ++i) {
        name += command[i];
      }

      std::string parent;

      for (size_t i = 8 + name.size(); command[i] != '.'; ++i) {
        parent += command[i];
      }

      pos += 8;

      int from = -1, to = -1;
      std::string num;
      for (size_t i = pos;; ++i) {
        if (command[i] >= '0' && command[i] <= '9') {
          num += command[i];
          continue;
        }

        if (from == -1) {
          from = std::stoi(num);
        } else {
          to = std::stoi(num);
        }
        num = "";
        if (command[i] == ')') {
          break;
        }
      }

      lists[name] = MyList(lists[parent], from, to);
    } else if ((pos = command.find(".get")) != std::string::npos) {
      std::string name;
      for (size_t i = 0; i < pos; ++i) {
        name += command[i];
      }

      pos += 5;
      std::string num;
      int x = 0;
      for (size_t i = pos;; ++i) {
        if (command[i] >= '0' && command[i] <= '9') {
          num += command[i];
          continue;
        }

        x = std::stoi(num);
        break;
      }
      std::cout << lists[name].Get(x) << std::endl;
    } else if ((pos = command.find(".add")) != std::string::npos) {
      std::string name;
      for (size_t i = 0; i < pos; ++i) {
        name += command[i];
      }

      pos += 5;
      std::string num;
      int x = 0;
      for (size_t i = pos;; ++i) {
        if (command[i] >= '0' && command[i] <= '9') {
          num += command[i];
          continue;
        }

        x = std::stoi(num);
        break;
      }

      lists[name].Add(x);
    } else if ((pos = command.find(".set")) != std::string::npos) {
      std::string name;
      for (size_t i = 0; i < pos; ++i) {
        name += command[i];
      }

      pos += 5;
      std::string num;
      int p = -1;
      int x = 0;

      for (size_t i = pos;; ++i) {
        if (command[i] >= '0' && command[i] <= '9') {
          num += command[i];
          continue;
        }

        if (p == -1) {
          p = std::stoi(num);

        } else {
          x = std::stoi(num);
        }
        num = "";
        if (command[i] == ')') {
          break;
        }
      }

      lists[name].Set(p, x);
    }
  }
}