ALTER TABLE initiatives RENAME TO projects;
ALTER TABLE initiative_translations RENAME COLUMN initiative_id TO project_id;
ALTER TABLE initiative_translations RENAME TO project_localizations;
ALTER TABLE initiative_images RENAME COLUMN initiative_id TO project_id;
ALTER TABLE initiative_images RENAME TO project_images;
