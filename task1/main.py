class Order:
    def __init__(self, user_id, amount, price, side):
        self.user_id = user_id
        self.amount = amount
        self.price = price
        self.side = side  # True for buy, False for sell

class BalanceChange:
    def __init__(self, user_id, value, currency):
        self.user_id = user_id
        self.value = value
        self.currency = currency

class Balance:
    def __init__(self):
        self.balances = {}  # Dictionary to store balances

    def update_balance(self, user_id, value, currency):
        if user_id in self.balances:
            self.balances[user_id][currency] += value
        else:
            self.balances[user_id] = {currency: value}

    def print_balance(self):
        print("Current Balances:")
        for user_id, balances in self.balances.items():
            print("User", user_id, ":", balances)


class OrderBook:
    def __init__(self):
        self.buy_orders = []  # List of buy orders sorted by price (descending)
        self.sell_orders = []  # List of sell orders sorted by price (ascending)
        self.balance = Balance()  # Initialize balance tracker

    def place_order(self, order):
        if order.side:  # Buy order
            self.buy_orders.append(order)
            self.buy_orders.sort(key=lambda x: x.price, reverse=True)
        else:  # Sell order
            self.sell_orders.append(order)
            self.sell_orders.sort(key=lambda x: x.price)

        self.match_orders()

    def match_orders(self):
        while self.buy_orders and self.sell_orders:
            buy_order = self.buy_orders[0]
            sell_order = self.sell_orders[0]

            if buy_order.price >= sell_order.price:
                matched_amount = min(buy_order.amount, sell_order.amount)
                matched_price = sell_order.price

                # Calculate total price
                total_price = matched_amount * matched_price

                # Emit balance changes
                buy_change = BalanceChange(buy_order.user_id, -total_price, "USD")
                sell_change = BalanceChange(sell_order.user_id, matched_amount, "UAH")

                print("Matched order: User", buy_order.user_id, "buys", matched_amount, "UAH at", matched_price, "USD")
                print("Matched order: User", sell_order.user_id, "sells", matched_amount, "UAH at", matched_price, "USD")

                # Update balances
                self.balance.update_balance(buy_order.user_id, -total_price, "USD")
                self.balance.update_balance(sell_order.user_id, matched_amount, "UAH")

                # Remove matched amounts from orders
                buy_order.amount -= matched_amount
                sell_order.amount -= matched_amount

                # Remove filled orders
                if buy_order.amount == 0:
                    self.buy_orders.pop(0)
                if sell_order.amount == 0:
                    self.sell_orders.pop(0)
            else:
                break

order_book = OrderBook()

orders = [
    Order(1, 50, 25, True),
    Order(2, 50, 24, False),
    Order(3, 30, 26, True),
    Order(4, 20, 23, False),
    Order(5, 70, 27, True),
    Order(6, 60, 22, False),
    Order(7, 40, 28, True)
]

for order in orders:
    order_book.place_order(order)

order_book.balance.print_balance()
