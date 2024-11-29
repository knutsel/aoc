def read_file(filename):
    with open('input/'+filename) as file:
        lines = file.read().splitlines()
    return lines
