

all: clean test

test:
	g++ -Wall -Werror -Wextra -Wpedantic -Wshadow -std=c++23 task_f.cc -o t
	./t < test.txt

clean:
	rm -rf t