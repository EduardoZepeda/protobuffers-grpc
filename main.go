package main

// Remote Procedure Call (RPC): protocolo que oculta la implementación en el backend de la petición que hizo un cliente, aunque el cliente sepa como hacer la petición y pueda invocarla como si fuese suya.

// gRPC: framework creado por Google para trabajar RPC con más eficiencia y alto rendimiento.

//     El transporte de datos funciona con HTTP2.
//         Permite crear multiplexación a la hora de enviar mensajes: más mensajes en la conexión TCP de manera simultanea.
//         Permite serializar datos.
//     Usa los protobuffers como estructura para intercambio de datos.
//         Permite serializar y deserializar datos más rápido.

// Métodos de gRPC

// Unary: similar a como funciona REST. Envía una petición al servidor, y el servidor la responde.

// Streaming: permite constante envío de data en un canal.

//     Del lado del cliente: el cliente envía muchas peticiones, y el servidor responde una sola vez.
//     Del lado del servidor: el cliente realiza una sola petición, y el servidor responde enviando la data en partes.
//     Bidireccional: cliente y servidor deciden ambos comunicarse por streaming de data.

// La serialización y deserialización de ambos formatos siempre ocurre. Los protobbufers tienen mucha menor latencia que los JSON al hacerlo.

// A travès de la historia hemos tenido los siguientes formatos:

//     EDI
//     XML
//     JSON

// JSON: formato de mensajes eficiente para JavaScript.

//     Pares de llave y valor.

//     Es más fácil de leer al ojo humano.

//     Es costoso en rendimiento si se quiere trabajar con otro lenguaje distinto de JavaScript.

// Protobuffers: formato de mensaje agnóstico a cualquier lenguaje de programación.

//     Un compilador se encarga de convertir la sintaxis de protobuffer al lenguaje correspondiente.
//     Esta compilación solo ocurre en tiempo de creación o modificación, no en tiempo de ejecución.
//     Se puede llamar archivos .proto desde otros archivos .proto.

// ¿Cuándo usar?

//     JSON: cuando la aplicación requiere que la data sea más flexible.
//     Protobuffers: Cuando la aplicación necesita correr muy rápido; cuando los procesos de serialización y deserialización deben ocurrir rápido.
