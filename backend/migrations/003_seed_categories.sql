-- Migration: 003_seed_categories
-- Insert default categories for expense tracking

INSERT INTO categories (name, type, priority, color_code, is_active) VALUES
('Food', 'EXPENSE', 'ESSENTIAL', '#FF5733', true),
('Transport', 'EXPENSE', 'ESSENTIAL', '#3B82F6', true),
('Shopping', 'EXPENSE', 'IMPORTANT', '#8B5CF6', true),
('Bills', 'EXPENSE', 'ESSENTIAL', '#EF4444', true),
('Entertainment', 'EXPENSE', 'LUXURY', '#10B981', true),
('Healthcare', 'EXPENSE', 'ESSENTIAL', '#F59E0B', true),
('Education', 'EXPENSE', 'IMPORTANT', '#6366F1', true),
('Others', 'EXPENSE', 'IMPORTANT', '#6B7280', true),
('Salary', 'INCOME', 'ESSENTIAL', '#059669', true),
('Freelance', 'INCOME', 'IMPORTANT', '#0D9488', true),
('Investment', 'INCOME', 'LUXURY', '#7C3AED', true)
ON CONFLICT (name) DO NOTHING;