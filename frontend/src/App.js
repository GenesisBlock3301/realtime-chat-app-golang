import logo from './logo.svg';
import './App.css';
import {useEffect, useState} from "react";

const socket = new WebSocket("ws://localhost:8080/ws")

function App() {
  const [message, setMessage] = useState([])
  const [inputValue, setInputValue] = useState('')

  useEffect(() => {
    socket.onopen = () => {
      console.log("Connected")
      setMessage('Connected')
    };

    socket.onmessage = (e) => {
      setMessage("Get message from server: " + e.data)
    };
    socket.onclose=()=>{
      console.log("closed")
      socket.close()
    }
  }, [])
  const handleChange = (e) =>{
    setInputValue(e.target.value)
  }
  const handleClick = (e)=>{
    e.preventDefault()
    socket.send(JSON.stringify({
      message: inputValue
    }))
  }
  return (
    <div className="App">
      <input type="text" value={inputValue} onChange={handleChange}/>
      <button onClick={handleClick}>Send</button>
      <pre>{message}</pre>
    </div>
  );
}

export default App;
