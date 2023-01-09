package interfaces

type DustSensorValues struct {
	ID int64 `db:"id" json:"id"`
	SensorID int64 `db:"sensor_id" json:"sensorId"`
	Moment string `db:"moment" json:"moment"`
	DustDensity float64 `db:"dust_density" json:"dustDensity"`
}

type DustSensorList struct {
	SensorID int64 `db:"sensor_id" json:"sensorId"`
}