#include <bits/stdc++.h>

class Base {
private:
	int *a;
public:
	Base() {
		std::cout << "Base::Constructor\n";
		a = new int(10);
	}
	~Base() {
		delete a;
		std::cout << "Base::Destructor\n";
	}
};

template<typename T>
class shared_ptr {
private:
	T* val;
	int* cnt;
public:
	shared_ptr() = default;

	shared_ptr(T* v) {
		val = v;
		cnt = new int(1);
	}

	shared_ptr(const shared_ptr<T>& v) {
		std::cout << "shared_ptr::Copy Constructor\n";
		this->val = v.val;
		this->cnt = v.cnt;
		++(*v.cnt);
	}

	~shared_ptr() {
		--(*cnt);
		std::cout << "shared_ptr::Deconstructing with " << *cnt << " left\n";
		if (*cnt > 0) return;
		delete val;
		delete cnt;
	}
};

int main() {
	shared_ptr<Base> a(new Base());
	shared_ptr<Base> b(a);

	return 0;
}