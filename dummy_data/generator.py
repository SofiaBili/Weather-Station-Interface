import pandas as pd
import numpy as np

from random import uniform, randrange
from time import sleep
from sys import argv
from dotenv import dotenv_values

from db import Database

BLUE = '\033[1;34m'
RED = '\033[0;31m'
GREEN = '\033[1;32m'
DEFAULT = '\033[0m'

SLEEP_INTERVAL = float(argv[1]) # Get the sleep interval from the system
NUM_OF_REPEATS = int(argv[2]) # Get how many times the loop is gonna function from the system

credentials = dotenv_values(".env")

if __name__ == '__main__':
  print(f"{BLUE}Trying to connect to the database... Please wait...{DEFAULT}")
  try:
    db = Database(credentials['HOST'], credentials['USER'], credentials['PASSWORD'], credentials['DATABASE'])

  except:
    print(f"{RED}Couldn't connect to the database. Exiting...{DEFAULT}")
    exit(-1)

  dataset = pd.read_csv('data.csv')
  print(f"{BLUE}Print the dataset:\n{GREEN}{dataset}")

  current_values = dict()
  print(f"{BLUE}\nGenerate the initial values")
  for entry in dataset.values:
    if entry[3] != 0:
      current_values[entry[0]] = uniform(entry[1], entry[2])
    else:
      current_values[entry[0]] = randrange(0, 2, 1)

  for i in range(NUM_OF_REPEATS):
    print(f"{GREEN}{current_values}{DEFAULT}")
    db.insert_values(current_values)
    for entry in dataset.values:
      if entry[3] != 0:
        current_values[entry[0]] += uniform(-entry[3], entry[3])
        if current_values[entry[0]] < entry[1]:
          current_values[entry[0]] = entry[1]
        if current_values[entry[0]] > entry[2]:
          current_values[entry[0]] = entry[2]
      else:
        current_values[entry[0]] = randrange(0, 2, 1)

    sleep(SLEEP_INTERVAL * 60)
