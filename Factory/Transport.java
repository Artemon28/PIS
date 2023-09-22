abstract class TransportService {
    protected String name;

    public TransportService(String name) {
        this.name = name;
    }

    public abstract double costTransportation(double distance);
}

abstract class TransportCompany {
    protected String name;

    public TransportCompany(String name) {
        this.name = name;
    }

    public abstract TransportService create(String name, int param);
}

class TaxiServices extends TransportService {
    private int category;

    public TaxiServices(String name, int category) {
        super(name);
        this.category = category;
    }

    public double costTransportation(double distance) {
        return distance * category;
    }

    public String toString() {
        return String.format("Фирма %s, поездка категории %d", name, category);
    }
}

class Shipping extends TransportService {
    private double tariff;

    public Shipping(String name, double tariff) {
        super(name);
        this.tariff = tariff;
    }

    public double costTransportation(double distance) {
        return distance * tariff;
    }

    public String toString() {
        return String.format("Фирма %s, доставка по тарифу %.2f", name, tariff);
    }
}

class TaxiTransCom extends TransportCompany {
    public TaxiTransCom(String name) {
        super(name);
    }

    public TransportService create(String name, int category) {
        return new TaxiServices(name, category);
    }
}

class ShipTransCom extends TransportCompany {
    public ShipTransCom(String name) {
        super(name);
    }

    public TransportService create(String name, int tariff) {
        return new Shipping(name, tariff);
    }
}

public class Main {
    public static void main(String[] args) {
        // Create a taxi transportation company
        TransportCompany trCom = new TaxiTransCom("Служба такси");

        // Create a taxi service
        TransportService compService = trCom.create("Такси", 1);

        // Set the distance
        double dist = 15.5;

        // Print service information
        printServiceInfo(compService, dist);

        // Create a shipping company
        TransportCompany gCom = new ShipTransCom("Служба перевозок");

        // Create a shipping service
        compService = gCom.create("Грузоперевозки", 2);

        // Set the distance
        double distg = 150.5;

        // Print service information
        printServiceInfo(compService, distg);
    }

    private static void printServiceInfo(TransportService compTax, double distg) {
        System.out.printf("Компания %s, расстояние %.2f, стоимость: %.2f%n",
                compTax.toString(), distg, compTax.costTransportation(distg));
    }
}
