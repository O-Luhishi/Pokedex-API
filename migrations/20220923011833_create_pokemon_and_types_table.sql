-- +goose Up
-- +goose StatementBegin
-- Create Type Then Add Type As FK To Pokemon Table
CREATE TABLE IF NOT EXISTS pokemon_types
(
    "id"           BIGSERIAL PRIMARY KEY,
    "name"         VARCHAR(255) NOT NULL,
    "description"  VARCHAR(255) NOT NULL,
    "created_at"   TIMESTAMP    NOT NULL
);

CREATE TABLE IF NOT EXISTS pokemon
(
    "id"           BIGSERIAL PRIMARY KEY,
    "name"         VARCHAR(255) NOT NULL,
    "abilities"    TEXT [] NOT NULL,
    "type_id"      BIGINT NOT NULL,
    "created_at"   TIMESTAMP NOT NULL
);

ALTER TABLE "pokemon" ADD FOREIGN KEY ("type_id") REFERENCES pokemon_types ("id");
CREATE INDEX ON "pokemon" ("type_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pokemon;
DROP TABLE IF EXISTS pokemon_types;
-- +goose StatementEnd
