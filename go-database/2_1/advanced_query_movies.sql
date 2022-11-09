/*
Mostrar el título y el nombre del género de todas las series.
Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.

Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
*/

SELECT title, name FROM series INNER JOIN genres ON series.genre_id = genres.id;
SELECT ep.title, ac.first_name, ac.last_name FROM actor_episode acep INNER JOIN episodes ep ON acep.episode_id = ep.id INNER JOIN actors ac ON ac.id = acep.actor_id;
SELECT ser.title, COUNT(sea.serie_id) FROM seasons sea INNER JOIN series ser ON sea.serie_id=ser.id GROUP BY ser.title;
SELECT gen.name, COUNT(gen.name) AS total_gen FROM genres gen INNER JOIN movies mov ON gen.id=mov.genre_id GROUP BY gen.name HAVING total_gen >= 3;
SELECT DISTINCT ac.first_name, ac.last_name FROM actor_movie acmo INNER JOIN actors ac ON acmo.actor_id = ac.id INNER JOIN movies mov ON acmo.movie_id = mov.id WHERE mov.title LIKE "La Guerra de las galaxias%";
