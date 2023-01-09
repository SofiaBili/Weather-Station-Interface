package interfaces

type BME280Values struct {
	ID int64 `db:"id" json:"id"`
	SensorID int64 `db:"sensor_id" json:"sensorId"`
	Moment string `db:"moment" json:"moment"`
	Temperature float64 `db:"temperature" json:"temperature"`
	Humidity float64 `db:"humidity" json:"humidity"`
	BarometricPressure float64 `db:"barometric_pressure" json:"barometricPressure"`
}

type BME280List struct {
	SensorID int64 `db:"sensor_id" json:"sensorId"`
}