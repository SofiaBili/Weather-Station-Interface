package interfaces

type MQ135Values struct {
	ID int64 `db:"id" json:"id"`
	SensorID int64 `db:"sensor_id" json:"sensorId"`
	Moment string `db:"moment" json:"moment"`
	CO2 float64 `db:"co2" json:"co2"`
	Tolueno float64 `db:"tolueno" json:"tolueno"`
	NH4 float64 `db:"nh4" json:"nh4"`
	Acetone float64 `db:"acetone" json:"acetone"`
}

type MQ135List struct {
	SensorID int64 `db:"sensor_id" json:"sensorId"`
}