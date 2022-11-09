/* Populate database */
INSERT INTO empleado VALUES ('E-0001', 'Cesar', 'Pinero', 'Vendedor', '2018-05-12', 80000, 15000, 'D-000-4');
INSERT INTO empleado VALUES ('E-0002', 'Yosep', 'Kowaleski', 'Analista', '2015-07-14', 140000, 0, 'D-000-2');
INSERT INTO empleado VALUES ('E-0003', 'Mariela', 'Barrios', 'Director', '2014-06-05', 185000, 0, 'D-000-3');
INSERT INTO empleado VALUES ('E-0004', 'Jonathan', 'Brezezicki', 'Vendedor', '2015-06-03', 85000, 10000, 'D-000-4');
INSERT INTO empleado VALUES ('E-0005', 'Daniel', 'Aguilera', 'Vendedor', '2018-03-03', 83000, 10000, 'D-000-4');
INSERT INTO empleado VALUES ('E-0006', 'Mito', 'Barchuk', 'Presidente', '2014-06-05', 190000, 0, 'D-000-3');
INSERT INTO empleado VALUES ('E-0007', 'Emilio', 'Galarza', 'Desarrollador', '2014-08-02', 60000, 0, 'D-000-1');

INSERT INTO departamento VALUES ('D-000-1', 'Software', 'Los Tigres');
INSERT INTO departamento VALUES ('D-000-2', 'Sistemas', 'Guadalupe');
INSERT INTO departamento VALUES ('D-000-3', 'Contabilidad', 'La Roca');
INSERT INTO departamento VALUES ('D-000-4', 'Ventas', 'Plata');

/*
Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
Visualizar los departamentos con más de cinco empleados.
Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
Mostrar el nombre del empleado que tiene el salario más bajo.
Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
*/

SELECT dep.nombre_depto, dep.localidad, emp.puesto FROM departamento dep INNER JOIN empleado emp ON dep.depto_nro=emp.departamento_depto_nro;
SELECT dep.nombre_depto, COUNT(emp.departamento_depto_nro) AS total_gen FROM departamento dep INNER JOIN empleado emp ON dep.depto_nro=emp.departamento_depto_nro GROUP BY emp.departamento_depto_nro HAVING total_gen > 5;
SELECT emp.nombre, emp.salario, dep.nombre_depto FROM departamento dep INNER JOIN empleado emp ON dep.depto_nro=emp.departamento_depto_nro WHERE emp.puesto=(SELECT puesto FROM empleado WHERE nombre='Mito' AND apellido = 'Barchuk');
SELECT emp.* FROM departamento dep INNER JOIN empleado emp ON dep.depto_nro=emp.departamento_depto_nro WHERE dep.nombre_depto='Contabilidad' ORDER BY emp.nombre;
SELECT nombre FROM empleado ORDER BY salario ASC LIMIT 1;
SELECT emp.* FROM departamento dep INNER JOIN empleado emp ON dep.depto_nro=emp.departamento_depto_nro WHERE dep.nombre_depto='Ventas' ORDER BY emp.salario DESC LIMIT 1;
