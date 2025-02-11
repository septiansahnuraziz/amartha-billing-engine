-- +migrate Up notransaction
CREATE TABLE loan_schedules (
    id SERIAL PRIMARY KEY,
    loan_id INT NOT NULL,
    week_number INT NOT NULL, -- Minggu keberapa dari total pinjaman
    due_date DATE NOT NULL, -- Tanggal jatuh tempo
    amount DECIMAL(15,2) NOT NULL, -- Jumlah yang harus dibayar
    status VARCHAR(20) DEFAULT 'pending', -- pending, paid, missed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (loan_id) REFERENCES loans(id)
);

INSERT INTO borrowers (name, email, phone, created_at) VALUES
    ('Andi Pratama', 'andi.pratama@email.com', '081234567890', NOW()),
    ('Budi Santoso', 'budi.santoso@email.com', '081298765432', NOW()),
    ('Citra Dewi', 'citra.dewi@email.com', '081345678901', NOW()),
    ('Dian Permata', 'dian.permata@email.com', '081456789012', NOW()),
    ('Eko Wijaya', 'eko.wijaya@email.com', '081567890123', NOW()),
    ('Fajar Hidayat', 'fajar.hidayat@email.com', '081678901234', NOW()),
    ('Gita Maharani', 'gita.maharani@email.com', '081789012345', NOW()),
    ('Hendra Saputra', 'hendra.saputra@email.com', '081890123456', NOW());

-- +migrate Down