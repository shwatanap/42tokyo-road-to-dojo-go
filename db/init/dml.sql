use 42game;

INSERT INTO
    users (name, token, high_score, coin)
VALUES
    ("Amy", UUID(), 0, 0);

INSERT INTO
    user_character_possessions (user_id, character_id)
VALUES
    (1, 1),
    (1, 2),
    (1, 3);

INSERT INTO
    rarities (rarity)
VALUES
    ("N"),
    ("R"),
    ("SR");

INSERT INTO
    characters (name, rarity_id)
VALUES
    ("Ken", 3),
    ("Tom", 2),
    ("Jaon", 2),
    ("Amy", 2),
    ("Jack", 2),
    ("Kebin", 1),
    ("Tommy", 1),
    ("Boby", 1),
    ("Kenny", 1),
    ("Ren", 1);
