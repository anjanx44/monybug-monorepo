export interface Category {
  id: number;
  name: string;
  type: 'INCOME' | 'EXPENSE';
  priority?: string;
  color_code?: string;
  is_active: boolean;
}

export interface Transaction {
  id: number;
  category_id: number;
  date: string;
  amount: number;
  currency: string;
  description: string;
  tags?: string;
  created_at: string;
  category?: Category;
}