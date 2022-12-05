import React from 'react';
import SideBar from "../components/SideBar"
import HomePage from './HomePage'
import DeviceList from './Devices/DeviceList'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'


function App() {
  return (
    <div className="flex">
        <Router>
          <SideBar />
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/devices" element={<DeviceList />} />
          </Routes>
        </Router>
    </div>
  );
}

export default App;
