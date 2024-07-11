def read_file(filename):
    with open(filename) as file:
        lines = file.read().splitlines()
    return lines
