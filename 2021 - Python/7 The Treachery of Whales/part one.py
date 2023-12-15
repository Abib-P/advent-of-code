class Bcolors:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKCYAN = '\033[96m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'


class Crabs:
    def __init__(self):
        self.crabs = [-1 for i in range(1000)]

    def add_crabs(self, number: int):
        self.crabs[self.crabs.index(-1)] = number

    def __str__(self):
        return str(self.crabs)

    def __repr__(self):
        return str(self.crabs)


def add_crab(number_of_crabs: int = 0):
    number_of_crabs += 1
    print(bcolors.OKCYAN + "A crab has been added. There are now " + str(number_of_crabs) + " crabs." + bcolors.ENDC)


if __name__ == '__main__':
    input_file = open('input.txt', 'r')
    crabs_position = input_file.readlines()[0].split(",")
    number_of_crabs = len(crabs_position)
    print(bcolors.OKCYAN + "There are " + str(number_of_crabs) + " crabs." + bcolors.ENDC)

    crabs = Crabs()
    print(crabs)