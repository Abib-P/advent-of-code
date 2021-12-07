input_file = open('input.txt', 'r')
lines = input_file.readlines()

count = 0
last_number = None

for line in lines:
    if last_number == None:
        last_number = int(line)
        continue
    if last_number < int(line):
        count += 1
    last_number = int(line)

print("result is " + str(count))