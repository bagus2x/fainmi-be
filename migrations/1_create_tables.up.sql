CREATE TABLE profile(
    profile_id SERIAL,
    photo VARCHAR(512),
    username VARCHAR(25) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at INT NOT NULL,
    updated_at INT NOT NULL,
    PRIMARY KEY(profile_id)
);

CREATE TABLE background(
    background_id SERIAL,
    name VARCHAR(25) UNIQUE NOT NULL,
    description VARCHAR(512),
    created_at INT NOT NULL,
    updated_at INT NOT NULL,
    PRIMARY KEY(background_id)
);

CREATE TABLE button(
    button_id SERIAL,
    name VARCHAR(25) UNIQUE NOT NULL,
    description VARCHAR(512),
    created_at INT NOT NULL,
    updated_at INT NOT NULL,
    PRIMARY KEY(button_id)
);

CREATE TABLE font(
    font_id SERIAL,
    name VARCHAR(25) UNIQUE NOT NULL,
    description VARCHAR(512),
    created_at INT NOT NULL,
    updated_at INT NOT NULL,
    PRIMARY KEY(font_id)
);

CREATE TABLE style(
    profile_id INT NOT NULL UNIQUE,
    background_id INT,
    button_id INT,
    font_id INT,
    created_at INT NOT NULL,
    updated_at INT NOT NULL,
    PRIMARY KEY(profile_id),
    CONSTRAINT fk_style_profile FOREIGN KEY(profile_id) REFERENCES profile(profile_id)
    ON DELETE CASCADE,
    CONSTRAINT fk_style_background FOREIGN KEY(background_id) REFERENCES background(background_id)
    ON DELETE SET NULL,
    CONSTRAINT fk_style_button FOREIGN KEY(button_id) REFERENCES button(button_id)
    ON DELETE SET NULL,
    CONSTRAINT fk_style_font FOREIGN KEY(font_id) REFERENCES font(font_id)
    ON DELETE SET NULL
);

CREATE TABLE link(
    link_id SERIAL,
    profile_id INT NOT NULL,
    "order" INT NOT NULL,
    title VARCHAR(25),
    url VARCHAR(255),
    display BOOLEAN NOT NULL,
    created_at INT NOT NULL,
    updated_at INT NOT NULL,
    PRIMARY KEY(link_id, profile_id),
    CONSTRAINT fk_link_profile FOREIGN KEY(profile_id) REFERENCES profile(profile_id)
    ON DELETE CASCADE
);

CREATE TABLE "like"(
    link_id INT NOT NULL,
    owner_id INT NOT NULL,
    liker_id INT NOT NULL,
    created_at INT NOT NULL,
    PRIMARY KEY(link_id, owner_id, liker_id),
    CONSTRAINT fk_likes_link FOREIGN KEY(link_id, owner_id) REFERENCES link(link_id, profile_id)
    ON DELETE CASCADE,
    CONSTRAINT fk_likes_liker FOREIGN KEY(liker_id) REFERENCES profile(profile_id)
    ON DELETE CASCADE
);