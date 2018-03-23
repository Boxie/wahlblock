-- +migrate Up

CREATE TABLE  blocks (
  id bigserial PRIMARY KEY,
  proof INT NOT NULL,
  previous_hash VARCHAR(255) NOT NULL,
  created_at TIMESTAMP
);

CREATE TABLE transactions (
  id bigserial PRIMARY KEY,
  block_index BIGINT NOT NULL REFERENCES blocks,
  ballot_id VARCHAR (255) NOT NULL,
  voting VARCHAR(255) NOT NULL,
  created_at TIMESTAMP
);

CREATE TABLE nodes (
  id bigserial PRIMARY KEY,
  host VARCHAR (255) NOT NULL,
  port INT NOT NULL,
  active BIT NOT NULL,
  last_communication TIMESTAMP
);