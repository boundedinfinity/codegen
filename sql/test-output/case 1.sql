PRAGMA foreign_keys = ON;

CREATE TABLE name (
    id TEXT PRIMARY KEY,
    name_type_id TEXT,
    person_id TEXT,
    FOREIGN KEY (id) REFERENCES name_first(name_id),
    FOREIGN KEY (id) REFERENCES name_middle(name_id),
    FOREIGN KEY (id) REFERENCES name_last(name_id),
    FOREIGN KEY (name_type_id) REFERENCES name_type(id),
    FOREIGN KEY (person_id) REFERENCES person(id)
) WITHOUT ROWID;

CREATE INDEX idx_name ON name (name_type_id);

CREATE TABLE name_first (
    id TEXT PRIMARY KEY,
    name TEXT,
    "index" INTEGER,
    name_id TEXT,
    FOREIGN KEY (name_id) REFERENCES name(id)
) WITHOUT ROWID;

CREATE INDEX idx_name_first ON name_first (name);

CREATE TABLE name_middle (
    id TEXT PRIMARY KEY,
    name TEXT,
    "index" INTEGER,
    name_id TEXT,
    FOREIGN KEY (name_id) REFERENCES name(id)
) WITHOUT ROWID;

CREATE INDEX idx_name_middle ON name_middle (name);

CREATE TABLE name_last (
    id TEXT PRIMARY KEY,
    name TEXT,
    "index" INTEGER,
    name_id TEXT,
    FOREIGN KEY (name_id) REFERENCES name(id)
) WITHOUT ROWID;

CREATE INDEX idx_name_last ON name_last (name);

CREATE TABLE name_type (
    id TEXT PRIMARY KEY,
    name TEXT,
    description TEXT
) WITHOUT ROWID;

CREATE INDEX idx_name_type ON name_type (name);

CREATE TABLE person (
    id TEXT PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES name(person_id)
);

CREATE TABLE label (
    id TEXT PRIMARY KEY,
    name TEXT,
    description TEXT
);

CREATE INDEX idx_label ON label (name);

CREATE TABLE person_label (
    person_id TEXT,
    label_id TEXT,
    FOREIGN KEY (person_id) REFERENCES person(id),
    FOREIGN KEY (label_id) REFERENCES label(id)
) WITHOUT ROWID;

CREATE INDEX idx_person_label ON person_label (person_id);

CREATE INDEX idx_person_label ON person_label (label_id);