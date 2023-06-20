/*
  Warnings:

  - You are about to drop the column `deposit` on the `DailySummery` table. All the data in the column will be lost.
  - You are about to drop the column `deposit` on the `EmployeeDailySummery` table. All the data in the column will be lost.
  - You are about to drop the column `deposit` on the `MonthlySummery` table. All the data in the column will be lost.
  - The primary key for the `Sales` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `createdAt` on the `Sales` table. All the data in the column will be lost.
  - You are about to drop the column `date` on the `Sales` table. All the data in the column will be lost.
  - You are about to drop the column `deletedAt` on the `Sales` table. All the data in the column will be lost.
  - You are about to drop the column `id` on the `Sales` table. All the data in the column will be lost.
  - You are about to drop the column `initialPrice` on the `Sales` table. All the data in the column will be lost.
  - You are about to drop the column `note` on the `Sales` table. All the data in the column will be lost.
  - You are about to drop the column `sellerId` on the `Sales` table. All the data in the column will be lost.
  - You are about to drop the column `updatedAt` on the `Sales` table. All the data in the column will be lost.
  - Added the required column `amount` to the `Sales` table without a default value. This is not possible if the table is not empty.
  - Added the required column `orderId` to the `Sales` table without a default value. This is not possible if the table is not empty.

*/
-- CreateEnum
CREATE TYPE "OrderStatus" AS ENUM ('Pending', 'PaidPending', 'Delivering', 'Delivered');

-- DropForeignKey
ALTER TABLE "Sales" DROP CONSTRAINT "Sales_sellerId_fkey";

-- AlterTable
ALTER TABLE "DailySummery" DROP COLUMN "deposit";

-- AlterTable
ALTER TABLE "EmployeeDailySummery" DROP COLUMN "deposit";

-- AlterTable
ALTER TABLE "MonthlySummery" DROP COLUMN "deposit";

-- AlterTable
ALTER TABLE "Sales" DROP CONSTRAINT "Sales_pkey",
DROP COLUMN "createdAt",
DROP COLUMN "date",
DROP COLUMN "deletedAt",
DROP COLUMN "id",
DROP COLUMN "initialPrice",
DROP COLUMN "note",
DROP COLUMN "sellerId",
DROP COLUMN "updatedAt",
ADD COLUMN     "amount" INTEGER NOT NULL,
ADD COLUMN     "orderId" TEXT NOT NULL,
ADD CONSTRAINT "Sales_pkey" PRIMARY KEY ("orderId", "inventoryId");

-- CreateTable
CREATE TABLE "Order" (
    "id" TEXT NOT NULL,
    "online" BOOLEAN NOT NULL,
    "orderNumber" TEXT NOT NULL,
    "note" TEXT NOT NULL,
    "totalPrice" DOUBLE PRECISION NOT NULL,
    "totalProfit" DOUBLE PRECISION NOT NULL,
    "paid" BOOLEAN NOT NULL DEFAULT false,
    "date" DATE NOT NULL,
    "sellerId" TEXT,
    "companyId" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "Order_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "OnlineOrderDetail" (
    "id" TEXT NOT NULL,
    "delivery" BOOLEAN NOT NULL,
    "latitude" TEXT,
    "longitude" TEXT,
    "locationName" TEXT,
    "detail" TEXT NOT NULL,
    "phoneNumber" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "clientFullName" TEXT NOT NULL,
    "redirectUrl" TEXT NOT NULL,
    "status" "OrderStatus" NOT NULL,
    "extra" JSONB,

    CONSTRAINT "OnlineOrderDetail_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "OnlineOrderPayment" (
    "id" TEXT NOT NULL,
    "method" TEXT NOT NULL,
    "txRef" VARCHAR(255) NOT NULL,
    "reference" VARCHAR(255) NOT NULL,
    "paid" BOOLEAN NOT NULL,
    "paidDate" TIMESTAMP(3) NOT NULL,
    "extra" JSONB,

    CONSTRAINT "OnlineOrderPayment_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Order_orderNumber_key" ON "Order"("orderNumber");

-- CreateIndex
CREATE UNIQUE INDEX "OnlineOrderPayment_txRef_key" ON "OnlineOrderPayment"("txRef");

-- CreateIndex
CREATE UNIQUE INDEX "OnlineOrderPayment_reference_key" ON "OnlineOrderPayment"("reference");

-- AddForeignKey
ALTER TABLE "Order" ADD CONSTRAINT "Order_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "Company"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Order" ADD CONSTRAINT "Order_sellerId_fkey" FOREIGN KEY ("sellerId") REFERENCES "User"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "OnlineOrderDetail" ADD CONSTRAINT "OnlineOrderDetail_id_fkey" FOREIGN KEY ("id") REFERENCES "Order"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "OnlineOrderPayment" ADD CONSTRAINT "OnlineOrderPayment_id_fkey" FOREIGN KEY ("id") REFERENCES "Order"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Sales" ADD CONSTRAINT "Sales_orderId_fkey" FOREIGN KEY ("orderId") REFERENCES "Order"("id") ON DELETE CASCADE ON UPDATE CASCADE;
