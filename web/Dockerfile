FROM node:21-alpine as builder

WORKDIR /app

COPY package.json package-lock.json* ./

RUN npm install

COPY . .

RUN npm run build

# Stage 2: Running the code
FROM node:21-alpine

WORKDIR /app

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json

EXPOSE 3000

CMD ["npm", "start"]
