// Home.js
import React from "react";
import './Home.css';
import logo from './../logo.svg';

function Home() {
    return (
        <div className="App">   
            <h1>Welcome to Home Page</h1>
            <img src={logo} alt="Logo" className="App-logo" />
        </div>
    );
}

export default Home;
