package services

import "weather-station-api/interfaces"

// This is the Data Access Object for the Stations.
// It is the lowest programmatically logic for the Stations.

func GetStations() ([]interfaces.WeatherStation, error) {
	var stations []interfaces.WeatherStation
	_, err := DbMap.Select(&stations, "SELECT * FROM weather_station")
	return stations, err
}

type StationDetails struct {
	StationID int64 `db:"station_id" json:"stationId"`
	Latitude string `db:"latitude" json:"latitude"`
	Longitude string `db:"longitude" json:"longitude"`
	Bme280 []interfaces.BME280Values `db:"bme280" json:"bme280"`
	MQ2 []interfaces.MQ2Values `db:"mq2" json:"mq2"`
	MQ3 []interfaces.MQ3Values `db:"mq3" json:"mq3"`
	MQ4 []interfaces.MQ4Values `db:"mq4" json:"mq4"`
	MQ135 []interfaces.MQ135Values `db:"mq135" json:"mq135"`
	RainSensor []interfaces.RainSensorValues `db:"rain_sensor" json:"rainSensor"`
	DustSensor []interfaces.DustSensorValues `db:"dust_sensor" json:"dustSensor"`
	SolarModule []interfaces.SolarModuleValues `db:"solar_module" json:"solarModule"`
}

func GetStationDetails(id string) (StationDetails, error) {
	var station StationDetails
	err := DbMap.SelectOne(
		&station,
		`SELECT
			station_id,
			latitude,
			longitude,
			JSON_ARRAYAGG(DISTINCT JSON_OBJECT(
				'id', bme280_values.id,
				'sensorId', bme280_id,
				'moment', bme280_values.moment,
				'temperature', temperature,
				'humidity', humidity,
				'barometricPressure', barometric_pressure
			)) bme280,
			JSON_ARRAYAGG(DISTINCT JSON_OBJECT(
				'id', mq2_values.id,
				'sensorId', mq2_id,
				'moment', mq2_values.moment,
				'h2', h2,
				'lpg', lpg,
				'propane', propane
			)) mq2,
			JSON_ARRAYAGG(DISTINCT JSON_OBJECT(
				'id', mq3_values.id,
				'sensorId', mq3_id,
				'moment', mq3_values.moment,
				'alcohol', alcohol,
				'benzene', benzene,
				'hexane', hexane
			)) mq3,
			JSON_ARRAYAGG(DISTINCT JSON_OBJECT(
				'id', mq4_values.id,
				'sensorId', mq4_id,
				'moment', mq4_values.moment,
				'ch4', ch4,
				'co', co,
				'smoke', smoke
			)) mq4,
			JSON_ARRAYAGG(DISTINCT JSON_OBJECT(
				'id', mq135_values.id,
				'sensorId', mq135_id,
				'moment', mq135_values.moment,
				'co2', co2,
				'tolueno', tolueno,
				'nh4', nh4,
				'acetone', acetone
			)) mq135,
			JSON_ARRAYAGG(DISTINCT JSON_OBJECT(
				'id', rain_sensor_values.id,
				'sensorId', rain_sensor_id,
				'moment', rain_sensor_values.moment,
				'sensorValue', rain_sensor_values.sensor_value
			)) rain_sensor,
			JSON_ARRAYAGG(DISTINCT JSON_OBJECT(
				'id', dust_sensor_values.id,
				'sensorId', dust_sensor_id,
				'moment', dust_sensor_values.moment,
				'sensorValue', dust_sensor_values.dust_density
			)) dust_sensor,
			JSON_ARRAYAGG(DISTINCT JSON_OBJECT(
				'id', solar_module_values.id,
				'sensorId', solar_module_id,
				'solarCharge', solar_module_values.solar_charge,
				'solarDone', solar_module_values.solar_done,
				'solarWarning', solar_module_values.solar_warning
			)) solar_module
		FROM weather_station
			INNER JOIN bme280_values ON bme280_id = bme280_values.sensor_id
			INNER JOIN mq2_values ON mq2_id = mq2_values.sensor_id
			INNER JOIN mq3_values ON mq2_id = mq3_values.sensor_id
			INNER JOIN mq4_values ON mq2_id = mq4_values.sensor_id
			INNER JOIN mq135_values ON mq2_id = mq135_values.sensor_id
			INNER JOIN rain_sensor_values ON rain_sensor_id = rain_sensor_values.sensor_id
			INNER JOIN dust_sensor_values ON dust_sensor_id = dust_sensor_values.sensor_id
			INNER JOIN solar_module_values ON solar_module_id = solar_module_values.sensor_id
		WHERE weather_station.station_id=?
		`,
		id,
	)
	return station, err
}