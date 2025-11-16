import React, { useState, useEffect } from 'react';
import ExpenseForm from './components/ExpenseForm';
import SummaryCards from './components/SummaryCards';
import TransactionList from './components/TransactionList';
import { Transaction } from './types';
import { api } from './services/api';
import './App.css';

function App() {
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  const loadTransactions = async () => {
    try {
      const response = await api.getTransactions();
      setTransactions(response.data);
    } catch (error) {
      console.error('Failed to load transactions:', error);
    }
  };

  useEffect(() => {
    loadTransactions();
  }, []);

  return (
    <div className="container">
      <h1>MonyBug Tracker</h1>
      <ExpenseForm onTransactionAdded={loadTransactions} />
      <SummaryCards transactions={transactions} />
      <TransactionList 
        transactions={transactions} 
        onTransactionUpdated={loadTransactions}
      />
    </div>
  );
}

export default App;
