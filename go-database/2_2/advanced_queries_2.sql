/*
Agregar una película a la tabla movies.
Agregar un género a la tabla genres.
Asociar a la película del punto 1. con el género creado en el punto 2.
Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.
Crear una tabla temporal copia de la tabla movies.
Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
Obtener la lista de todos los géneros que tengan al menos una película.
Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
Crear un índice sobre el nombre en la tabla movies.
Chequee que el índice fue creado correctamente.
En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.
¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta
*/
INSERT INTO movies VALUES (22, null, null, 'Enola Holmes 2', 8.5, 3, '2022-11-05', 125, 3);
INSERT INTO genres VALUES (13, '2022-11-08', null, 'Romance', 13, 1);
UPDATE movies SET genre_id=13 WHERE id=22;
UPDATE actors SET favorite_movie_id=22 WHERE id=40;
CREATE TEMPORARY TABLE MOV SELECT * FROM movies;
DELETE FROM MOV WHERE awards < 5;
SELECT gen.name, COUNT(mov.genre_id) AS count_per_genre FROM movies mov INNER JOIN genres gen ON mov.genre_id=gen.id GROUP BY mov.genre_id HAVING count_per_genre>=1;
SELECT act.* FROM actor_movie actmov INNER JOIN actors act ON actmov.actor_id=act.id INNER JOIN movies mov ON actmov.movie_id=mov.id WHERE mov.awards>3;
CREATE INDEX nombre_tabla ON movies (title);
SHOW INDEX FROM movies;
/* Si habría una mejora dado que la busqueda por titulo revisaría uno por uno y con el indice esto mejoraria.
Este mismo proceso se podría realizar en la de series con el mismo objetivo que se hizo en movies */
