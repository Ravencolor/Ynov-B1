public class Templar extends Character implements Healer, Tank {
    private final int healCapacity;
    private final int shield;

    public Templar(String name, int maxHealth, int healCapacity, int shield, Weapon weapon) {
        super(name, maxHealth, weapon);
        this.healCapacity = healCapacity;
        this.shield = shield;
    }

    public int getHealCapacity() {
        return healCapacity;
    }

    @Override
    public void attack(Character character) throws DeadCharacterException {
        if (currentHealth <= 0 ) throw new DeadCharacterException(this);
        character.takeDamage(this.getWeapon().getDamage());
    }

    @Override
    public void takeDamage(int damage) throws DeadCharacterException {
        if (currentHealth <= 0 ) throw new DeadCharacterException(this);
        int damageTaken = damage - this.shield;
        if (damageTaken < 0) {
            damageTaken = 0;
        }
        this.currentHealth = Math.max(this.currentHealth - damageTaken, 0);
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

    public int getShield() {
        return shield;
    }

    @Override
    public String toString() {
        String description;
        if (this.currentHealth == 0) {
            description = this.getName() + " has been beaten, even with its " + this.shield + " shield. So bad, it could heal " + this.healCapacity + " HP.";
        } else {
            description = this.getName() + " is a strong Templar with " + this.currentHealth + " HP. It can heal " + this.healCapacity + " HP and has a shield of " + this.shield + ".";
        }

        return description + " He has the weapon " + this.getWeapon().toString() + ".";
    }
}