package interfaces

type WeatherStation struct {
	StationID int64 `db:"station_id" json:"stationId"`
	Latitude string `db:"latitude" json:"latitude"`
	Longitude string `db:"longitude" json:"longitude"`
	Bme280ID int64 `db:"bme280_id" json:"bme280Id"`
	MQ2ID int64 `db:"mq2_id" json:"mq2Id"`
	MQ3ID int64 `db:"mq3_id" json:"mq3Id"`
	MQ4ID int64 `db:"mq4_id" json:"mq4Id"`
	MQ135ID int64 `db:"mq135_id" json:"mq135Id"`
	RainSensorID int64 `db:"rain_sensor_id" json:"rainSensorId"`
	DustSensorID int64 `db:"dust_sensor_id" json:"dustSensorId"`
	SolarModuleID int64 `db:"solar_module_id" json:"solarModuleId"`
}
