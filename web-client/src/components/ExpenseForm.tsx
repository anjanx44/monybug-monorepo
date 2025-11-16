import React, { useState, useEffect } from 'react';
import { Category } from '../types';
import { api } from '../services/api';

interface Props {
  onTransactionAdded: () => void;
}

const ExpenseForm: React.FC<Props> = ({ onTransactionAdded }) => {
  const [categories, setCategories] = useState<Category[]>([]);
  const [formData, setFormData] = useState({
    amount: '',
    category_id: '',
    description: '',
    date: new Date().toISOString().split('T')[0]
  });

  useEffect(() => {
    loadCategories();
  }, []);

  const loadCategories = async () => {
    try {
      const response = await api.getCategories();
      setCategories(response.data);
    } catch (error) {
      console.error('Failed to load categories:', error);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await api.createTransaction({
        ...formData,
        amount: parseFloat(formData.amount),
        category_id: parseInt(formData.category_id)
      });
      setFormData({
        amount: '',
        category_id: '',
        description: '',
        date: new Date().toISOString().split('T')[0]
      });
      onTransactionAdded();
    } catch (error) {
      console.error('Failed to create transaction:', error);
    }
  };

  return (
    <div className="card">
      <h2>Quick Entry</h2>
      <form id="expense-form" onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="amount">Amount (৳)</label>
          <input
            type="number"
            id="amount"
            placeholder="1250.00"
            step="0.01"
            value={formData.amount}
            onChange={(e) => setFormData({...formData, amount: e.target.value})}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="category">Category</label>
          <select
            id="category"
            value={formData.category_id}
            onChange={(e) => setFormData({...formData, category_id: e.target.value})}
            required
          >
            <option value="">Select Category</option>
            {categories.map(cat => (
              <option key={cat.id} value={cat.id}>{cat.name}</option>
            ))}
          </select>
        </div>

        <div className="form-group">
          <label htmlFor="description">Description</label>
          <input
            type="text"
            id="description"
            placeholder="Enter description"
            value={formData.description}
            onChange={(e) => setFormData({...formData, description: e.target.value})}
          />
        </div>

        <div className="form-group">
          <label htmlFor="date">Date</label>
          <input
            type="date"
            id="date"
            value={formData.date}
            onChange={(e) => setFormData({...formData, date: e.target.value})}
            required
          />
        </div>

        <button type="submit" className="btn-success">
          + Add Transaction
        </button>
      </form>
    </div>
  );
};

export default ExpenseForm;