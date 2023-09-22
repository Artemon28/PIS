using System;
public abstract class TransportService
{
    public string Name { get; set; }

    public TransportService(string name)
    {
        Name = name;
    }

    public abstract double CostTransportation(double distance);
}

public abstract class TransportCompany
{
    public string Name { get; set; }

    public TransportCompany(string name)
    {
        Name = name;
    }

    public override string ToString()
    {
        return Name;
    }
    public abstract TransportService Create(string name, int param);
}

public class TaxiServices : TransportService
{
    public int Category { get; set; }

    public TaxiServices(string name, int cat) : base(name)
    {
        Category = cat;
    }

    public override double CostTransportation(double distance)
    {
        return distance * Category;
    }

    public override string ToString()
    {
        return $"Фирма {Name}, поездка категории {Category}";
    }
}

public class Shipping : TransportService
{
    public double Tariff { get; set; }

    public Shipping(string name, double taff) : base(name)
    {
        Tariff = taff;
    }

    public override double CostTransportation(double distance)
    {
        return distance * Tariff;
    }

    public override string ToString()
    {
        return $"Фирма {Name}, доставка по тарифу {Tariff}";
    }
}

public class DrunkDriverServices : TransportService
{
    public DrunkDriverServices(string name) : base(name)
    {
    }

    public override double CostTransportation(double distance)
    {
        double fixedCost = 20.0;
        return fixedCost * distance;
    }

    public override string ToString()
    {
        return $"Фирма {Name}, услуга - пьяный водитель";
    }
}

public class TaxiTransCom : TransportCompany
{
    public TaxiTransCom(string name) : base(name)
    {
    }

    public override TransportService Create(string name, int c)
    {
        return new TaxiServices(Name, c);
    }
}

public class ShipTransCom : TransportCompany
{
    public ShipTransCom(string name) : base(name)
    {
    }

    public override TransportService Create(string name, int t)
    {
        return new Shipping(Name, t);
    }
}

public class DrunkDriverTransCom : TransportCompany
{
    public DrunkDriverTransCom(string name) : base(name)
    {
    }

    public override TransportService Create(string name, int param)
    {
        return new DrunkDriverServices(Name);
    }
}

class Program
{
    static void Main()
    {
        TransportCompany trCom = new TaxiTransCom("Служба такси");
        TransportService compService = trCom.Create("Такси", 1);
        double dist = 15.5;
        Print(compService, dist);
        TransportCompany gCom = new ShipTransCom("Служба перевозок");
        compService = gCom.Create("Грузоперевозки", 2);
        double distg = 150.5;
        Print(compService, distg);
        TransportCompany ddCom = new DrunkDriverTransCom("Служба пьяного водителя");
        compService = ddCom.Create("Пьяный водитель", 0);
        double distdd = 5.0;
        Print(compService, distdd);
    }

    private static void Print(TransportService compTax, double distg)
    {
        Console.WriteLine($"Компания {compTax}, расстояние {distg}, стоимость: {compTax.CostTransportation(distg):C}");
    }
}


//Adding new function DrunkDriver wasn't hard, because
//all i need is
//Created a New Concrete Service Class
//Created a New Concrete Company Class:
//Updated the Main Program