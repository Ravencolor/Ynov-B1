public class DeadCharacterException extends Exception {

    private Character character;

    public DeadCharacterException(Character character) {
        this.character = character;
    }

    public String getMessage() {
        return "The " + character.getClass().getSimpleName().toLowerCase() + " " + character.getName() + " is dead.";
    }
}
