import axios from "axios";

const url = "http://localhost:3000";

const getStations = async () => {
  const { data: stations } = await axios.get(url);
  return stations;
};

const getStation = async (id) => {
  const { data: station } = await axios.get(`${url}/${id}`);
  return station;
};

export default {
  getStations,
  getStation,
};
