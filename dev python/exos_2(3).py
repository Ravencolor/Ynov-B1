class Dispositif:
    def __init__(self, marque, modele):
        self.marque = marque
        self.modele = modele

class OrdinateurPortable(Dispositif):
    def __init__(self, marque, modele, taille_ecran):
        super().__init__(marque, modele)
        self.taille_ecran = taille_ecran

class OrdinateurFixe(Dispositif):
    def __init__(self, marque, modele, configuration):
        super().__init__(marque, modele)
        self.configuration = configuration

class ConsoleJeu(Dispositif):
    def __init__(self, marque, modele, type_console):
        super().__init__(marque, modele)
        self.type_console = type_console

class Switch(ConsoleJeu):
    def __init__(self, marque, modele):
        super().__init__(marque, modele, "Switch")

class Xbox(ConsoleJeu):
    def __init__(self, marque, modele):
        super().__init__(marque, modele, "Xbox")

class Playstation(ConsoleJeu):
    def __init__(self, marque, modele):
        super().__init__(marque, modele, "Playstation")

class Vehicule:
    def __init__(self, marque, modele):
        self.marque = marque
        self.modele = modele

class Voiture(Vehicule):
    def __init__(self, marque, modele, nb_portes):
        super().__init__(marque, modele)
        self.nb_portes = nb_portes

class Camion(Vehicule):
    def __init__(self, marque, modele, capacite_charge):
        super().__init__(marque, modele)
        self.capacite_charge = capacite_charge

class Velo(Vehicule):
    def __init__(self, marque, modele, type_velo):
        super().__init__(marque, modele)
        self.type_velo = type_velo

class Trottinette(Vehicule):
    def __init__(self, marque, modele):
        super().__init__(marque, modele)
