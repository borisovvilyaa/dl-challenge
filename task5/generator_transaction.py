import csv
import random

def generate_transactions_csv(file_path, num_transactions):
    with open(file_path, 'w', newline='') as file:
        writer = csv.writer(file)
        writer.writerow(["tx_id", "tx_size", "tx_fee"])  # Writing header
        for i in range(num_transactions):
            tx_id = f"{i+1}"
            tx_size = random.randint(200, 500)  # Random size between 100 and 1000 bytes
            tx_fee = random.randint(100, 10000)  # Random fee between 1000 and 10000 satoshis
            writer.writerow([tx_id, tx_size, tx_fee])

if __name__ == "__main__":
    num_transactions = 100000
    generate_transactions_csv("transactions.csv", num_transactions)
    print(f"{num_transactions} transactions generated and saved to transactions.csv.")
