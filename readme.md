# QuickFix Backend

Este repositorio contiene el código fuente del backend para QuickFix, una plataforma que conecta usuarios con profesionales que ofrecen servicios a domicilio.

## Requisitos

Antes de comenzar, asegúrate de tener instalado lo siguiente:

- Docker
- Docker Compose

## Configuración del Proyecto

1. **Clonar el Repositorio**

   ```bash
   git clone git@github.com:jizambrana5/quick-fix-back.git
   cd quick-fix-back

2. **Levantar infraestructura**

   ```bash
   docker-compose up -d --build
   ```



## Colección de Postman

Puedes encontrar la colección de Postman para esta API en el siguiente enlace:

[Descargar Colección de Postman](./docs/QuickFix-back.postman_collection.json)

Para usar la colección de Postman:
1. Descarga el archivo JSON.
2. Importa la colección en tu aplicación Postman.
3. Explora y prueba los diferentes endpoints de la API.

## Endpoints Disponibles

### Ping
- **GET** `/ping`
   - **Descripción**: Verifica si el servidor está en línea.
   - **Respuesta Exitosa**: Código 200 OK

### Órdenes
- **GET** `/orders/:order_id`
   - **Descripción**: Obtiene una orden por su ID.

- **GET** `/orders/user/:user_id`
   - **Descripción**: Obtiene órdenes asociadas a un usuario específico.

- **GET** `/orders/professional/:professional_id`
   - **Descripción**: Obtiene órdenes asociadas a un profesional específico.

- **POST** `/orders/`
   - **Descripción**: Crea una nueva orden.
   - **Body**: Datos de la orden a crear.

- **PUT** `/orders/:order_id/accept`
   - **Descripción**: Marca una orden como aceptada por un profesional.

- **PUT** `/orders/:order_id/complete`
   - **Descripción**: Marca una orden como completada.

- **PUT** `/orders/:order_id/cancel`
   - **Descripción**: Cancela una orden.

### Registro de Usuarios y Profesionales
- **POST** `/user/`
   - **Descripción**: Crea un nuevo usuario.
   - **Body**: Datos del usuario a crear.

- **GET** `/user/:user_id`
   - **Descripción**: Obtiene un usuario por su ID.

- **POST** `/professional/`
   - **Descripción**: Crea un nuevo profesional.
   - **Body**: Datos del profesional a crear.

- **GET** `/professional/:professional_id`
   - **Descripción**: Obtiene un profesional por su ID.

### Profesionales por Ubicación
- **GET** `/professionals/:department/:district`
   - **Descripción**: Obtiene profesionales por departamento y distrito en Mendoza.

### Ubicaciones
- **GET** `/locations`
   - **Descripción**: Obtiene todas las ubicaciones disponibles en Mendoza.