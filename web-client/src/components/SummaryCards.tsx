import React from 'react';
import { Transaction } from '../types';

interface Props {
  transactions: Transaction[];
}

const SummaryCards: React.FC<Props> = ({ transactions }) => {
  const today = new Date().toISOString().split('T')[0];
  const currentMonth = new Date().toISOString().slice(0, 7);

  const todayTotal = transactions
    .filter(t => t.date.startsWith(today))
    .reduce((sum, t) => sum + t.amount, 0);

  const monthTotal = transactions
    .filter(t => t.date.startsWith(currentMonth))
    .reduce((sum, t) => sum + t.amount, 0);

  return (
    <div id="summary-cards">
      <div className="summary-card">
        <h3>Today's Expenses</h3>
        <p>৳{todayTotal.toFixed(2)}</p>
      </div>
      <div className="summary-card">
        <h3>This Month</h3>
        <p>৳{monthTotal.toFixed(2)}</p>
      </div>
      <div className="summary-card budget-card">
        <h3>Total Transactions</h3>
        <p>{transactions.length}</p>
      </div>
    </div>
  );
};

export default SummaryCards;