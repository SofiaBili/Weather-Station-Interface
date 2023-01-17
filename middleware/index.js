import express from "express";
import cors from "cors";
import axios from "axios";

const app = express();
app.use(express.json());
app.use(cors());

const URL = "http://localhost:3000";

app.get("/", async (req, res) => {
  const { data: stations } = await axios.get(URL);
  res.send(stations);
});

app.get("/:id", async (req, res) => {
  const { id } = req.params;
  const { data: station } = await axios.get(`${URL}/${id}`);
  res.send(station);
});

app.listen(5000, "0.0.0.0", () => console.log("Middleware is ready"));
