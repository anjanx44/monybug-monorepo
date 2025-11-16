-- MonyBug Analytics Queries
-- Sample queries for expense analysis and reporting

-- 1. Monthly Category Breakdown
SELECT 
    c.name as category,
    c.priority,
    SUM(t.amount) as total_spent,
    COUNT(t.transaction_id) as transaction_count
FROM transactions t 
JOIN categories c ON t.category_id = c.category_id
WHERE t.transaction_date >= DATE_TRUNC('month', CURRENT_DATE)
GROUP BY c.name, c.priority
ORDER BY total_spent DESC;

-- 2. Budget vs Actual Spending
SELECT 
    c.name as category,
    b.monthly_limit,
    COALESCE(SUM(t.amount), 0) as actual_spent,
    (b.monthly_limit - COALESCE(SUM(t.amount), 0)) as remaining_budget,
    ROUND((COALESCE(SUM(t.amount), 0) / b.monthly_limit * 100), 2) as budget_used_percent
FROM budgets b
JOIN categories c ON b.category_id = c.category_id
LEFT JOIN transactions t ON c.category_id = t.category_id 
    AND t.transaction_date >= b.start_date 
    AND t.transaction_date <= b.end_date
WHERE b.is_active = true
GROUP BY c.name, b.monthly_limit;

-- 3. Daily Spending Trends (Last 30 days)
SELECT 
    transaction_date,
    SUM(amount) as daily_total,
    COUNT(*) as transaction_count
FROM transactions 
WHERE transaction_date >= CURRENT_DATE - INTERVAL '30 days'
    AND transaction_date <= CURRENT_DATE
GROUP BY transaction_date
ORDER BY transaction_date;

-- 4. Top Spending Categories (Current Month)
SELECT 
    c.name,
    c.priority,
    SUM(t.amount) as total_spent
FROM transactions t
JOIN categories c ON t.category_id = c.category_id
WHERE EXTRACT(MONTH FROM t.transaction_date) = EXTRACT(MONTH FROM CURRENT_DATE)
    AND EXTRACT(YEAR FROM t.transaction_date) = EXTRACT(YEAR FROM CURRENT_DATE)
    AND c.type = 'EXPENSE'
GROUP BY c.name, c.priority
ORDER BY total_spent DESC
LIMIT 10;

-- 5. Monthly Comparison (This Month vs Last Month)
WITH current_month AS (
    SELECT SUM(amount) as current_total
    FROM transactions 
    WHERE EXTRACT(MONTH FROM transaction_date) = EXTRACT(MONTH FROM CURRENT_DATE)
        AND EXTRACT(YEAR FROM transaction_date) = EXTRACT(YEAR FROM CURRENT_DATE)
),
last_month AS (
    SELECT SUM(amount) as last_total
    FROM transactions 
    WHERE transaction_date >= DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
        AND transaction_date < DATE_TRUNC('month', CURRENT_DATE)
)
SELECT 
    cm.current_total,
    lm.last_total,
    (cm.current_total - lm.last_total) as difference,
    ROUND(((cm.current_total - lm.last_total) / lm.last_total * 100), 2) as percent_change
FROM current_month cm, last_month lm;

-- 6. Expense Pattern by Day of Week
SELECT 
    EXTRACT(DOW FROM transaction_date) as day_of_week,
    CASE EXTRACT(DOW FROM transaction_date)
        WHEN 0 THEN 'Sunday'
        WHEN 1 THEN 'Monday'
        WHEN 2 THEN 'Tuesday'
        WHEN 3 THEN 'Wednesday'
        WHEN 4 THEN 'Thursday'
        WHEN 5 THEN 'Friday'
        WHEN 6 THEN 'Saturday'
    END as day_name,
    AVG(amount) as avg_spending,
    COUNT(*) as transaction_count
FROM transactions
WHERE transaction_date >= CURRENT_DATE - INTERVAL '90 days'
GROUP BY EXTRACT(DOW FROM transaction_date)
ORDER BY day_of_week;