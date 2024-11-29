# TODOS

## Authorizer

- [ ] Implementar el Authorizer para la API Gateway

## Swagger

- [ ] Implementar el Swagger para los servicios

## Tests

- [ ] Implementar los tests para los servicios

## Bot Service

- [ ] Agregar el nombre del restaurant a la configuraron del System Prompt
- [ ] Mejorar el System Prompt base del bot, para que la personalidad del bot sea mas natural
- [ ] Agregar Function calls para los bots, para que puedan registrar reservaciones y pedidos.
- [ ] Implementar la lógica para que si es el primer mensaje del día de un usuario, el bot envié automáticamente la carta del restaurant mas un mensaje de bienvenida

## Restaurant Service

- [ ] Agregar mas información al modelo de restaurantes, como dirección, teléfono, etc.
- [ ] Evaluar si convertir esta entidad en algo mas genérico, como "Organización" o "Negocio", para ampliar el alcance de la aplicación

## Auth

- [ ] Implementar los middlewares de autenticación y autorización
- [ ] Implementar la verificación del token
- [ ] Modificar la generación del token para que incluya "Bearer"
- [ ] Desacoplar la lógica de autenticación y autorización, de el servicio de usuarios a Shared

## Notifications

- [ ] Implementar el servicio de notificaciones para notificar al usuario cuando se ha registrado una reservación o pedido

## General

- [ ] Mejorar el método PUT de los servicios para que no sea necesario enviar todos los campos

## Documentación

- [ ] Agregar la documentación de los servicios, Diagramas de clases y Diagramas de secuencia

## Twilio

- [ ] Registra la cuenta de Twilio como una de pago, para poder obtener un número de teléfono personalizados para los bots

## AWS

- [ ] Implementar SNS para las notificaciones

## Carga de imágenes

- [ ] Implementar la carga de imágenes para los restaurantes, para que puedan subir su menú y el sistema se encargue de extraer los items, y guardar el documento en S3 