import axios from 'axios';
import { Category, Transaction } from '../types';

const API_BASE_URL = 'http://localhost:8080/api/v1';

export const api = {
  // Categories
  getCategories: () => axios.get<Category[]>(`${API_BASE_URL}/categories`),
  createCategory: (category: Partial<Category>) => 
    axios.post<Category>(`${API_BASE_URL}/categories`, category),
  
  // Transactions
  getTransactions: () => axios.get<Transaction[]>(`${API_BASE_URL}/transactions`),
  createTransaction: (transaction: Partial<Transaction>) => 
    axios.post<Transaction>(`${API_BASE_URL}/transactions`, transaction),
  updateTransaction: (id: number, transaction: Partial<Transaction>) => 
    axios.put(`${API_BASE_URL}/transactions/${id}`, transaction),
  deleteTransaction: (id: number) => 
    axios.delete(`${API_BASE_URL}/transactions/${id}`)
};