public class Monster extends Character {
    public Monster(String name, int maxHealth, Weapon weapon) {
        super(name, maxHealth, weapon);
    }
    
    @Override
    public void attack(Character character) throws DeadCharacterException {
        if (currentHealth <= 0 ) throw new DeadCharacterException(this);
        character.takeDamage(this.getWeapon().getDamage());
    }

    @Override
    public void takeDamage(int damage) throws DeadCharacterException {
        if (currentHealth <= 0 ) throw new DeadCharacterException(this);
        int newDamage = (int) Math.floor(damage * 0.8);
        this.currentHealth = Math.max(this.currentHealth - newDamage, 0);
    }


    @Override
    public String toString() {
        String description = (getCurrentHealth()<=0 ? "<name> is a monster and is dead" : "<name> is a monster with <currentHealth> HP")
        .replace("<name>", getName())
        .replace("<currentHealth>", Integer.toString(this.getCurrentHealth()));
        
        return description + " He has the weapon " + this.getWeapon().getName() + " which deals " + this.getWeapon().toString() + ".";
    }
}
