import copy
from functools import reduce

class Point:
    def __init__(self, x: int, y: int):
        if (x > Map.max):
            Map.max = x + 2
        if (y > Map.max):
            Map.max = y + 2
        self.x = x
        self.y = y

    @classmethod
    def from_str(cls, coordinates_str: str):
        coordinates = coordinates_str.split(',')
        return cls(int(coordinates[0]), int(coordinates[1]))

    def __str__(self):
        return str(self.x) + "," + str(self.y)

    def __repr__(self):
        return str(self.x) + "," + str(self.y)


class Line:
    def __init__(self, start_point: Point, end_point: Point):
        self.start = start_point
        self.end = end_point

    @classmethod
    def from_str(cls, line_str: str):
        points = line_str.split(" -> ")
        start_point = Point.from_str(points[0])
        end_point = Point.from_str(points[1])
        return cls(start_point, end_point)

    def __str__(self):
        return str(self.start) + " -> " + str(self.end)

    def __repr__(self):
        return str(self.start) + " -> " + str(self.end)


class Map:
    max: int = 1

    def __init__(self):
        self.map = [[0] * Map.max for i in range(Map.max)]

    def add_line(self, line: Line):
        line_to_add = copy.copy(line)

        while line_to_add.start.x < line_to_add.end.x and line_to_add.start.y < line_to_add.end.y:
            self.map[line_to_add.start.x][line_to_add.start.y] += 1
            line_to_add.start.x += 1
            line_to_add.start.y += 1
        while line_to_add.start.x < line_to_add.end.x and line_to_add.start.y > line_to_add.end.y:
            self.map[line_to_add.start.x][line_to_add.start.y] += 1
            line_to_add.start.x += 1
            line_to_add.start.y -= 1
        while line_to_add.start.x > line_to_add.end.x and line_to_add.start.y < line_to_add.end.y:
            self.map[line_to_add.start.x][line_to_add.start.y] += 1
            line_to_add.start.x -= 1
            line_to_add.start.y += 1
        while line_to_add.start.x > line_to_add.end.x and line_to_add.start.y > line_to_add.end.y:
            self.map[line_to_add.start.x][line_to_add.start.y] += 1
            line_to_add.start.x -= 1
            line_to_add.start.y -= 1

        while line_to_add.start.x < line_to_add.end.x:
            self.map[line_to_add.start.x][line_to_add.start.y] += 1
            line_to_add.start.x += 1
        while line_to_add.start.x > line_to_add.end.x:
            self.map[line_to_add.start.x][line_to_add.start.y] += 1
            line_to_add.start.x -= 1
        while line_to_add.start.y < line_to_add.end.y:
            self.map[line_to_add.start.x][line_to_add.start.y] += 1
            line_to_add.start.y += 1
        while line_to_add.start.y > line_to_add.end.y:
            self.map[line_to_add.start.x][line_to_add.start.y] += 1
            line_to_add.start.y -= 1
        self.map[line_to_add.start.x][line_to_add.start.y] += 1


def parse_one_line(line: str):
    line = line[:-1]
    return Line.from_str(line)


def part_two(file_name: str):
    input_file = open(file_name, 'r')
    file_lines = input_file.readlines()

    lines = [parse_one_line(line) for line in file_lines]

    maps = Map()

    [maps.add_line(line) for line in lines]

    points = [[i for i in row if i > 1] for row in maps.map]

    number_of_overlap = len(reduce(list.__add__, points))

    print("the result is " + str(number_of_overlap))
