// src/App.tsx
import React from 'react';
import './App.css';
import BalanceSheet from './components/BalanceSheet';

const App: React.FC = () => {
    return (
        <div className="App">
           
            <BalanceSheet />
        </div>
    );
};

export default App;
