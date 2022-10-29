from mysql import connector

class Database:
  def __init__(self, host: str, user: str, passwd: str, database: str):
    super().__init__()
    self.__weather_station_database = connector.connect(
      host=host,
      user=user,
      passwd=passwd,
      database=database
    )
    self.__cursor = self.__weather_station_database.cursor()

  def insert_values(self, values: dict) -> None:
    self.__cursor.execute(f"INSERT INTO bme280_values (sensor_id, temperature, humidity, barometric_pressure) VALUES ({1}, {values['Temperature']}, {values['Humidity']}, {values['Pressure']})")
    self.__cursor.execute(f"INSERT INTO mq2_values (sensor_id, h2, lpg, propane) VALUES ({1}, {values['H2']}, {values['LGP']}, {values['Benzine']})")
    self.__cursor.execute(f"INSERT INTO mq3_values (sensor_id, alcohol, benzene, hexane) VALUES ({1}, {values['Alcohol']}, {values['Benzine']}, {values['Hexane']})")
    self.__cursor.execute(f"INSERT INTO mq4_values (sensor_id, ch4, co, smoke) VALUES ({1}, {values['CH4']}, {values['CO']}, {values['Smoke']})")
    self.__cursor.execute(f"INSERT INTO mq135_values (sensor_id, co2, tolueno, nh4, acetone) VALUES ({1}, {values['CO2']}, {values['Tolueno']}, {values['NH4']}, {values['Benzine']})")
    self.__cursor.execute(f"INSERT INTO dust_sensor_values (sensor_id, dust_density) VALUES ({1}, {values['Dust Sensor']})")
    self.__cursor.execute(f"INSERT INTO solar_module_values (sensor_id, solar_charge, solar_done, solar_warning) VALUES ({1}, {values['Solar Charge']}, {values['Solar Done']}, {values['Solar Warning']})")
    self.__cursor.execute(f"INSERT INTO rain_sensor_values (sensor_id, sensor_value) VALUES ({1}, {values['Rain Sensor']})")
    self.__weather_station_database.commit()
