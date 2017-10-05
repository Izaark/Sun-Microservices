# Sun-Microservices
**Este repositorio es una ejemplo de como se realizaran los siguientes microservicios !**

# Reglas
- La rama master siempre debe seguir la versión que tengamos en producción. Por lo cual queda prohibido hacer push directamente sobre master.

- Para poder subir cambios a producción debemos solicitar un pull request para hacer las pruebas necesarias y revisar el código.
- Cada mimebro del equipo es responsable de aceptar un cambio en especifico para produccion **(es el cuello de botella para lanzar a produccion)**

- Evita tiempo en revisar código, **cuidado al hacer merge** xD
- [![N|Solid](https://cldup.com/lldn2xu-wh.gif)](https://nodesource.com/products/nsolid)



### Conceptos considerables a leer: 

* [API Management] - Que es y para que sirve?
* [Patron MVC] - Por que MVC ?
* [Microservicios] - Entendiendo los Microservicios !

And of course Dillinger itself is open source with a [public repository][dill]
 on GitHub.

### Instalacion

Instala Docker por este medio [Docker](https://www.digitalocean.com/community/tutorials/como-instalar-y-usar-docker-en-ubuntu-16-04-es)
Una vez instalado docker Haz un pull a la imagen y sigue estas [instrucciones](https://store.docker.com/images/golang) para tener el contenedor corriendo **Golang**
```sh
$ docker pull golang
```
En todos los microservicios usaremos el  API Managemente **WSO2** si no lo tienes sigue [aquí](https://github.com/wso2/docker-apim/tree/master/dockerfile) las instrucciones para instalar WSO2. De igual manera funciona en una imagen contenida por docker

| Plugin | README |
| ------ | ------ |
| Gin/Gonic | [plugins/GinGonic/README.md] [PlDb] |
| rethinkdb | [plugins/rethinkdb/docs] [PlGh] |


### Desarrollo
Nuestra arquitectura MVC esta compuesta de la siguiente manera:
[![N|Solid](https://cldup.com/RfII1ZbtON.png)](https://nodesource.com/products/nsolid)
 - **pokemon**: Es el nombre del repositorio del microservicio.
 - **assets**: Dentro de este folder se encuentra la carpeta *images*, aqui estarán las carpetas que contienen las fotos!. Hay un nivel mas de carpeta con el nombre *pokemons* este nombre varia dependiendo del la relación que lleven las imagenes. **Para los microservicios esta ruta estará en AWS S3, sólo como referencia se dejo en esta prueba !:**
 - **config**: Esta carpeta contiene a *database.go* el cual realiza la coneccion a la base de datos.
 - **controllers**: La carpeta contiene el archivo *pokemon.go* en donde se tiene la lógica de programación.
 - **models**: En esta carpeta se encuentran el archivo *pokemon.go* el cual contiene los structs que serán consultados a la base de datos.
 - **main.go**: Es nuestro archivo principal, el cual se encuentra en el nivel principal del repositorio, *main.go* es el archivo que se compila en la terminal !.
 - **pokemon_swagger.json**: Adicionalmente, en el nivel principal del repo se encuentra el archivo *pokemon_swagger.json*, este sirve para diseñar el API para montarlo en WSO2. swagger se codea con YAML para generar un JSON.

   [Patron MVC]: <https://codigofacilito.com/articles/22>
   [Microservicios]: <https://openwebinars.net/blog/microservicios-que-son/>
   [API Management]: <https://www.paradigmadigital.com/dev/api-management-que-es-y-para-que-sirve/>


   [PlDb]: <https://github.com/gin-gonic/gin>
   [PlGh]: <https://rethinkdb.com/docs/>
