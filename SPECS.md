# Especificaciones

Archivo `docker-compose.yml` por medio del cual se instancian **5** contenedores:

## Contenedor A - MongoDB

* Se encargará de correr un servidor de base de datos de `MongoDB`. Esta base de datos almacenará la información de nuestra aplicación en general

## Contenedor B - DBMS/DBAdmin

* Se encargará de ejecutar alguno de los `DBMS` que vimos en clase, dependiendo del motor de base de datos elegido. También es posible investigar e implementar algún otro `DBMS` que funcione sobre `Docker` y que sea más sencillo y/o a gusto del equipo
* Este contenedor se conectará con la base de datos del contenedor `A` de tal manera que podamos visualizar el esquema de datos creado para este proyecto

## Contenedor C - API Scraper/Fetcher/Consumer

* Este contenedor se encargará de correr un scraper/consumer para obtener los datos de una API
* La API es opcional de entre las disponibles en el listado de [Public APIs](https://github.com/public-apis/public-apis), pero debe de contar con al menos 3 diferentes endpoints/recursos. Por ejemplo, tenemos la API de [Marvel](https://developer.marvel.com/docs), y `/v1/public/characters`, `/v1/public/comics`, `/v1/public/creators`, `/v1/public/events`, `/v1/public/series` y `/v1/public/stories` son cada uno diferentes recursos, es decir que tenemos **6** endpoints/recursos en esta API
* La API elegida debe de ser diferente entre cada uno de los **4** equipos, y no puede elegirse la API de Marvel que se mostró en el ejemplo del punto anterior
* Se deben de obtener todos los datos disponibles en la `API` para cada endpoint/recurso y estos no deben de guardarse en una sola tabla, por lo que hay que analizar, definir y estructurar correctamente la base de datos al momento de diseñarla

## Contenedor D - Front-End

* Para esta parte puedes utilizar `Flask` o cualquier otro framework `MVC` como `Django` por ejemplo. Incluso puedes desarrollarlo "_from scratch_" si lo consideras pertinente y necesario :wink:
* Este contenedor estará a cargo de correr una pequeña "_web-app_" que muestre un índice de todos los endpoints/recursos obtenidos de la `API` a través de la página principal. Después, para cada endpoint/recurso en específico tendremos un link a una página nueva
* En la página nueva se deben de mostrar todos los datos obtenidos para cada endpoint/recurso en específico, es decir, partiendo del ejemplo de la API de Marvel, suponiendo que yo el usuario di click al link del endpoint `/characters/78` en la página del índice, la nueva página me deberá de mostrar toda la información relacionada al personaje con ese `ID = 78`
* Se puede crear una página "_índice_" para cada uno de los diferentes endpoints/recursos, o bien, una sola página de este tipo que agrupe a todos los endpoints/recursos pero separados de alguna manera explícita (por ejemplo una etiqueta `<h2>` o una etiqueta `<hr />`)
* Si usas `Flask`, deberás de utilizar el sistema de templates [`Jinja`](https://flask.palletsprojects.com/en/1.1.x/tutorial/templates/) nativo en el framework de `Flask` para mostrar el `HTML`/markup de la aplicación. Queda a tu criterio el uso de `CSS`, `JS` y otras artimañas para darle formato a la web-app :wink:
* Finalmente, para cada uno de los endpoints/recursos que hayas agregado, habilita la opción de crear registros por medio del mismo, es decir, ser capaz de hacer un request de tipo [`HTTP - POST`](https://developer.mozilla.org/es/docs/Web/HTTP/Methods/POST) a esa ruta en tu frontend, de tal manera que alguien que consuma tu "_web-app_" también pueda crear registros por medio de la misma. Partiendo una vez más del ejemplo de Marvel, si yo quisiera crear un nuevo personaje en la `BD`, entonces yo tendría que mandar un `POST` request al endpoint de `/characters/` en mi "_web-app_". Asegúrate de que para cada endpoint/recurso se pasen todos los datos necesarios para poder crear un registro en la BD con éxito. Puedes hacer pruebas para este punto en específico con herramientas gratuitas como [`Postman`](https://www.postman.com/) o [`Insomnia`](https://insomnia.rest/)

## Contenedor E - Message Queue Broker

* Este contenedor se encargará de ejecutar un message queue broker como [RabbitMQ](https://www.rabbitmq.com/), [Apache Kafka](https://kafka.apache.org/) o [Celery](https://docs.celeryproject.org/en/stable/getting-started/introduction.html)
* El nuevo servicio pasará a ser el intermediario entre el contenedor `D` del front end y el contenedor `A` de la base de datos, de tal manera que habrá que sobreescribir un poco la lógica en el contenedor `D` para la creación de registros por medio de `HTTP POST` de tal manera que ahora soporte operaciones asíncronas por medio del contenedor `E`, es decir, si escogiste `RabbitMQ` como message queue broker, entonces al enviar peticiones `POST` para crear un nuevo registro estas pasarían por `RabbitMQ` primero, y este sería el que se encargaría de enviarlo al contenedor/servicio correcto que este suscrito a un tipo concreto de mensajes, en este caso el de `PostgreSQL`. Lo mismo aplica en la otra dirección, es decir `PostgreSQL` respondiéndole al servicio del front end por medio de `RabbitMQ`. Pueden crear un endpoint/recurso nuevo para este punto, de tal manera que la aplicación tenga soporte tanto síncrono como asíncrono, o bien, sobreescribir la funcionalidad del endpoint original, como se más sencillo para ustedes
