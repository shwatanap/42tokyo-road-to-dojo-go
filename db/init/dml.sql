use 42game;

INSERT INTO
    users (name, token, high_score, coin)
VALUES
    ("Amy", UUID(), 50, 10),
    ("Bob", UUID(), 60, 42),
    ("Cate", UUID(), 100, 30),
    ("Dian", UUID(), 30, 40),
    ("Ema", UUID(), 70, 50),
    ("Finn", UUID(), 40, 70),
    ("Genie", UUID(), 30, 60),
    ("Hello", UUID(), 60, 1000),
    ("Iron", UUID(), 90, 40),
    ("Jack", UUID(), 10, 660),
    ("Ken", UUID(), 30, 770),
    ("Len", UUID(), 30, 20),
    ("Mikele", UUID(), 90, 0),
    ("Nick", UUID(), 90, 340),
    ("Oruga", UUID(), 40, 420),
    ("Prin", UUID(), 30, 210),
    ("Queen", UUID(), 50, 720),
    ("Reck", UUID(), 80, 90);

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
