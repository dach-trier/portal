CREATE TABLE project_assets (
    project_id BIGINT REFERENCES projects(id),
    asset_id   BIGINT REFERENCES assets(id),

    PRIMARY KEY (project_id, asset_id)
);
