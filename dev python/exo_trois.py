class Ingredient:
    def __init__(self, name, state, unit, quantity):
        self.name = name
        self.state = state
        self.unit = unit
        self.quantity = quantity

    def equals(self, ingredient):
        return self.name == ingredient.name and self.state == ingredient.state


class CookableIngredient(Ingredient):
    def __init__(self, name, state, unit, quantity):
        super().__init__(name, state, unit, quantity)

    def cook(self):
        self.state = "cooked"


class CutableIngredient(Ingredient):
    def __init__(self, name, state, unit, quantity):
        super().__init__(name, state, unit, quantity)

    def cut(self):
        self.state = "cut"


class Dish:
    def __init__(self, name, lst):
        self.name = name
        self.list = lst

    def equals(self, dish):
        for i in range(len(self.list)):
            if not self.list[i].equals(dish.list[i]):
                return False
        return True

# Ingrédients
choucroute = CookableIngredient("choucroute", "raw", "g", "500")
lard = CookableIngredient("lard", "whole cooked", "g", "150")
saucisse = CookableIngredient("saucisse", "whole cooked", "piece", "2")
carotte = CutableIngredient("Carotte", "whole", "piece", "2")

# Plats
plat1 = Dish("choucroute", [choucroute, lard, saucisse])

choucroute.cook()
print(choucroute.state)
carotte.cut()
print(carotte.state)
