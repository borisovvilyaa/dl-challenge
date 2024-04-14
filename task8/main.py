import argparse

def maximize_income(N, C, gains, prices):
    # Creating a list of tuples (gain, price, index) and sorting it by decreasing gain
    laptops = sorted([(gains[i], prices[i], i) for i in range(len(gains))], reverse=True)
    
    capital = C  # Initial capital
    bought_laptops = []  # List of bought laptops
    
    for gain, price, index in laptops:
        if capital >= price and len(bought_laptops) < N:
            capital -= price
            capital += gain
            bought_laptops.append(index)
    
    return capital

def main():
    parser = argparse.ArgumentParser(description="Calculate capital at the end of the summer based on given parameters.")
    parser.add_argument("N", type=int, help="Number of laptops that can be bought")
    parser.add_argument("C", type=int, help="Initial capital")
    parser.add_argument("gains", type=str, help="Gain from each laptop")
    parser.add_argument("prices", type=str, help="Price of each laptop")
    args = parser.parse_args()

    # Parse gains and prices strings
    gains = [int(x) for x in args.gains.split(",")]
    prices = [int(x) for x in args.prices.split(",")]

    # Calculate result
    result = maximize_income(args.N, args.C, gains, prices)
    print("Capital at the end of the summer:", result)

if __name__ == "__main__":
    main()
