/* Populate database */
INSERT INTO AUTOR VALUES (1, 'Cesar', 'Mexicano');
INSERT INTO AUTOR VALUES (2, 'Manuel', 'Colombiano');
INSERT INTO AUTOR VALUES (3, 'Luisa', 'Americano');
INSERT INTO AUTOR VALUES (4, 'Maria', 'Argentino');
INSERT INTO AUTOR VALUES (5, 'Rogelio', 'Venezolano');

INSERT INTO ESTUDIANTE VALUES (1, 'Xiomara', 'Gaze', 'Calle 54', 'Profesora', 54);
INSERT INTO ESTUDIANTE VALUES (2, 'Alba', 'Gonzalez', 'Avenida 14', 'Actriz', 21);
INSERT INTO ESTUDIANTE VALUES (3, 'Jane', 'Villanueva', 'Calle 13', 'Cantante', 30);
INSERT INTO ESTUDIANTE VALUES (4, 'Alejandro', 'Urbina', 'Calle 12', 'Arquitecto', 64);
INSERT INTO ESTUDIANTE VALUES (5, 'Federico', 'Rueda', 'Avenida 1', 'Medico', 50);

INSERT INTO LIBRO VALUES (1, 'Cerebro de pan', 'Castellana', 'Nutricional');
INSERT INTO LIBRO VALUES (2, 'Catching Fire', 'Bonsai', 'Ficcion');
INSERT INTO LIBRO VALUES (3, 'Don Quijote', 'Barca', 'Clasico');
INSERT INTO LIBRO VALUES (4, 'Mocking Jay', 'Bonsai', 'Ficcion');
INSERT INTO LIBRO VALUES (5, 'The Giver', 'Avion de papel', 'Drama');

INSERT INTO LIBROAUTOR VALUES (1, 3);
INSERT INTO LIBROAUTOR VALUES (3, 2);
INSERT INTO LIBROAUTOR VALUES (2, 5);
INSERT INTO LIBROAUTOR VALUES (4, 1);
INSERT INTO LIBROAUTOR VALUES (5, 4);

INSERT INTO PRESTAMO VALUES ("2022-11-05", "2022-11-25", "NO", 1, 3);
INSERT INTO PRESTAMO VALUES ("2020-01-25", "2020-02-25", "SI", 3, 3);
INSERT INTO PRESTAMO VALUES ("2019-16-05", "2020-16-05", "SI", 5, 3);

/* Queries */ /*
Listar los datos de los autores.
Listar nombre y edad de los estudiantes
¿Qué estudiantes pertenecen a la carrera informática?
¿Qué autores son de nacionalidad francesa o italiana?
¿Qué libros no son del área de internet?
Listar los libros de la editorial Salamandra.
Listar los datos de los estudiantes cuya edad es mayor al promedio.
Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
¿Qué libros se prestaron al lector “Filippo Galli”?
Listar el nombre del estudiante de menor edad.
Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
Listar los libros que pertenecen a la autora J.K. Rowling.

Listar títulos de los libros que debían devolverse el 16/07/2021.
*/
SELECT * FROM AUTOR;
SELECT Nombre, Apellido FROM ESTUDIANTE;
SELECT * FROM ESTUDIANTE WHERE Carrera='Informatica';
SELECT * FROM AUTORES WHERE Nacionalidad='Francesa' OR Nacionalidad='Italiana';
SELECT * FROM LIBRO WHERE Area!='Internet';
SELECT * FROM LIBRO WHERE Editorial='Salamandra';
SELECT * FROM ESTUDIANTE WHERE Edad > (SELECT AVG(Edad) FROM ESTUDIANTE);
SELECT * FROM ESTUDIANTE WHERE Apellido LIKE 'G%';
SELECT auto.Nombre FROM LIBROAUTOR liauto INNER JOIN LIBRO li ON liauto.LIBRO_idLibro=li.idLibro INNER JOIN AUTOR auto ON liauto.AUTOR_idAutor=auto.idAutor WHERE li.Titulo='El Universo: Guía de viaje';
SELECT * FROM PRESTAMO prest INNER JOIN ESTUDIANTE est ON prest.ESTUDIANTE_idLector=est.idLector WHERE est.Nombre='Filippo' AND est.Apellido='Galli';
SELECT Nombre FROM ESTUDIANTE ORDER BY Edad ASC LIMIT 1;
SELECT auto.Nombre FROM PRESTAMO prest INNER JOIN LIBRO li ON prest.LIBRO_idLibro=li.idLibro INNER JOIN ESTUDIANTE est ON prest.ESTUDIANTE_idLector=est.idLector WHERE li.Area='Base de Datos';
SELECT li.* FROM LIBROAUTOR liauto INNER JOIN LIBRO li ON liauto.LIBRO_idLibro=li.idLibro INNER JOIN AUTOR auto ON liauto.AUTOR_idAutor=auto.idAutor WHERE auto.Nombre='J.K Rowling';
SELECT li.Titulo FROM PRESTAMO prest INNER JOIN LIBRO li ON prest.LIBRO_idLibro=li.idLibro WHERE prest.FechaDevolucion='2021-07-16';
