# Data for Rocket Calculations
To be able to do some ballpark calculations to check if the rocket equations work okay, I've provided some CSV files with data about various fuel tanks and rocket engines.

None of this is entirely accurate. I have taken some liberties. The column for engine thrust e.g. mixes thrust in vaccuum and at sea level (1 athmosphere pressure). Engines typically used for the first stage on earth have their values given for sea level. While smaller engines used for second stage have thrust and specific impulse given for vacuum.

# Format Description and Units
I am using SI unites for the most part but mass is given in metric tonnes and thrust is given in Kilo Newton. Keep that in mind when reading this file.
    
Rocket engines file has format:

    Name,  Company,  Mass (tonne),  Thrust (kN),  Specific Impulse (s),  Fuel consumption,  Burn time (s)
    
Most of these values should be self explanatory. "Fuel consumption" is something from Kerbal I am not sure about. Burn time, is how long the rocket engines are turned on in total for a given stage.