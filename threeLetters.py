# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(a, b):
    # Implement your solution here
    res = ""
    while a > b:
        if a - b > 1: 
            res += "aa"
            a -= 2
        else:
            res += "a"
            a -= 1

        if b :
            res += "b"
            b -= 1
    while b > a:
        if b - a > 1:
            res += "bb"
            b -= 2
        else:
            res += "b"
            b -= 1

        if a:
            res += "a"
            a -= 1

    for i in range(a):
        res += "a"
        res += "b"
    return res