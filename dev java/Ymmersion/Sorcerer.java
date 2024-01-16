public class Sorcerer extends Character implements Healer {
    private final int healCapacity;

    public Sorcerer(String name, int maxHealth, int healCapacity, Weapon weapon) {
        super(name, maxHealth, weapon);
        this.healCapacity = healCapacity;
    }

    public int getCurrentHealth() {
        return this.currentHealth;
    }

    @Override
    public void attack(Character character) throws DeadCharacterException {
        if (currentHealth <= 0 ) throw new DeadCharacterException(this);
        character.takeDamage(this.getWeapon().getDamage());
    }

    @Override
    public void takeDamage(int damage) throws DeadCharacterException {
        if (currentHealth <= 0 ) throw new DeadCharacterException(this);
        this.currentHealth = Math.max(this.currentHealth-damage, 0);
    }

    @Override
    public int getHealCapacity() {
        return healCapacity;
    }

    @Override
    public void heal(Character character) throws DeadCharacterException {
        if (currentHealth <= 0 ) throw new DeadCharacterException(this);
        int currentHealth = character.getCurrentHealth();
        int newHealth = currentHealth + this.healCapacity;
        if (newHealth > character.getMaxHealth()) {
            newHealth = character.getMaxHealth();
        }
        character.currentHealth = newHealth;
    }

    @Override
    public String toString() {
        String description;
        if (this.currentHealth == 0) {
            description = this.getName() + " is a dead sorcerer. So bad, it could heal " + this.healCapacity + " HP.";
        } else {
            description = this.getName() + " is a sorcerer with " + this.currentHealth + " HP. It can heal " + this.healCapacity + " HP.";
        }

        return description + " He has the weapon " + this.getWeapon().toString() + ".";
    }
}