import os, sys,collections as c  
sys.stdin = open("./../input.txt", "r")
sys.stdout = open("./../output.txt", 'w')

input = lambda: sys.stdin.readline().strip()

ans = 0
for _ in range(1000):
    arr = list(map(int, input().split()))
    for removed in range(len(arr)):
        pos = neg = 0
        safe = True
        for i in range(len(arr) - 1):
            if i == removed:
                continue
            if i + 1 == removed:
                if i + 2 < len(arr):
                    if not (1 <= abs(arr[i] - arr[i + 2]) <= 3):
                        safe = False
                        break
                    if arr[i] - arr[i + 2] > 0:
                        pos += 1
                    else:
                        neg += 1
            else:
                if not (1 <= abs(arr[i] - arr[i + 1]) <= 3):
                    safe = False
                    break
                if arr[i] - arr[i + 1] > 0:
                    pos += 1
                else:
                    neg += 1

        if pos and neg:
            safe = False
        
        if safe:
            ans += 1
            break
    
        


print(ans)