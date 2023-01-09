package interfaces

type RainSensorValues struct {
	ID int64 `db:"id" json:"id"`
	SensorID int64 `db:"sensor_id" json:"sensorId"`
	Moment string `db:"moment" json:"moment"`
	SensorValue float64 `db:"sensor_value" json:"sensorValue"`
}

type RainSensorList struct {
	SensorID int64 `db:"sensor_id" json:"sensorId"`
}