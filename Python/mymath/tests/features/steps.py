from lettuce import *

@step
def have_the_number(step, number):
    'I have the number (\d+)'
    world.number = int(number)

@step
def i_compute_its_factorial(step):
    world.number = factorial(world.number)

@step('I see the number (\d+)')
def check_number(step, expected):
    expected = int(expected)
    assert world.number == expected, \
        "Got %d" % world.number

def factorial(number):
    number = int(number)
    if (number == 0) or (number == 1):
        return 1
    else:
        return number * factorial(number-1)