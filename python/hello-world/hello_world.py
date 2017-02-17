#
# Skeleton file for the Python "Hello World" exercise.
#
import time


def hello(name=''):
    if name:
        return "Hello, %s!" % name
    else:
        return "Hello, World!"


def main():
    name = ''
    print hello(name)

main()
