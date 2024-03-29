
class Lanternfish:
    def __init__(self):
        self.lanternfish = [0 for i in range(9)]

    def add_initial_lanternfish(self, number: int):
        self.lanternfish[number] += 1

    def play_day(self):
        temp = [0 for i in range(9)]
        temp[0] = self.lanternfish[1]
        temp[1] = self.lanternfish[2]
        temp[2] = self.lanternfish[3]
        temp[3] = self.lanternfish[4]
        temp[4] = self.lanternfish[5]
        temp[5] = self.lanternfish[6]
        temp[6] = self.lanternfish[7] + self.lanternfish[0]
        temp[7] = self.lanternfish[8]
        temp[8] = self.lanternfish[0]
        self.lanternfish = temp

    def get_number_of_fish(self):
        result = 0
        for i in range(len(self.lanternfish)):
            result += self.lanternfish[i]
        return result

    def __str__(self):
        return str(self.lanternfish)

    def __repr__(self):
        return str(self.lanternfish)


def part_one(file_name: str):
    input_file = open(file_name, 'r')
    file_lines = input_file.readlines()

    lanternfish = Lanternfish()
    [lanternfish.add_initial_lanternfish(int(number)) for number in file_lines[0].split(",")]
    for i in range(80):
        lanternfish.play_day()
    print("the result is " + str(lanternfish.get_number_of_fish()))
