CC = g++
CXXFLAGS = -Wall -Werror -Wextra -std=c++23

all: clean test

test:
		$(CC) $(CXXFLAGS) task_h.cc -o t
		./t < text.txt

clean:
		rm -rf t