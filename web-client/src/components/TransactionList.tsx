import React from 'react';
import { Transaction } from '../types';
import { api } from '../services/api';

interface Props {
  transactions: Transaction[];
  onTransactionUpdated: () => void;
}

const TransactionList: React.FC<Props> = ({ transactions, onTransactionUpdated }) => {
  const getCategoryIcon = (categoryName: string) => {
    const icons: { [key: string]: string } = {
      'Food': '🍔',
      'Transport': '⛽',
      'Shopping': '🛒',
      'Bills': '💡',
      'Entertainment': '🎬',
      'Others': '📝'
    };
    return icons[categoryName] || '📝';
  };

  const handleDelete = async (id: number) => {
    if (window.confirm('Are you sure you want to delete this transaction?')) {
      try {
        await api.deleteTransaction(id);
        onTransactionUpdated();
      } catch (error) {
        console.error('Failed to delete transaction:', error);
      }
    }
  };

  const groupedTransactions = transactions.reduce((groups, transaction) => {
    const date = new Date(transaction.date).toLocaleDateString();
    if (!groups[date]) {
      groups[date] = [];
    }
    groups[date].push(transaction);
    return groups;
  }, {} as { [key: string]: Transaction[] });

  return (
    <div id="transaction-list-container">
      <div className="card">
        <h2>Recent Transactions</h2>
        {Object.keys(groupedTransactions).length === 0 ? (
          <p>No transactions found. Add your first expense above!</p>
        ) : (
          Object.entries(groupedTransactions).map(([date, dayTransactions]) => (
            <div key={date} className="date-group">
              <div className="date-header">{date}</div>
              {dayTransactions.map(transaction => (
                <div key={transaction.id} className="transaction-item">
                  <div className="category-icon">
                    {getCategoryIcon(transaction.category?.name || 'Others')}
                  </div>
                  <div className="details">
                    <p className="description">{transaction.description}</p>
                    <p className="time">{transaction.category?.name}</p>
                  </div>
                  <div className="amount-display">
                    ৳{transaction.amount.toFixed(2)}
                  </div>
                  <div className="actions">
                    <button 
                      className="btn-danger"
                      onClick={() => handleDelete(transaction.id)}
                    >
                      🗑️
                    </button>
                  </div>
                </div>
              ))}
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default TransactionList;