# Información General sobre el Proyecto

Por ahora todo puede funcionar gratis, menos el tema de los números de Twilio, si queremos un numero propio hay que comprarlo, para probar podemos usar el del sandbox.

## Twilio

- Ahora estamos usando la capa gratuita de Twilio, con su sandbox el cual nos provee de un numero de WhatsApp gratis pero tiene limitaciones.

- Para poder adquirir números de teléfono Chilenos hay que registrarse como empresa en Twilio, osea necesitamos un RUT de empresa.

## OPENAI API

- Estamos usando un token de Github para poder acceder a la API de OpenAI, el token tiene limitaciones, 150 request al día.

- Por cada mensaje del usuario usamos una request para generar el embedding del mensaje, y otra request para obtener la respuesta del BOT.

## BOT

- Podemos registrar BOTs para cada restaurante, incluso pueden tener mas de uno. Se puede modificar el Nombre del bot y su "identidad" o personalidad.

- Cada bot debe tener un numero de WhatsApp asociado, este numero es el que se usara para enviar y recibir mensajes. En base a su numero, único para cada bot, podemos obtener su nombre y su identidad.

- Podemos agregarle conocimientos al bot, para este caso que estamos trabajando, de restaurantes y eso, agregamos items de menú, el bot puede acceder a esta información según su restaurante asociado.

- El BOT puede acceder a un historial de chats para poder tener contexto completo de la conversación. La query que hace eso trae las ultimas 5 interacciones, 10 mensajes, mensaje de usuario - mensaje de bot, del dia actual.

- La idea con el tema de los BOTs es que se puedan registrar y configurar de manera sencilla, y que puedan ser utilizados por los restaurantes u organizaciones para atención al cliente.

### FUNCTION CALLS

- La API de OpenAI permite que el bot pueda escoger cuando responder con un mensaje para el usuario, o cuando responder con una estructura de datos, esa estructura puede ser usada para desencadenar alguna acción en el sistema, como reservar una mesa, o pedir un delivery, esto todavía no esta implementado.

## CHAT_HISTORY

- El historial de chats almacena el numero de quien envía un mensaje, osea el cliente o usuario, el numero del bot a quien se le envía el mensaje, el mensaje del usuario, la respuesta del bot.

- Se almacena la fecha y hora de creación, actualización y eliminación del registro, esto es por el ORM que estamos usando.

## RESTAURANT

- El restaurant por ahora es simple, solo almacenamos el nombre y su usuario asociado.

- Tengo la idea de que podemos expandir mas esto, no solo a restaurantes, sino a organizaciones de todo tipo, los BOTS de atención al publico pueden ser útiles en cualquier rubro.

- la relación entre usuario y restaurant es de uno a N, un usuario puede tener muchos restaurantes/organizaciones.

## MENU

- Esta tabla es la base de conocimiento del BOT, ahora son items de menú pero podemos expandir esto a cualquier tipo de información que el BOT pueda necesitar para responder a las consultas de los usuarios.

- La relación entre restaurant y menu es de uno a N, un restaurant puede tener muchos items de menú.

## USER

- El usuario es el que tiene acceso a la plataforma, por ahora manejamos dos roles, admin y user.

## BASE DE DATOS

- Por ahora estamos usando Supabase, porque es gratis y su dashboard es muy amigable, pero la capa gratuita tiene limitaciones, aunque para empezar son mas que suficientes.

- Supabase es un Wraper de alguna Cloud, por lo que es mas caro que usar directamente la Cloud, el plan normal son 20 Dolares al mes, pero con la capa gratuita podemos empezar a trabajar.

## APP MOVIL

- Tengo pensado hacer una app Movil para que el usuario pueda configurar su BOT como quiera, agregar y modificar la información de su restaurante, agregar items de menú, ver el historial de chats, etc.

- La idea es que el usuario pueda hacer todo esto desde su celular, y que el BOT pueda ser utilizado desde cualquier plataforma, WhatsApp, Telegram, Facebook Messenger, etc.

## FRONTEND

- Se puede hacer una plataforma web para que el usuario pueda hacer lo mismo que en la app Movil o expandir las funcionalidades, hay que desarrollar mas esta idea.

## EL NEGOCIO

- Yo creo que el negocio esta en vender el servicio de atención al cliente, con BOTS, a empresas, restaurantes, organizaciones, etc.

- Podemos cobrar por cada BOT nuevo que quiera registrar la organización de hecho hay que hacerlo porque los números de Twilio no son gratis.

## GASTOS ESTIMADOS

La estimación la hice con Claude, arique habría que revisarla, pero para empezar sirve.

Basándome en el documento proporcionado, realizaré una estimación de gastos para los servicios mencionados:

1. Twilio
- Actualmente usando la capa gratuita (Sandbox)
- Para números chilenos: Requiere registro de empresa
- Costo estimado para números de WhatsApp: 
  * Número local: ~$1-2 USD/mes
  * Mensajes: ~$0.005-0.01 USD por mensaje

2. OpenAI API
- Limitación actual: 150 requests/día con token de GitHub
- Costo estimado:
  * Plan básico: $0.0020 por 1000 tokens de entrada
  * Plan básico: $0.0060 por 1000 tokens de salida
  * Asumiendo 10 interacciones diarias: ~$5-10 USD/mes

3. Supabase (Base de Datos)
- Capa gratuita actualmente
- Plan normal: $20 USD/mes
- Costos adicionales por almacenamiento y transferencia

4. Infraestructura Adicional
- Hosting de aplicación: $5-10 USD/mes (servicios como Heroku, DigitalOcean)

Estimación Total de Gastos Mensuales:
- Mínimo (actual): ~$0-5 USD/mes
- Proyección inicial: ~$30-50 USD/mes
- Proyección con crecimiento: $50-100 USD/mes

Consideraciones adicionales:
- Los costos aumentarán proporcionalmente al número de bots y mensajes
- Recomendable tener un fondo de contingencia para escalamiento
- Considerar cobrar por bot para cubrir estos gastos operativos

Basándome en el análisis anterior, el total estimado de gastos mensuales sería:

**Total de Gastos Mensuales: $50 USD**

Este monto incluye:
- Twilio
- OpenAI API
- Supabase
- Infraestructura básica

La estimación considera un escenario inicial de crecimiento moderado, con flexibilidad para ajustes según la demanda real de los servicios.