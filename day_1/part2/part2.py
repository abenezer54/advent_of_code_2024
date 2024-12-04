import os, sys,collections as c  
sys.stdin = open("./input.txt", "r")
sys.stdout = open("./output.txt", 'w')

input = lambda: sys.stdin.readline().strip()

a, b = [], []
for _ in range(1000):
    x, y = map(int, input().split())
    a.append(x)
    b.append(y)

cnt = c.Counter(b)

ans = 0
for i in range(1000):
    ans += a[i] * cnt[a[i]]
print(ans)