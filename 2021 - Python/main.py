sonar_sweep = __import__('1_Sonar_Sweep')
dive = __import__('2_Dive')
binary_diagnostic = __import__('3_Binary_Diagnostic')
giant_squid = __import__('4_Giant_Squid')


def main():
    sonar_sweep.part_one("./1_Sonar_Sweep/input.txt")
    sonar_sweep.part_two("./1_Sonar_Sweep/input.txt")
    dive.part_one("./2_Dive/input.txt")
    dive.part_two("./2_Dive/input.txt")
    binary_diagnostic.part_one("./3_Binary_Diagnostic/input.txt")
    binary_diagnostic.part_two("./3_Binary_Diagnostic/input.txt")
    giant_squid.part_one("./4_Giant_Squid/input.txt")
    giant_squid.part_two("./4_Giant_Squid/input.txt")


if __name__ == "__main__":
    main()
