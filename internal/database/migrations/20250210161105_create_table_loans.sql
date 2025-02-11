-- +migrate Up notransaction
CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    borrower_id INT NOT NULL,
    loan_amount DECIMAL(15,2) NOT NULL, -- Jumlah pinjaman
    interest_rate DECIMAL(5,2) NOT NULL, -- Suku bunga per tahun
    total_amount DECIMAL(15,2) NOT NULL, -- Total yang harus dibayar (loan_amount + bunga)
    weekly_payment DECIMAL(15,2) NOT NULL, -- Jumlah pembayaran per minggu
    total_weeks INT NOT NULL DEFAULT 50, -- Durasi pinjaman dalam minggu
    outstanding_amount DECIMAL(15,2) NOT NULL, -- Sisa saldo yang harus dibayar
    status VARCHAR(20) DEFAULT 'active', -- active, completed, delinquent
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (borrower_id) REFERENCES borrowers(id)
);

-- +migrate Down