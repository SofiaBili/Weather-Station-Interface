package interfaces

type SolarModuleValues struct {
	ID int64 `db:"id" json:"id"`
	SensorID int64 `db:"sensor_id" json:"sensorId"`
	Moment string `db:"moment" json:"moment"`
	SolarCharge int64 `db:"solar_charge" json:"solarCharge"`
	SolarDone int64 `db:"solar_done" json:"solarDone"`
	SolarWarning int64 `db:"solar_warning" json:"solarWarning"`
}

type SolarModuleList struct {
	SensorID int64 `db:"sensor_id" json:"sensorId"`
}