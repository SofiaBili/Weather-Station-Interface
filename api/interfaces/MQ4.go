package interfaces

type MQ4Values struct {
	ID int64 `db:"id" json:"id"`
	SensorID int64 `db:"sensor_id" json:"sensorId"`
	Moment string `db:"moment" json:"moment"`
	CH4 float64 `db:"ch4" json:"ch4"`
	CO float64 `db:"co" json:"co"`
	Smoke float64 `db:"smoke" json:"smoke"`
}

type MQ4List struct {
	SensorID int64 `db:"sensor_id" json:"sensorId"`
}