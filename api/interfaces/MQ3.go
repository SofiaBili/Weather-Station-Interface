package interfaces

type MQ3Values struct {
	ID int64 `db:"id" json:"id"`
	SensorID int64 `db:"sensor_id" json:"sensorId"`
	Moment string `db:"moment" json:"moment"`
	Alcohol float64 `db:"alcohol" json:"alcohol"`
	Benzene float64 `db:"benzene" json:"benzene"`
	Hexane float64 `db:"hexane" json:"hexane"`
}

type MQ3List struct {
	SensorID int64 `db:"sensor_id" json:"sensorId"`
}