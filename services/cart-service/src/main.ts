import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(AppModule, {
    transport: Transport.NATS,
    options: {
      servers: ['http://localhost:4222/'], 
    },
  });

  await app.listen();
  console.log('📡 NestJS microservice listening to NATS....');
}
bootstrap();


