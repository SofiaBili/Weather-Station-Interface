import utils from "./utils.js";

const getData = async (id) => {
  const response = await fetch(`http://localhost:5000/${id}`);
  const station = await response.json();
  return station;
};

const createCharts = async (id) => {
  const station = await getData(id);

  let chartName = document.getElementById("temp").getContext("2d");
  const temperatures = station.bme280.map((data) => data.temperature);
  utils.gaugeChart(
    chartName,
    temperatures.slice(-1),
    0,
    [...temperatures],
    "Temperature",
    utils.tempGaugeColors
  );

  chartName = document.getElementById("hum").getContext("2d");
  const humidities = station.bme280.map((data) => data.humidity);
  utils.gaugeChart(
    chartName,
    humidities.slice(-1),
    0,
    [...humidities],
    "Humidity",
    utils.tempGaugeColors
  );

  chartName = document.getElementById("bar").getContext("2d");
  const barometricPressures = station.bme280.map(
    (data) => data.barometricPressure
  );
  utils.gaugeChart(
    chartName,
    barometricPressures.slice(-1),
    0,
    [...barometricPressures],
    "Barometric Pressure",
    utils.tempGaugeColors
  );

  chartName = document.getElementById("bme").getContext("2d");
  let dataset = utils.fillDataset(
    [temperatures, humidities, barometricPressures],
    ["Temperature", "Humidity", "Barometric Pressure"]
  );
  const bme280Timestamps = station.bme280.map((data) => data.moment);
  utils.dynamicLineChart(
    chartName,
    bme280Timestamps,
    dataset,
    "Temperature, Humidity and Barometric Pressure"
  );

  chartName = document.getElementById("dust");
  const dustDensities = station.dustSensor.map((data) => data.dustDensity);
  const dustTimestamps = station.dustSensor.map((data) => data.moment);
  utils.lineChart(
    chartName,
    dustTimestamps,
    dustDensities,
    "Dust",
    utils.dustColor
  );

  chartName = document.getElementById("rain");
  const rainSensorValues = station.rainSensor.map((data) => data.sensorValue);
  const rainSensorTimestamps = station.rainSensor.map((data) => data.moment);
  utils.lineChart(
    chartName,
    rainSensorTimestamps,
    rainSensorValues,
    "Dust",
    utils.rainColor
  );

  chartName = document.getElementById("mq2");
  const h2Values = station.mq2.map((data) => data.h2);
  const lpgValues = station.mq2.map((data) => data.lpg);
  const propaneValues = station.mq2.map((data) => data.propane);
  dataset = utils.fillDataset(
    [h2Values, lpgValues, propaneValues],
    ["H2", "LPG", "Propane"]
  );
  const mq2Timestamps = station.mq2.map((data) => data.moment);
  utils.dynamicLineChart(chartName, mq2Timestamps, dataset, "MQ2");

  chartName = document.getElementById("mq3");
  const alcoholValues = station.mq3.map((data) => data.alcohol);
  const benzeneValues = station.mq3.map((data) => data.benzene);
  const hexaneValues = station.mq3.map((data) => data.hexane);
  dataset = utils.fillDataset(
    [alcoholValues, benzeneValues, hexaneValues],
    ["Alcohol", "Benzene", "Hexane"]
  );
  const mq3Timestamps = station.mq3.map((data) => data.moment);
  utils.dynamicLineChart(chartName, mq3Timestamps, dataset, "MQ3");

  chartName = document.getElementById("mq4");
  const ch4Values = station.mq4.map((data) => data.ch4);
  const coValues = station.mq4.map((data) => data.co);
  const smokeValues = station.mq4.map((data) => data.smoke);
  dataset = utils.fillDataset(
    [ch4Values, coValues, smokeValues],
    ["CH4", "CO", "Smoke"]
  );
  const mq4Timestamps = station.mq4.map((data) => data.moment);
  utils.dynamicLineChart(chartName, mq4Timestamps, dataset, "MQ4");

  chartName = document.getElementById("mq135");
  const acetoneValues = station.mq135.map((data) => data.acetone);
  const co2Values = station.mq135.map((data) => data.co2);
  const nh4Values = station.mq135.map((data) => data.nh4);
  const toluenoValues = station.mq135.map((data) => data.tolueno);
  dataset = utils.fillDataset(
    [acetoneValues, co2Values, nh4Values, toluenoValues],
    ["Acetone", "CO2", "NH4", "Tolueno"]
  );
  const mq135Timestamps = station.mq135.map((data) => data.moment);
  utils.dynamicLineChart(chartName, mq135Timestamps, dataset, "MQ135");

  chartName = document.getElementById("aq");
  utils.lineChart(
    chartName,
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    "Air Quality",
    utils.airQColor
  );
};

createCharts(1);

export { createCharts };
