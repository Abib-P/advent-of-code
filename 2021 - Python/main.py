sonar_sweep = __import__('1_Sonar_Sweep')
dive = __import__('2_Dive')


def main():
    sonar_sweep.part_one("./1_Sonar_Sweep/input.txt")
    sonar_sweep.part_two("./1_Sonar_Sweep/input.txt")
    dive.part_one()
    dive.part_two()


if __name__ == "__main__":
    main()
