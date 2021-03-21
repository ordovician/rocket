# Tank Data
To be able to do some ballpark calculations to check if the rocket equations work okay, I've provided some CSV files with data about various fuel tanks and rocket engines. In this package we got the fuel tank data. Not all of this is accurate.

The Falcon 9 and Falcon 1 fuel tank numbers required a bit of guesswork, since SpaceX does not give us those numbers. So for e.g. Falcon 9, first stage I would take the dry weight (without fuel) assumption and substract the weight of 9 Merlin 1D engines. While for Falcon 1, I would substract just the weight of 1 engine.

# Format Description and Units
While generally sticking to SI units the mass numbers are given in metric tonnes. So in 1000 Kg.

Fuel tanks file is formated like:

    Name, Company,  Total Mass (tonne),  Dry Mass (tonne)

The dry mass is the weight of the fuel tank without any propellant inside it. I have simplified this characterization by not specifying details such as the oxiders and fuel. The tanks will contain both, so propellant is just the mass of the oxidizer plus the fuel.
