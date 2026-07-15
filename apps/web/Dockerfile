FROM node:24.18.0-alpine

WORKDIR /apps

RUN corepack enable

COPY package.json pnpm-lock.yaml ./
RUN pnpm install

COPY . .

EXPOSE 3000

CMD ["pnpm", "run", "dev"]