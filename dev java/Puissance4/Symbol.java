/* Rémi 

L'énumération Symbol définie les symboles par couleur de chaque jeton ou un symbole vide pour les cases vide:

JAUNE contient la valeur "X".

ROUGE contient la valeur "O".

VIDE contient la valeur ".".*/

public enum Symbol {

	JAUNE("X"),
	ROUGE("O"),
	VIDE(".");
	
	private final String name;
	
	Symbol(String name) {
		this.name = name;
	}
	
	public String toString() {
		return this.name;
	}
}
