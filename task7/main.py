import argparse

def calculate_candies(pinyatas):
    """
    Calculate the maximum number of candies that can be obtained by breaking pinyatas.

    @param pinyatas: A list of integers representing the pinyatas.
    @return: The maximum number of candies that can be obtained.
    """
    max_candies = 0
    for i in range(len(pinyatas)):
        left_pinyata = pinyatas[i - 1] if i - 1 >= 0 else pinyatas[0]
        right_pinyata = pinyatas[i + 1] if i + 1 < len(pinyatas) else pinyatas[0]
        candies = left_pinyata * pinyatas[i] * right_pinyata
        if candies > max_candies:
            max_candies = candies
    return max_candies

def main():
    """
    Parse command-line arguments and calculate maximum candies.
    """
    parser = argparse.ArgumentParser(description='Calculate maximum candies')
    parser.add_argument('pinyatas', metavar='N', type=int, nargs='+',
                        help='a list of integers representing pinyatas')
    args = parser.parse_args()
    max_candies = calculate_candies(args.pinyatas)
    print(max_candies)

if __name__ == "__main__":
    main()
