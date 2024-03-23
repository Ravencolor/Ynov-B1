class Watch:
    def __init__(self, hour, minute):
        self.hour = hour
        self.minute = minute

    def clone(self):
        return Watch(self.hour, self.minute)

    def advance(self, minutes=1):
        self.minute += minutes
        while self.minute >= 60:
            self.hour += 1
            self.minute -= 60
            if self.hour >= 24:
                self.hour = 0

    def get_time_string(self):    
        return f"{self.hour}h{self.minute}"


class Person:
    def __init__(self, name, watch=None):
        self.name = name
        self.watch = watch

    def wear(self, watch):
        if self.watch is None:
            self.watch = watch

    def remove(self):
        self.watch = None

    def ask_time(self, person):
        if person.watch is not None:
            return person.watch.get_time_string()
        else:
            return("la personne n'a pas de montre")






w = Watch(10, 10)
print(f"Création d'une montre : {w.get_time_string()}")

w2 = w.clone()
print(f"Clonage de la montre : {w2.get_time_string()}")

w.advance(30)
print(f"Avancement de la montre : {w.get_time_string()}")

tintin = Person("Tintin")
print(f"Création de la personne : {tintin.name}")

dupond = Person("Dupond", w2)
print(f"Création de la personne : {dupond.name}")

tintin.wear(w)
print(f"{tintin.name} porte une montre : {tintin.watch.get_time_string()}")

tintin.remove()
print(f"{tintin.name} a enlevé sa montre")

time = tintin.ask_time(dupond)
print(f"{tintin.name} demande l'heure à {dupond.name} , il est {time}")


dupond.wear(w2)
print(f"{dupond.name} porte une montre : {dupond.watch.get_time_string()}")

dupond.remove()
print(f"{dupond.name} a enlevé sa montre")


time = dupond.ask_time(dupond)
print(f"{dupond.name} demande l'heure à {tintin.name} , {time}")

time = tintin.ask_time(dupond)
print(f"{tintin.name} demande l'heure à {dupond.name} , {time}")
