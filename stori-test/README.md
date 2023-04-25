# Stori-test

Test para Stori, se divide en dos funciones lambda, conectadas por un queue en sqs

```bash
.
├── README.md                   <-- Instucciones
├── process-file                <-- Funcion que procesa un archivo en S3 para calcular los valores requeridos
├── send-email                  <-- Funcion que genera un email y lo procesa por SES
└── template.yaml               <-- Plantilla Cloudformation
```

## Requisitos

* AWS CLI already configured with Administrator permission
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Setup process

### Compilando funciones

```shell
sam build
```
### Deployando cambios
```bash
sam deploy --guided
```

Seguir las instrucciones y cambiar el nombre del email que se usara para enviar y recibir:


## IMPORTANTE

Una vez que se deploye la pila de cloudformation se enviara un email para comprobar la direccion de email, es importante verificar este punto.
