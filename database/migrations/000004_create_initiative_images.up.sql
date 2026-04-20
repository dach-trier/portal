CREATE TABLE initiative_images (
    initiative_id TEXT NOT NULL,
    image_id INT NOT NULL,

    CONSTRAINT pk_initiative_images PRIMARY KEY (initiative_id, image_id),

    CONSTRAINT fk_initiative_images_initiative
        FOREIGN KEY (initiative_id)
        REFERENCES initiatives(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_initiative_images_image
        FOREIGN KEY (image_id)
        REFERENCES images(id)
        ON DELETE CASCADE
);
