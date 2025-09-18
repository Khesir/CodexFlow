import axios from "axios";

const api = axios.create({
  baseURL: "/api", // this works in both dev & prod (proxy handles it in dev)
});

export default api;