CREATE TABLE transactions(
     id BIGSERIAL PRIMARY KEY,
     from_wallet_id BIGINT NOT NULL,
     to_wallet_id BIGINT NOT NULL,
     amount BIGINT NOT NULL DEFAULT 0,
     type VARCHAR(20) NOT NULL,
     status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);