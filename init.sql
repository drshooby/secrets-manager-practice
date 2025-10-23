-- Spanish Pop Songs Database (seeding)

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    artist VARCHAR(100) NOT NULL
);

INSERT INTO songs (title, artist) VALUES
('Lentamente', 'Ana Mena'),
('Ya Es Hora', 'Ana Mena'),
('El Chisme', 'Ana Mena'),
('Plan Perfecto', 'David Bisbal'),
('Ajedrez', 'David Bisbal'),
('A Partir De Hoy', 'David Bisbal'),
('Enemigos', 'Aitana'),
('Quieres', 'Aitana'),
('Muy Loco', 'Carlos Baute'),
('Mesa Para Dos', 'Julia Medina'),
('Quiero Decirte', 'Abraham Mateo'),
('Clavaito', 'Abraham Mateo'),
('Ambulancia', 'Camilo'),
('LAS BURBUJAS DEL JACUZZI', 'India Martinez'),
('Feliz Por Conocerte', 'Niko Rubio');