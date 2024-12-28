import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import MariaDB  from './mariadb/Mariadb';
import Home from './home/Home';
import './App.css';
function App() {
  return (
    <Router>
      <div>
        <nav className='App-header'>
              <Link to="/" className='App-link'>Home</Link>
              <Link to="/mariadb" className='App-link'>MariaDB</Link>
        </nav>
        <Routes>
          <Route path="/" element={<Home/>} />
          <Route path="/mariadb" element={<MariaDB/>} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;