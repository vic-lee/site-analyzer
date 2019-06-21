import './App.css';
import Dashboard from './Dashboard';

const React = require('react');


const App: React.FC = () => {

  return (
    <div className="App">
      <div className="query-panel">
      </div>
      <div className="dashboard">
        <Dashboard />
      </div>
    </div>
  );
}

export default App;
