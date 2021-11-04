#include <bits/stdc++.h>

template<typename T>
T* find_if(T* begin, T* end, std::function<bool(const T&)> fn) {
	std::cout << sizeof(fn) << "\n";
	for (; begin != end; begin++) {
		if (fn(*begin)) return begin;
	}
	return end;
}

int main() {
	constexpr int N = 10;
	int* a = new int[N];
	std::iota(a, a+N, 0);

	auto it = find_if<int>(a, a+N, [](const int& v) -> bool {
		return v == 5;
	});
	if (it == a+N) std::cout << "not found";
	else std::cout << "index " << (it - a);

	delete [] a;
	return 0;
}
