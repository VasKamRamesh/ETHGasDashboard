// src/App.jsx
import React from "react";
import DiamondViewer from "./components/EthModel.jsx";

function App() {
  return (
    <div className="App">
      <header className="header">
        <h1 className="title">ETH Gas Tracker</h1>
        <div className="viewer-container">
          <DiamondViewer />
        </div>
      </header>
    </div>
  );
}

export default App;
