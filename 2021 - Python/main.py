sonar_sweep = __import__('1_Sonar_Sweep')
dive = __import__('2_Dive')
binary_diagnostic = __import__('3_Binary_Diagnostic')


def main():
    sonar_sweep.part_one("./1_Sonar_Sweep/input.txt")
    sonar_sweep.part_two("./1_Sonar_Sweep/input.txt")
    dive.part_one()
    dive.part_two()
    binary_diagnostic.part_one()
    binary_diagnostic.part_two()


if __name__ == "__main__":
    main()
