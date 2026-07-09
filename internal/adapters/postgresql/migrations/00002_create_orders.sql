-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
  "id" BIGSERIAL PRIMARY KEY,
  "userId" INT NOT NULL,
  "total" DECIMAL(10, 2) NOT NULL,
  "status" VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK ("status" IN ('pending', 'completed', 'cancelled')),
  "address" TEXT NOT NULL,
  "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_items (
  "id" BIGSERIAL PRIMARY KEY,
  "orderId" INT NOT NULL,
  "productId" INT NOT NULL,
  "quantity" INT NOT NULL,
  "price" DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY ("orderId") REFERENCES orders("id") ON DELETE CASCADE,
  FOREIGN KEY ("productId") REFERENCES products("id")
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd