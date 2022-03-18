# Dummy Data Generator

This is a temporary script used to test the interface without having any data from real sensors.
The script is implemented with Python and it uses ``random``, ``numpy`` and ``pandas``.

## How to run

To run this script, it needs to have ``numpy`` and ``pandas``. To install them, run this command:

```bash
pip3 install -r requirements.txt
```

This read from the ``requirements.txt`` file the modules used for the script.

Then just run from the CLI:

```bash
python3 generator.py
```

## Explanation

### CSV file

The CSV file contains the values to be used from the script to generate the data that will be used. Each line represents a field and each comma represents a value from the key. The fields are:

* **sensor**: which sensor is the data coming from
* **start**: which is the minimum value the sensor can read to look realistic
* **end**: which is the maximum value the sensor can read to look realistic
* **step**: how much it changes overtime to look realistic

Pandas is used to read the CSV file and parse it to a numpy array so it will be available to use the data in Python.

```py
dataset = pd.read_csv('data.csv')
print(f"{BLUE}Print the dataset:\n{GREEN}{dataset}")
```

### Initial Values

To make a dummy data generator that looks realistic, first it is needed to actually initialize a value in each sensor as a starting point.
A dictionary is made to keep the current values of each sensor in a format that functions similar to a JSON, same as the format that the API will read from.

```py
current_values = dict()
```

Next, it reads the data from the dataset. If the fourth field is empty, that means that it stores only two values and it picks randomly between those two values. Otherwise, ``random.uniform`` is needed to pick a random real number in the range of the minimum and maximum that the sensor can read.

```py
for entry in dataset.values:
  if entry[3] != 0:
    current_values[entry[0]] = uniform(entry[1], entry[2])
  else:
    current_values[entry[0]] = randrange(0, 2, 1)
```

### Calculate next value

Final step for the generator is to calculate the next value of each sensor every 3 seconds. The logic behind this is similar to the previous problem. As far as the sensors that read only two values, it still picks randomly between those two values. Otherwise, the third entry is used to find a random number to add to the current value. If the result is smaller than the minimum value then it picks the minimum value. Similarly, if the result is bigger than the maximum value then it picks the maximum value. In any other case, it keeps the result.

```py
while True:
  print(f"{GREEN}{current_values}{DEFAULT}")
  for entry in dataset.values:
    if entry[3] != 0:
      current_values[entry[0]] += uniform(-entry[3], entry[3])
      if current_values[entry[0]] < entry[1]:
        current_values[entry[0]] = entry[1]
      if current_values[entry[0]] > entry[2]:
        current_values[entry[0]] = entry[2]
    else:
      current_values[entry[0]] = randrange(0, 2, 1)

  sleep(3)
```
