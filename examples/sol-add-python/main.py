
n = int(input())
elems = input().strip().split()
assert len(elems) == n

summ = 0
for e in elems:
    summ += int(e)
print(summ)
