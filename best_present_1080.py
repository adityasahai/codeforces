
import sys

def calculate(a, b):
    return a*(a+1) - (b*b)

def formula(pos):
    ceil, floor = int(pos/2) + 1, int(pos/2)
    if pos % 2 == 0:
        return calculate(ceil, floor)
    else:
        return calculate(floor, ceil)

if __name__ == "__main__":
    numLines = int(input())
    for i in range(numLines):
        inputs = input().split(" ")
        e, s = int(inputs[0]), int(inputs[1])
        print(int(formula(e) - formula(s-1)))