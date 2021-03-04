# Data for Rocket Calculations
To be able to do some ballpark calculations to check if the rocket equations work okay, I've provided some CSV files with data about various fuel tanks and rocket engines.

None of this is entirely accurate. I have taken some liberties. The column for engine thrust e.g. mixes thrust in vaccuum and at sea level (1 athmosphere pressure). Engines typically used for the first stage on earth have their values given for sea level. While smaller engines used for second stage have thrust and specific impulse given for vacuum.

The Falcon 9 and Falcon 1 fuel tank numbers required a bit of guesswork, since SpaceX does not give us those numbers. So for e.g. Falcon 9, first stage I would take the dry weight (without fuel) assumption and substract the weight of 9 Merlin 1D engines. While for Falcon 1, I would substract just the weight of 1 engine.

# Format Description and Units
Mass for things like tanks and rocket engines are given in tonnes (metric tons, 1000 kg). Trust is gien in Kilo Newton. Apart from that I am generally sticking to SI units such as  Kg, metre, second etc.

Fuel tanks file is formated like:

    Name, Company,  Total Mass (tonne),  Dry Mass (tonne)

The dry mass is the weight of the fuel tank without any propellant inside it. I have simplified this characterization by not specifying details such as the oxiders and fuel. The tanks will contain both, so propellant is just the mass of the oxidizer plus the fuel.
    
Rocket engines file has format:

    Name,  Company,  Mass (tonne),  Thrust (kN),  Specific Impulse (s),  Fuel consumption,  Burn time (s)
    
Most of these values should be self explanatory. "Fuel consumption" is something from Kerbal I am not sure about. Burn time, is how long the rocket engines are turned on in total for a given stage.