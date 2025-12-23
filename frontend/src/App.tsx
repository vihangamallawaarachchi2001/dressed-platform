// src/App.tsx
import './App.css'
import { BrowserRouter as Router } from 'react-router-dom';
import Routes from './routes';

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-white">
        <Routes />
      </div>
    </Router>
  );
}

export default App;
