import { PrismaClient } from "@prisma/client";
const prisma = new PrismaClient();

// (async () => {
//   const produto = await prisma.produto.count();
//   console.log(produto);
// })();

// (async () => {
//   const produto = await prisma.produto.create({
//     data: {
//       nome: "Nivea Men Black&White",
//       Ean: "8795543",
//     },
//   });
//   console.log("produto criado: ", produto);
// })();

// (async () => {
//   const produto = await prisma.produto.findMany({
//     where: {
//       nome: {
//         startsWith: "R",
//       },
//     },
//   });
//   console.log(produto);
// })();

// (async () => {
//   const produto = await prisma.produto.create({
//     data: {
//       nome: "Colagate",
//       Ean: "3451",
//       loja: {
//         create: {
//           loja: "Filial 01",
//         },
//       },
//     },
//   });
//   console.log(produto);
// })();

// (async () => {
//   const produto = await prisma.produto.findMany({
//     select: {
//       id: true,
//       nome: true,
//       Ean: true,
//       loja: true,
//     },
//   });
//   console.log(produto);
// })();

(async () => {
  const loja = await prisma.loja.findMany({
    where: {
      loja: {
        startsWith: "F",
      },
    },
  });
  console.log(loja);
})();
