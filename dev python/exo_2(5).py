import random
from abc import ABC, abstractmethod

class Character(ABC):
    def __init__(self, name, position, hitbox):
        self.hp = random.randint(50, 100)
        self.mp = random.randint(0, 50)
        self.name = name
        self.position = position
        self.hitbox = hitbox

    @abstractmethod
    def special_ability(self):
        pass

class Hero(Character):
    count = 0
    def __init__(self, name, position, hitbox):
        super().__init__(name, position, hitbox)
        self.level = 1
        Hero.count += 1

class Tank(Hero):
    def special_ability(self):
        return "Shield"

class Mage(Hero):
    def special_ability(self):
        return "Magic"

class Healer(Hero):
    def special_ability(self):
        return "Heal"

class Warrior(Hero):
    def __init__(self, name, position, hitbox):
        super().__init__(name, position, hitbox)
        self.mp = 0

    def special_ability(self):
        return "Attack"

class Monster(Character):
    count = 0
    def __init__(self, name, position, hitbox):
        super().__init__(name, position, hitbox)
        Monster.count += 1

class Minion(Monster):
    def special_ability(self):
        return "Swarm"

class Buldozer(Monster):
    def special_ability(self):
        return "Crush"
