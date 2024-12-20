import os, sys,collections as c  
sys.stdin = open("./input.txt", "r")
sys.stdout = open("./output.txt", 'w')

input = lambda: sys.stdin.readline().strip()
a, b = [], []
for _ in range(1000):
    x, y = map(int, input().split())
    a.append(x)
    b.append(y)

a.sort()
b.sort()

ans = 0
for i in range(1000):
    ans += abs(a[i] - b[i])
print(ans)