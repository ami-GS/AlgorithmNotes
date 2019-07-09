#include <iostream>
#include <chrono>
#include <string>
#include <sstream>

//17700?
void fizzbuzz0(int n) {
    for (int i = 0; i < n; i++) {
        if (i % 15 == 0) {
            std::cout << "FizzBuzz" << std::endl;
        } else if (i % 3 == 0) {
            std::cout << "Fizz" << std::endl;
        } else if (i % 5 == 0) {
            std::cout << "Buzz" << std::endl;
        } else {
            std::cout << i << std::endl;
        }
    }
}

//15593.682
void fizzbuzz1(int n) {
    for (int i = 0; i < n; ++i) {
        bool printed = false;
        if (i % 3 == 0) {
            std::cout << "Fizz";
            printed = true;
        }
        if (i % 5 == 0) {
            std::cout << "Buzz";
            printed = true;
        }
        if (!printed) {
            std::cout << i;
        }
        std::cout << std::endl;
    }
}

//18297.462000[ms]
void fizzbuzz2(int n) {
    for (int i = 0; i < n; i++) {
        bool printed = false;
        if (i % 3 == 0) {
            std::cerr << "Fizz";
            printed = true;
        }
        if (i % 5 == 0) {
            std::cerr << "Buzz";
            printed = true;
        }
        if (!printed) {
            std::cerr << i;
        }
        std::cout << std::endl;
    }
}

// string stream
void fizzbuzz3(int n) {
    std::stringstream s;
    for (int i = 0; i < n; i++) {
        bool printed = false;
        if (i % 3 == 0) {
            s << "Fizz";
            printed = true;
        }
        if (i % 5 == 0) {
            s << "Buzz";
            printed = true;
        }
        if (!printed) {
            //std::cout << i;
            s << i;
            //s << std::to_string(i);
        }
        s << std::endl;
    }
    std::cout << s.str();
}


// alloc first
void fizzbuzz4(int n) {
    std::string s(5*n, '\0');
    std::string fizz("Fizz");
    std::string buzz("Buzz");
    int sIdx = 0;
    for (int i = 1; i <= n; i++) {
        bool printed = false;
        if (i % 3 == 0) {
            memcpy(&s[sIdx], &fizz[0], 4);
            sIdx+=4;
            printed = true;
        }
        if (i % 5 == 0) {
            memcpy(&s[sIdx], &buzz[0], 4);
            sIdx+=4;
            printed = true;
        }
        if (!printed) {
            auto tmp = std::to_string(i);
            memcpy(&s[sIdx], &tmp[0], tmp.size());
            sIdx += tmp.size();
        }
        s[sIdx++] = '\n';
    }
    std::cout << s;
}

// alloc first
void fizzbuzz5(int n) {
    std::string s(5*n, '\0');
    const std::string fizz("Fizz");
    const std::string buzz("Buzz");
    const std::string fizzbuzz("FizzBuzz");
    int sIdx = 0;
    for (int i = 1; i <= n; i++) {
        if (i % 15 == 0) {
            memcpy(&s[sIdx], &fizzbuzz[0], 8);
            sIdx+=8;
        } else if (i % 5 == 0) {
            memcpy(&s[sIdx], &buzz[0], 4);
            sIdx+=4;
        } else if (i % 3 == 0) {
            memcpy(&s[sIdx], &fizz[0], 4);
            sIdx+=4;
        } else {
            auto tmp = std::to_string(i);
            memcpy(&s[sIdx], &tmp[0], tmp.size());
            sIdx += tmp.size();
        }
        s[sIdx++] = '\n';
    }
    std::cout << s;
}

int main() {
    std::chrono::system_clock::time_point start, end;

    start = std::chrono::system_clock::now();
    for (int i = 0; i < 10000; i++) {
        fizzbuzz5(10000);
    }
    end = std::chrono::system_clock::now();
    double time = static_cast<double>(std::chrono::duration_cast<std::chrono::microseconds>(end - start).count() / 1000.0);
    printf("time %lf[ms]\n", time);
    return 0;
}
