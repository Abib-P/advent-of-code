def part_two():
    input_file = open('input.txt', 'r')
    lines = input_file.readlines()
    aim = 0
    depth = 0
    horizontal = 0
    for line in lines:
        words = line.split()
        match words[0]:
            case "forward":
                horizontal += int(words[1])
                depth += aim * int(words[1])
            case "up":
                aim -= int(words[1])
            case "down":
                aim += int(words[1])
    print("result is " + str(depth * horizontal))


part_two()
