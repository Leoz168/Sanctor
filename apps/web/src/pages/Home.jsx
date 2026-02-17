import { useState, useEffect } from 'react'
import '@styles/Home.css'
import { healthCheck } from '@services/api'

function Home() {
  const [apiStatus, setApiStatus] = useState('Checking...');
  const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

  useEffect(() => {
    healthCheck()
      .then(data => setApiStatus(data.message))
      .catch(() => setApiStatus('API not available'))
  }, [])

  return (
    <div className="App">
      <header className="App-header">
        <h1>Sanctor</h1>
        <p>Go Backend + React Frontend</p>
        <div className="status">
          <strong>Backend Status:</strong> {apiStatus}
        </div>
      </header>
    </div>
  );
}

export default Home
