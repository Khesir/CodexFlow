
import { useState } from 'react';
import './App.css'
import api from './api';

function App() {
  const [msg, setMsg] = useState("");

  const pingServer = async () => {
    try {
      const res = await api.get("/ping");
      setMsg(res.data.message);
      console.log(res.data.message)
    } catch (err) {
      console.error("Error:", err);
    }
  };

  return (
    <div>
      <h1>Scrum App</h1>
      <button onClick={pingServer}>Ping Server</button>
      {msg && <p>Server says: {msg}</p>}
    </div>

  )
}

export default App
