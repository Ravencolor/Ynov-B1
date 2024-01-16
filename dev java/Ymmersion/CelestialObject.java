// CelestialObject.java

public class CelestialObject {
    private double x;
    private double y;
    private double z;
    private String name;

    public static final double KM_IN_ONE_AU = 150_000_000.0;

    public CelestialObject() {
        this.x = 0.0;
        this.y = 0.0;
        this.z = 0.0;
        this.name = "Soleil";
    }

    public CelestialObject(String name, double x, double y, double z) {
        this.name = name;
        this.x = x;
        this.y = y;
        this.z = z;
    }

    public double getX() {
        return x;
    }

    public void setX(double x) {
        this.x = x;
    }

    public double getY() {
        return y;
    }

    public void setY(double y) {
        this.y = y;
    }

    public double getZ() {
        return z;
    }

    public void setZ(double z) {
        this.z = z;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
    
    public static double getDistanceBetween(CelestialObject obj1, CelestialObject obj2) {
        double deltaX = obj1.getX() - obj2.getX();
        double deltaY = obj1.getY() - obj2.getY();
        double deltaZ = obj1.getZ() - obj2.getZ();
        return Math.sqrt(deltaX * deltaX + deltaY * deltaY + deltaZ * deltaZ);
    }

    public static double getDistanceBetweenInKm(CelestialObject obj1, CelestialObject obj2) {
        return getDistanceBetween(obj1, obj2) * KM_IN_ONE_AU;
    }

    @Override
    public String toString() {
        return String.format("%s is positioned at (%.3f, %.3f, %.3f)", name, x, y, z);
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) return true;
        if (obj == null || getClass() != obj.getClass()) return false;

        CelestialObject other = (CelestialObject) obj;

        return Double.compare(other.x, x) == 0 &&
               Double.compare(other.y, y) == 0 &&
               Double.compare(other.z, z) == 0 &&
               name.equals(other.name);
    }

    @Override
    public int hashCode() {
        int result = 17;
        long temp;

        temp = Double.doubleToLongBits(x);
        result = 31 * result + (int) (temp ^ (temp >>> 32));

        temp = Double.doubleToLongBits(y);
        result = 31 * result + (int) (temp ^ (temp >>> 32));

        temp = Double.doubleToLongBits(z);
        result = 31 * result + (int) (temp ^ (temp >>> 32));

        result = 31 * result + name.hashCode();

        return result;
    }
}
