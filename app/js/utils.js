const tempGaugeColors = ["#032580", "#0582CA", "#F85E00", "#E81717"];
const gaugeColors = ["#5F9324", "#FFB30F", "#F85E00", "#E81717"];
const colors = [
  "#fa4dc8",
  "#7E34D1",
  "#0582CA",
  "#490F60",
  "#5F9324",
  "#FFB30F",
  "#F85E00",
  "#E81717",
];
const dustColor = ["#5F9324", "#5F932480"];
const rainColor = ["#437F97", "#437F9780"];
const airQColor = ["#FFB30F", "#FFB30F80"];

const fillDataset = (dataArray, labels) => {
  let i = 0;
  return dataArray.map((data) => {
    return {
      label: labels[i],
      data: data,
      fill: false,
      backgroundColor: colors[i],
      borderColor: colors[i++],
    };
  });
};

const lineChart = (chartName, xValues, yValues, chartText, color) => {
  const chart = new Chart(chartName, {
    type: "line",
    data: {
      labels: xValues,
      datasets: [
        {
          fill: true,
          lineTension: 0,
          backgroundColor: color[1],
          borderColor: color[0],
          data: yValues,
        },
      ],
    },
    options: {
      legend: { display: false },
      title: {
        display: true,
        text: chartText,
      },
    },
  });

  chart.render();
};

const dynamicLineChart = (chartName, xValues, datasetData, chartText) => {
  const chart = new Chart(chartName, {
    type: "line",
    data: {
      labels: xValues,
      datasets: datasetData,
    },
    options: {
      legend: {
        position: "left",
        labels: {
          fontSize: 12,
          boxWidth: 1,
        },
      },
      title: {
        display: true,
        text: chartText,
      },
    },
  });

  chart.render();
};

const gaugeChart = (
  chartName,
  currentValue,
  currentMinValue,
  currentLimit,
  chartText,
  colors
) => {
  const chart = new Chart(chartName, {
    type: "gauge",
    data: {
      datasets: [
        {
          value: currentValue,
          minValue: currentMinValue,
          backgroundColor: colors,
          data: currentLimit,
        },
      ],
    },
    options: {
      needle: {
        radiusPercentage: 12,
        widthPercentage: 3.0,
        lengthPercentage: 80,
        color: "#011627",
      },
      valueLabel: {
        display: true,
        formatter: (value) => {
          return Math.round(value);
        },
        fontSize: 12,
        color: "rgba(255, 255, 255, 1)",
        backgroundColor: "#011627",
        borderRadius: 10,
        padding: {
          top: 1,
          bottom: 10,
        },
      },
      title: {
        display: true,
        text: chartText,
      },
    },
  });

  chart.render();
};

export default {
  tempGaugeColors,
  gaugeColors,
  colors,
  dustColor,
  rainColor,
  airQColor,
  fillDataset,
  lineChart,
  dynamicLineChart,
  gaugeChart,
};
