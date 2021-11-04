import numpy as np
n = 10000
# blue = 0, red = 1
a = np.random.randint(0, 2, n)
while len(a) > 1:
	u = np.random.randint(0, n)
	np.delete(a, u)
	print(u)
	pass
print(a)