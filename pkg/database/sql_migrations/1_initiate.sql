-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE menu (
    menu_id SERIAL PRIMARY KEY,
    menu_name VARCHAR(256) NOT NULL,
    menu_description VARCHAR(256) NOT NULL,
    menu_price INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE htrans (
    order_id VARCHAR(256) PRIMARY KEY,
    htrans_subtotal VARCHAR(256) NOT NULL,
    htrans_status INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE dtrans (
    order_id VARCHAR(256) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES htrans(order_id),
    menu_id BIGINT NOT NULL,
    FOREIGN KEY (menu_id) REFERENCES menu(menu_id),
    dtrans_quantity INT NOT NULL,
    dtrans_total VARCHAR(256) NOT NULL
);

CREATE TABLE carts (
    order_id VARCHAR(256) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES htrans(order_id),
    menu_id BIGINT NOT NULL,
    FOREIGN KEY (menu_id) REFERENCES menu(menu_id),
    cart_quantity INT NOT NULL
);

-- +migrate StatementEnd