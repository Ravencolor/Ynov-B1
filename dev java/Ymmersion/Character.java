import java.util.ArrayList;
import java.util.List;

public abstract class Character {
    private final int maxHealth;
    protected int currentHealth;
    private final String name;
    private static List<Character> allCharacters = new ArrayList<Character>();
    private Weapon weapon;

    public Character(String name, int maxHealth, Weapon weapon) {
        this.name = name;
        this.maxHealth = maxHealth;
        this.currentHealth = maxHealth;
        this.weapon = weapon;
        addCharacter(this);
    }

    public int getMaxHealth() {
        return this.maxHealth;
    }

    public int getCurrentHealth() {
        return this.currentHealth;
    }

    public String getName() {
        return this.name;
    }

    public Weapon getWeapon() {
        return this.weapon;
    }

    @Override
    public String toString() {
        if (this.currentHealth == 0) {
            return this.name + " : KO";
        }
        return this.name + " : " + this.currentHealth + "/" + this.maxHealth;
    }

    public abstract void takeDamage(int damage) throws DeadCharacterException;

    public abstract void attack(Character character) throws DeadCharacterException;

    public void addCharacter(Character character) {
        allCharacters.add(character);
    }

    public static List<Character> getAllCharacters() {
        return allCharacters;
    }

    public static String printStatus()  {
        if(allCharacters.isEmpty()) return "------------------------------------------\n" + "Nobody's fighting right now !\n" + "------------------------------------------";
        StringBuilder charactersToString = new StringBuilder();
        for(Character c : allCharacters) charactersToString.append(" - <character.toString>".replace("<character.toString>", c.toString())).append("\n");
        return "------------------------------------------\n" + "Characters currently fighting :\n" + charactersToString + "------------------------------------------";
    }

    public static Character fight(Character character1, Character character2) {
        while (character1.getCurrentHealth() > 0 && character2.getCurrentHealth() > 0) {
            try {
                character1.attack(character2);
            } catch (DeadCharacterException e) {
                System.out.println(e.getMessage());
            }
            if (character2.getCurrentHealth() > 0) {
                try {
                    character2.attack(character1);
                } catch (DeadCharacterException e) {
                    System.out.println(e.getMessage());
                }
            }
        }
        if (character1.getCurrentHealth() == 0) {
            return character2;
        }
        return character1;
    }
}
