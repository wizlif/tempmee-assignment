CREATE VIEW v_orders
AS
SELECT 
    o.id,
    o.price,
    o.total,
    o.status,
    b.title,
    u.email,
    o.created_at
FROM orders o 
JOIN books b ON b.id == o.book_id 
JOIN users u ON u.id == o.user_id 
