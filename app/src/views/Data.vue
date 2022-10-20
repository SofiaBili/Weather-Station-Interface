<template>
  <div class="main">
    <div class="content-area">
      <div class="container-fluid">
        <div class="row">
          <div class="col-sm-6" style="background-color: white">
            <h2>Weather Sensors</h2>
            <div class="row">
              <div class="col-sm-6" style="background-color: lightcyan">
                humidity
              </div>
              <div class="col-sm-6" style="background-color: lightgray">
                temperature
              </div>
              <div class="col-sm-6" style="background-color: lightgray">
                barometric pressure
              </div>
            </div>
            <div class="row">
              <div class="col-sm-12"><Chart
                :chartData="tempConfig"
                :chartOptions="tempOptions"
              /></div>
            </div>
            <div class="row">
              <div class="col-sm-6"><Chart
                :chartData="dustConfig"
                :chartOptions="dustOptions"
              /></div>
              <div class="col-sm-6"><Chart
                :chartData="rainConfig"
                :chartOptions="rainOptions"
              /></div>
            </div>
          </div>
          <div class="col-sm-6" style="background-color: white">
            <h2>Air Quality Sensors</h2>
            <div class="row">
              <div class="col-sm-6"><Chart
                :chartData="mq2Config"
                :chartOptions="mq2Options"
              /></div>
              <div class="col-sm-6"><Chart
                :chartData="mq3Config"
                :chartOptions="mq3Options"
              /></div>
            </div>
            <div class="row">
              <div class="col-sm-6"><Chart
                :chartData="mq4Config"
                :chartOptions="mq4Options"
              /></div>
              <div class="col-sm-6"><Chart
                :chartData="mq135Config"
                :chartOptions="mq135Options"
              /></div>
            </div>

            <div class="row">
              <div class="col-sm-12"><Chart
                :chartData="airQualityConfig"
                :chartOptions="airQualityOptions"
              /></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Chart from "@/components/Chart";

export default {
  name: "Data",
  components: { Chart },
  setup() {
    const backgroundColor = "rgba(0,0,255,1.0)";
    const borderColor = "rgba(0,0,255,0.1)";

    const tempConfig = {
      labels: [50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150],
      datasets: [
        {
          label: 'Temperature, humidity and barometric pressure',
          backgroundColor: backgroundColor,
          borderColor: borderColor,
          data: [7, 8, 8, 9, 9, 9, 10, 11, 14 ,14, 15],
        },
      ],
    }

    const tempOptions = {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        title: {
          display: true,
          text: 'Temperature, humidity and barometric pressure'
        }
      }
    }

    const dustConfig = {
      labels: [1, 2, 7, 8, 9, 10],
      datasets: [
        {
          label: 'Dust',
          backgroundColor: backgroundColor,
          borderColor: borderColor,
          data: [7, 8, 8, 9, 9, 9],
        },
      ],
    }

    const dustOptions = {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        title: {
          display: true,
          text: 'Dust'
        }
      }
    }

    const rainConfig = {
      labels: [1, 2, 7, 8, 9, 10],
      datasets: [
        {
          label: 'Rain',
          backgroundColor: backgroundColor,
          borderColor: borderColor,
          data: [7, 8, 8, 9, 9, 9],
        },
      ],
    }

    const rainOptions = {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        title: {
          display: true,
          text: 'Rain'
        }
      }
    }

    const mq2Config = {
      labels: [1, 2, 7, 8, 9, 10],
      datasets: [
        {
          label: 'MQ2',
          backgroundColor: backgroundColor,
          borderColor: borderColor,
          data: [7, 8, 8, 9, 9, 9],
        },
      ],
    }

    const mq2Options = {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        title: {
          display: true,
          text: 'MQ2'
        }
      }
    }

    const mq3Config = {
      labels: [1, 2, 7, 8, 9, 10],
      datasets: [
        {
          label: 'MQ3',
          backgroundColor: backgroundColor,
          borderColor: borderColor,
          data: [7, 8, 8, 9, 9, 9],
        },
      ],
    }

    const mq3Options = {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        title: {
          display: true,
          text: 'MQ3'
        }
      }
    }

    const mq4Config = {
      labels: [1, 2, 7, 8, 9, 10],
      datasets: [
        {
          label: 'MQ4',
          backgroundColor: backgroundColor,
          borderColor: borderColor,
          data: [7, 8, 8, 9, 9, 9],
        },
      ],
    }

    const mq4Options = {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        title: {
          display: true,
          text: 'MQ4'
        }
      }
    }

    const mq135Config = {
      labels: [1, 2, 7, 8, 9, 10],
      datasets: [
        {
          label: 'MQ135',
          backgroundColor: backgroundColor,
          borderColor: borderColor,
          data: [7, 8, 8, 9, 9, 9],
        },
      ],
    }

    const mq135Options = {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        title: {
          display: true,
          text: 'MQ135'
        }
      }
    }

    const airQualityConfig = {
      labels: [1, 2, 7, 8, 9, 10],
      datasets: [
        {
          label: 'Air Quality',
          backgroundColor: backgroundColor,
          borderColor: borderColor,
          data: [7, 8, 8, 9, 9, 9],
        },
      ],
    }

    const airQualityOptions = {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        title: {
          display: true,
          text: 'Air Quality'
        }
      }
    }

    return {
      tempConfig,
      tempOptions,
      dustConfig,
      dustOptions,
      rainConfig,
      rainOptions,
      mq2Config,
      mq2Options,
      mq3Config,
      mq3Options,
      mq4Config,
      mq4Options,
      mq135Config,
      mq135Options,
      airQualityConfig,
      airQualityOptions
    }
  }
};
</script>
