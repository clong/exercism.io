

def is_leap_year(year):
    # Divisible by 4
    if year % 4 == 0:
        if (year % 100 == 0 and year % 400 == 0):
            return True
        if (year % 100 == 0):
            return False
        return True
    else:
        return False
