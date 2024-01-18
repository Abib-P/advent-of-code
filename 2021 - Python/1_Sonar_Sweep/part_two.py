def part_two(file_name: str):
    input_file = open(file_name, 'r')
    lines = input_file.readlines()
    count = 0
    three_last_number = []
    for line in lines:
        if len(three_last_number) < 3:
            three_last_number.append(int(line))
            continue
        last_sum = sum(three_last_number)
        three_last_number.pop(0)
        three_last_number.append(int(line))
        if last_sum < sum(three_last_number):
            count += 1
    print("result is " + str(count))
