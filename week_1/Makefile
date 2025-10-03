

all: clean test

test:
	g++ -Wall -Wextra -Wpedantic -Wshadow -std=c++23 task_j.cc -o t
	./t < test.txt

clean:
	rm -rf t