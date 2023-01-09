package interfaces

type MQ2Values struct {
	ID int64 `db:"id" json:"id"`
	SensorID int64 `db:"sensor_id" json:"sensorId"`
	Moment string `db:"moment" json:"moment"`
	H2 float64 `db:"h2" json:"h2"`
	LPG float64 `db:"lpg" json:"lpg"`
	Propane float64 `db:"propane" json:"propane"`
}

type MQ2List struct {
	SensorID int64 `db:"sensor_id" json:"sensorId"`
}