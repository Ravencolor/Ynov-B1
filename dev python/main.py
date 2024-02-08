class CoffeeShop:

    def __init__(self,name,stock_quantity):
        self.name = name
        self.stock_quantity = stock_quantity
        self.state = None

    def get_state(self):
        if self.stock_quantity > 5:
            self.state = True
        else:
            self.state = False
        return self.state

    def restock(self):
        self.stock_quantity += 10
        print("Restocked\n")
        return self.stock_quantity

    def serve(self):
        if self.get_state() == True:
            print("Welcome to " + self.name + "!\n")
            number_of_coffee = int(input("How many coffees would you like?\n"))
            self.take_commande(number_of_coffee)
        else:
            print("Sorry, we are restocking. Please come back later.\n")
            print("Restocking...\n")
            print("Many hours later... \n")
            self.restock()
            self.serve()

    def take_commande(self, number_of_coffee):
        if self.stock_quantity >= number_of_coffee:
            print("Here are your " + str(number_of_coffee) + " coffees. Enjoy!\n")
            self.stock_quantity -= number_of_coffee
            print("Remaining stock: " + str(self.stock_quantity) + "\n")
        else:
            print("Sorry, we don't have enough coffee. Please come back later.\n")






coffee_shop = CoffeeShop("Coffee and Code", 100)
coffee_shop.serve()
