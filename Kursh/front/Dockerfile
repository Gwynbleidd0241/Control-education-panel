FROM node:20

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# CRA по умолчанию стартует на 3000
EXPOSE 3000

CMD ["npm", "start"]
