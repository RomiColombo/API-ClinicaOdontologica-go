## Examen final - Colombo, Romina

En el archivo sql se encuentra el script para crear las 3 bases de datos e insertar campos en cada una de ellas asi no se inicializan vacías.
A su vez, al final del script estan comentadas tres querys para dropear las tablas en caso de ser necesario.

El archivo clinicaOdontologica.postman_collection.json se puede importar en el postman para probar los endpoints
Hay 3 carpetas correspondientes a odontologo, paciente y turnos con todos los métodos creados.

En el caso de querer eliminar un odontologo o un paciente que tiene asociado un turno se produce un error de MySql ya que
no se pueden eliminar claves foráneas.

En caso de querer cambiar el path de la base de datos se puede hacer en el .env ya que esta almacenada como variable de entorno
