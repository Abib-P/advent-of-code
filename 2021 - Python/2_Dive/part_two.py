def part_two(file_name: str):
    input_file = open(file_name, 'r')
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
