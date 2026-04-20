CREATE TABLE initiative_translations (
    initiative_id TEXT NOT NULL,
    lang TEXT NOT NULL,
    name TEXT,
    description TEXT,

    CONSTRAINT pk_initiative_translations PRIMARY KEY (initiative_id, lang),
    CONSTRAINT fk_initiative_translations_initiative
        FOREIGN KEY (initiative_id)
        REFERENCES initiatives(id)
        ON DELETE CASCADE
);
