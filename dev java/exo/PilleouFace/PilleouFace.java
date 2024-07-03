import java.util.Scanner;

public class PilleouFace {
    public static void main(String[] args) {
        // Lancer la pièce au hasard
        int random = (int) (Math.random() * 2); // 0 pour pile, 1 pour face
        
        // Demander à l'utilisateur de saisir pile ou face
        Scanner scanner = new Scanner(System.in);
        System.out.print("Choisissez pile ou face: ");
        String choixUtilisateur = scanner.nextLine();
        
        // Vérifier le résultat et afficher le résultat
        String resultat = (random == 0) ? "pile" : "face";
        System.out.println("Le résultat est: " + resultat);
        
        // Vérifier si l'utilisateur a gagné ou perdu
        if (choixUtilisateur.equalsIgnoreCase(resultat)) {
            System.out.println("Vous avez gagné!");
        } else {
            System.out.println("Vous avez perdu!");
        }
        
        // Fermer le scanner
        scanner.close();
    }
}
