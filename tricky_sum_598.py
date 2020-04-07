import math


def get_pow_2_sum(count):
    max_power = math.log(count, 2)
    return 2**(int(max_power) + 1) - 1

def calculate_sum(count):
    s = (count * (count + 1)) / 2
    p = 1
    while p <= count:
        s -= p * 2
        p *= 2
    return s

if __name__ == '__main__':
    inputs = int(input())
    for i in range(inputs):
        print(int(calculate_sum(int(input()))))

    