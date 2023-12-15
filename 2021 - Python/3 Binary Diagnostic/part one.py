input_file = open('input.txt', 'r')
lines = input_file.readlines()

number_of_bit_by_line = len(lines[0]) - 1
gamma_rate = []

for i in range(number_of_bit_by_line):
    count_0 = 0
    count_1 = 0
    for line in lines:
        match line[i]:
            case '0':
                count_0 += 1
            case '1':
                count_1 += 1
    if count_0 > count_1:
        gamma_rate.append(0)
    else:
        gamma_rate.append(1)

epsilon_rate = [i ^ 1 for i in gamma_rate]

gamma_rate = gamma_rate[::-1]
epsilon_rate = epsilon_rate[::-1]

gamma_result = 0
epsilon_result = 0

for i in range(number_of_bit_by_line):
    gamma_result += gamma_rate[i] * pow(2, i)
    epsilon_result += epsilon_rate[i] * pow(2, i)

print("result is " + str(gamma_result * epsilon_result))
