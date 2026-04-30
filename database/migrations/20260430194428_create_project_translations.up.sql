CREATE TABLE project_translations (
    project_id BIGINT REFERENCES projects(id),
    lang       TEXT,
    name       TEXT,
    body       TEXT,

    PRIMARY KEY (project_id, lang)
);
