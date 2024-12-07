from aocd import get_data


def get_input(for_example, day):
    if for_example:
        with open("/tmp/ex") as file:
            return file.read().splitlines()
    else:
        return get_data(year=2024, day=day).splitlines()