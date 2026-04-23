ALTER TABLE projects RENAME TO initiatives;
ALTER TABLE project_localizations RENAME COLUMN project_id TO initiative_id;
ALTER TABLE project_localizations RENAME TO initiative_translations;
ALTER TABLE project_images RENAME COLUMN project_id TO initiative_id;
ALTER TABLE project_images RENAME TO initiative_images;
