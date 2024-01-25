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


class Number_in_board:
    def __init__(self, number):
        self.number = number
        self.has_been_drawn = False

    def drawn(self):
        self.has_been_drawn = True

    def __str__(self):
        if self.has_been_drawn:
            return Bcolors.OKGREEN + str(self.number) + Bcolors.ENDC
        return Bcolors.WARNING + str(self.number) + Bcolors.ENDC

    def __repr__(self):
        if self.has_been_drawn:
            return Bcolors.OKGREEN + str(self.number) + Bcolors.ENDC
        return Bcolors.WARNING + str(self.number) + Bcolors.ENDC


class Board:
    def __init__(self):
        self.matrix = []

    def add_row_number(self, lines):
        number_in_row = [Number_in_board(int(numbers)) for numbers in lines[0].split()]
        self.matrix.append(number_in_row)
        lines.pop(0)

    def use_drawn_number(self, number):
        [[i.drawn() for i in row if i.number == number] for row in self.matrix]

    def total_undraw(self):
        result = []
        [[result.append(i.number) for i in row if i.has_been_drawn == False] for row in self.matrix]
        return sum(result)

    def has_won(self):
        for i in range(len(self.matrix)):
            has_won_row = True
            has_won_column = True
            for j in range(len(self.matrix[i])):
                if self.matrix[i][j].has_been_drawn == False:
                    has_won_row = False
                if self.matrix[j][i].has_been_drawn == False:
                    has_won_column = False

            if has_won_row:
                return True
            if has_won_column:
                return True

        return False

    def __str__(self):
        return str(self.matrix)

    def __repr__(self):
        return str(self.matrix)


class Bingo:
    boards = []

    def add_boards(self, lines):
        board = Board()
        for i in range(5):
            board.add_row_number(lines)
        self.boards.append(board)

    def use_drawn_number(self, number):
        [board.use_drawn_number(number) for board in self.boards]

    def board_has_won(self):
        won_board = [board for board in self.boards if board.has_won()]
        return len(won_board) > 0

    def calculate_result(self, number):
        won_board = [board for board in self.boards if board.has_won()]
        return won_board[0].total_undraw() * number

    def __str__(self):
        return str(self.boards)

    def __repr__(self):
        return str(self.boards)


def create_bingo(lines):
    bingo_game = Bingo()
    while len(lines) > 0:
        lines.pop(0)
        bingo_game.add_boards(lines)
    return bingo_game


def part_one():
    input_file = open('input.txt', 'r')
    lines = input_file.readlines()

    lines[0] = lines[0][:-1]
    drawn_numbers = [int(numbers) for numbers in lines[0].split(',')]

    lines.pop(0)
    bingo = create_bingo(lines)
    last_drawn_number = 0
    while len(drawn_numbers) > 0 and bingo.board_has_won() == False:
        last_drawn_number = drawn_numbers[0]
        bingo.use_drawn_number(last_drawn_number)
        drawn_numbers.pop(0)

    print("the result is " + str(bingo.calculate_result(last_drawn_number)))


part_one()
