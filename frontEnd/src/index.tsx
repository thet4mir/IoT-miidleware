import React from 'react';
import './static/global.css';
import ReactDOM from 'react-dom/client';
import App from './pages/App';
import UserService from './common/UserService'


const dashboard = () => {
  const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
  );

  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>
  );
}

UserService.initKeycloak(dashboard)
