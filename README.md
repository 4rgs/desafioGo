Backend API on GO-lang


# Go Backend Palindrom Challenge

Servicio API REST para el desafio de productos palindromes desarrollado en GO

## API Reference

#### Find Product

```http
  GET /api/productos/busqueda
```

| Parameter     | Type     | Description                       |
| :--------     | :------- | :-------------------------------- |
| `query`       | `String` | String of iem to search           |
| `page`        | `Number` | page of the document              |


## Deployment

To deploy this proyect first we need to deploy this repo and have Docker installed:
```bash
git clone https://github.com/walmartdigital/products-db
```
```bash
cd products-db
```
```bash
sudo make database-docker-up
```
```bash
sudo make database-provision

```
```bash
cd ..
```
Then we can deploy our service that will run on previous DATABASE

```bash
git clone https://github.com/4rgs/desafioGo.git
```
```bash
cd desafio-backend
```
```bash
npm install
```
```bash
npm run start
```


## Authors

- [@AlvaroGonzalez](https://github.com/4rgs)


## Demo on Lab

https://api.spids.cl