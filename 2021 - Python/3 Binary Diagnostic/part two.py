input_file = open('input.txt', 'r')
oxygen_generator_rating_lines = input_file.readlines()
CO2_scrubber_rating_lines = oxygen_generator_rating_lines

number_of_bit_by_line = len(oxygen_generator_rating_lines[0]) - 1

for i in range(number_of_bit_by_line):
    count_0 = 0
    count_1 = 0
    for line in oxygen_generator_rating_lines:
        match line[i]:
            case '0':
                count_0 += 1
            case '1':
                count_1 += 1
    if count_0 > count_1:
        new_lines = []
        for line in oxygen_generator_rating_lines:
            match line[i]:
                case '0':
                    new_lines.append(line)
        oxygen_generator_rating_lines = new_lines
    else:
        new_lines = []
        for line in oxygen_generator_rating_lines:
            match line[i]:
                case '1':
                    new_lines.append(line)
        oxygen_generator_rating_lines = new_lines
    if (len(oxygen_generator_rating_lines) == 1):
        break

for i in range(number_of_bit_by_line):
    count_0 = 0
    count_1 = 0
    for line in CO2_scrubber_rating_lines:
        match line[i]:
            case '0':
                count_0 += 1
            case '1':
                count_1 += 1
    if count_0 > count_1:
        new_lines = []
        for line in CO2_scrubber_rating_lines:
            match line[i]:
                case '1':
                    new_lines.append(line)
        CO2_scrubber_rating_lines = new_lines
    else:
        new_lines = []
        for line in CO2_scrubber_rating_lines:
            match line[i]:
                case '0':
                    new_lines.append(line)
        CO2_scrubber_rating_lines = new_lines
    if (len(CO2_scrubber_rating_lines) == 1):
        break

oxygen_generator_rating = oxygen_generator_rating_lines[0][::-1].replace('\n', '')
CO2_scrubber_rating = CO2_scrubber_rating_lines[0][::-1].replace('\n', '')

oxygen_generator_result = 0
CO2_scrubber_result = 0

for i in range(number_of_bit_by_line):
    oxygen_generator_result += int(oxygen_generator_rating[i]) * pow(2, i)
    CO2_scrubber_result += int(CO2_scrubber_rating[i]) * pow(2, i)

print("result is " + str(oxygen_generator_result * CO2_scrubber_result))
