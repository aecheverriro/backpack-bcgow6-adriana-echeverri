/* Crea tabla temporal The Walking Dead */

CREATE TEMPORARY TABLE TWD 
SELECT eps.* FROM seasons sea INNER JOIN series ser ON sea.serie_id=ser.id INNER JOIN episodes eps ON eps.season_id=sea.id WHERE ser.title='The Walking Dead';

SELECT * FROM TWD;


/* En la base de datos “movies”, seleccionar una tabla donde crear un índice y luego chequear la creación del mismo. */
CREATE INDEX episode_num ON episodes (number);
SHOW INDEX FROM episodes;
/* Se crea un indice para el numero del episodio ya que este es un criterio al que se le suele hacer busqueda y puede ahorrar tiempo */
