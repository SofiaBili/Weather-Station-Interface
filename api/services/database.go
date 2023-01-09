package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"

	"weather-station-api/interfaces"
)

// Type CustomTypeConverter and the next two functions solves issue: https://github.com/go-gorp/gorp/issues/142.
// The problem and the solution is described perfectly at: https://medium.com/@cikupin/storing-retrieving-postgresql-json-data-using-gorp-8eb6b095b890.
// TL;DR GORP doesn't read non-primitives and we need to create custom type converter driver for each json we want to parse.
type CustomTypeConverter struct{}

func (t CustomTypeConverter) ToDb(val interface{}) (interface{}, error) {
	switch t := val.(type) {
	case []interfaces.BME280Values:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	case []interfaces.MQ2Values:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	case []interfaces.MQ3Values:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	case []interfaces.MQ4Values:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	case []interfaces.MQ135Values:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	case []interfaces.RainSensorValues:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	case []interfaces.DustSensorValues:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	case []interfaces.SolarModuleValues:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return val, nil
}

func (t CustomTypeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	switch target.(type) {
	case *[]interfaces.BME280Values:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDB: Unable to convert BME280Values entry to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	case *[]interfaces.MQ2Values:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDB: Unable to convert MQ2Values entry to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	case *[]interfaces.MQ3Values:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDB: Unable to convert MQ3Values entry to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	case *[]interfaces.MQ4Values:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDB: Unable to convert MQ4Values entry to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	case *[]interfaces.MQ135Values:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDB: Unable to convert MQ135Values entry to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	case *[]interfaces.RainSensorValues:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDB: Unable to convert RainSensorValues entry to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	case *[]interfaces.DustSensorValues:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDB: Unable to convert DustSensorValues entry to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	case *[]interfaces.SolarModuleValues:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDB: Unable to convert SolarModuleValues entry to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	}
	return gorp.CustomScanner{}, false
}

// initDB is a function that initiates the DbMap object.
// DbMap is a Map that connects the Database with Golang and runs queries automatically with minimal code.
func initDB() *gorp.DbMap {
	db, err := sql.Open("mysql", "iakmastro:123@/talos_db")
	if err != nil {
		log.Fatalln("sql.Open failed", err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}, TypeConverter: CustomTypeConverter{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "weather-station-api", log.Lmicroseconds))
	dbmap.AddTableWithName(interfaces.BME280List{}, "bme280_list").SetKeys(true, "SensorID")
	dbmap.AddTableWithName(interfaces.BME280Values{}, "bme280_values").SetKeys(true, "ID")
	dbmap.AddTableWithName(interfaces.MQ2List{}, "mq2_list").SetKeys(true, "SensorID")
	dbmap.AddTableWithName(interfaces.MQ2Values{}, "mq2_values").SetKeys(true, "ID")
	dbmap.AddTableWithName(interfaces.MQ3List{}, "mq3_list").SetKeys(true, "SensorID")
	dbmap.AddTableWithName(interfaces.MQ3Values{}, "mq3_values").SetKeys(true, "ID")
	dbmap.AddTableWithName(interfaces.MQ4List{}, "mq4_list").SetKeys(true, "SensorID")
	dbmap.AddTableWithName(interfaces.MQ4Values{}, "mq4_values").SetKeys(true, "ID")
	dbmap.AddTableWithName(interfaces.MQ135List{}, "mq135_list").SetKeys(true, "SensorID")
	dbmap.AddTableWithName(interfaces.MQ135Values{}, "mq135_values").SetKeys(true, "ID")
	dbmap.AddTableWithName(interfaces.DustSensorList{}, "dust_sensor_list").SetKeys(true, "SensorID")
	dbmap.AddTableWithName(interfaces.DustSensorValues{}, "dust_sensor_values").SetKeys(true, "ID")
	dbmap.AddTableWithName(interfaces.RainSensorList{}, "rain_sensor_list").SetKeys(true, "SensorID")
	dbmap.AddTableWithName(interfaces.RainSensorValues{}, "rain_sensor_values").SetKeys(true, "ID")
	dbmap.AddTableWithName(interfaces.SolarModuleList{}, "solar_module_list").SetKeys(true, "SensorID")
	dbmap.AddTableWithName(interfaces.SolarModuleValues{}, "solar_module_values").SetKeys(true, "ID")
	dbmap.AddTableWithName(interfaces.WeatherStation{}, "weather_station").SetKeys(true, "StationID")
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalln("Create tables failed", err)
	}
	return dbmap
}

var DbMap = initDB()
