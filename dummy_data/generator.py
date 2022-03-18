import pandas as pd
import numpy as np

from random import uniform, randrange
from time import sleep

BLUE = '\033[1;34m'
RED = '\033[0;31m'
GREEN = '\033[1;32m'
DEFAULT = '\033[0m'

if __name__ == '__main__':
  dataset = pd.read_csv('data.csv')
  print(f"{BLUE}Print the dataset:\n{GREEN}{dataset}")

  current_values = dict()
  print(f"{BLUE}\nGenerate the initial values")
  for entry in dataset.values:
    if entry[3] != 0:
      current_values[entry[0]] = uniform(entry[1], entry[2])
    else:
      current_values[entry[0]] = randrange(0, 2, 1)

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
