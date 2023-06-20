/*
  Warnings:

  - You are about to drop the column `sellingPriceEstimation` on the `Inventory` table. All the data in the column will be lost.
  - You are about to drop the column `clientFullName` on the `OnlineOrderDetail` table. All the data in the column will be lost.
  - You are about to drop the column `email` on the `OnlineOrderDetail` table. All the data in the column will be lost.
  - You are about to drop the column `phoneNumber` on the `OnlineOrderDetail` table. All the data in the column will be lost.
  - You are about to drop the column `redirectUrl` on the `OnlineOrderDetail` table. All the data in the column will be lost.
  - You are about to drop the `DailySummery` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `EmployeeDailySummery` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `MonthlySummery` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `maxSellingPriceEstimation` to the `Inventory` table without a default value. This is not possible if the table is not empty.
  - Added the required column `minSellingPriceEstimation` to the `Inventory` table without a default value. This is not possible if the table is not empty.
  - Added the required column `redirectUrl` to the `OnlineOrderPayment` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "DailySummery" DROP CONSTRAINT "DailySummery_companyId_fkey";

-- DropForeignKey
ALTER TABLE "DailySummery" DROP CONSTRAINT "DailySummery_userId_fkey";

-- DropForeignKey
ALTER TABLE "EmployeeDailySummery" DROP CONSTRAINT "EmployeeDailySummery_companyId_fkey";

-- DropForeignKey
ALTER TABLE "EmployeeDailySummery" DROP CONSTRAINT "EmployeeDailySummery_managerId_fkey";

-- DropForeignKey
ALTER TABLE "MonthlySummery" DROP CONSTRAINT "MonthlySummery_companyId_fkey";

-- AlterTable
ALTER TABLE "Inventory" DROP COLUMN "sellingPriceEstimation",
ADD COLUMN     "maxSellingPriceEstimation" DOUBLE PRECISION NOT NULL,
ADD COLUMN     "minSellingPriceEstimation" DOUBLE PRECISION NOT NULL;

-- AlterTable
ALTER TABLE "OnlineOrderDetail" DROP COLUMN "clientFullName",
DROP COLUMN "email",
DROP COLUMN "phoneNumber",
DROP COLUMN "redirectUrl",
ADD COLUMN     "customerID" TEXT,
ADD COLUMN     "referralID" TEXT;

-- AlterTable
ALTER TABLE "OnlineOrderPayment" ADD COLUMN     "redirectUrl" TEXT NOT NULL;

-- DropTable
DROP TABLE "DailySummery";

-- DropTable
DROP TABLE "EmployeeDailySummery";

-- DropTable
DROP TABLE "MonthlySummery";

-- CreateTable
CREATE TABLE "Customer" (
    "id" TEXT NOT NULL,
    "fullName" TEXT NOT NULL,
    "phoneNumber" TEXT,
    "phoneNumberVerified" BOOLEAN DEFAULT false,
    "email" TEXT,
    "emailVerified" BOOLEAN NOT NULL DEFAULT false,
    "companyId" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "Customer_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TableCount" (
    "id" TEXT NOT NULL,
    "count" INTEGER NOT NULL DEFAULT 1,

    CONSTRAINT "TableCount_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Promoter" (
    "id" TEXT NOT NULL,
    "fullName" TEXT NOT NULL,
    "phoneNumber" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "companyId" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "Promoter_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "PromoterBankDetail" (
    "id" TEXT NOT NULL,
    "bankName" TEXT NOT NULL,
    "bankAccount" TEXT NOT NULL,
    "promoterId" TEXT NOT NULL,

    CONSTRAINT "PromoterBankDetail_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Referral" (
    "id" TEXT NOT NULL,
    "promoCode" TEXT NOT NULL,
    "discount" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "promoterCut" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "startDate" TIMESTAMP(3) NOT NULL,
    "endDate" TIMESTAMP(3),
    "companyId" TEXT NOT NULL,
    "promoterId" TEXT,
    "allProduct" BOOLEAN NOT NULL DEFAULT true,
    "extra" JSONB,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "Referral_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "ReferralProduct" (
    "referralId" TEXT NOT NULL,
    "productId" TEXT NOT NULL,

    CONSTRAINT "ReferralProduct_pkey" PRIMARY KEY ("referralId","productId")
);

-- CreateTable
CREATE TABLE "EmployeeDailySummary" (
    "id" TEXT NOT NULL,
    "earning" DOUBLE PRECISION NOT NULL,
    "profit" DOUBLE PRECISION NOT NULL,
    "managerAccepted" BOOLEAN NOT NULL DEFAULT false,
    "date" DATE NOT NULL,
    "managerId" TEXT,
    "employeeId" TEXT NOT NULL,
    "companyId" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "EmployeeDailySummary_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Summary" (
    "id" TEXT NOT NULL,
    "earning" DOUBLE PRECISION NOT NULL,
    "profit" DOUBLE PRECISION NOT NULL,
    "startDate" DATE NOT NULL,
    "endDate" DATE NOT NULL,
    "companyId" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "Summary_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Customer_phoneNumber_key" ON "Customer"("phoneNumber");

-- CreateIndex
CREATE UNIQUE INDEX "Promoter_companyId_phoneNumber_key" ON "Promoter"("companyId", "phoneNumber");

-- CreateIndex
CREATE UNIQUE INDEX "PromoterBankDetail_bankName_bankAccount_promoterId_key" ON "PromoterBankDetail"("bankName", "bankAccount", "promoterId");

-- CreateIndex
CREATE UNIQUE INDEX "Referral_promoCode_companyId_key" ON "Referral"("promoCode", "companyId");

-- CreateIndex
CREATE UNIQUE INDEX "EmployeeDailySummary_companyId_date_employeeId_key" ON "EmployeeDailySummary"("companyId", "date", "employeeId");

-- CreateIndex
CREATE UNIQUE INDEX "Summary_companyId_startDate_endDate_key" ON "Summary"("companyId", "startDate", "endDate");

-- AddForeignKey
ALTER TABLE "Customer" ADD CONSTRAINT "Customer_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "Company"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "OnlineOrderDetail" ADD CONSTRAINT "OnlineOrderDetail_customerID_fkey" FOREIGN KEY ("customerID") REFERENCES "Customer"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "OnlineOrderDetail" ADD CONSTRAINT "OnlineOrderDetail_referralID_fkey" FOREIGN KEY ("referralID") REFERENCES "Referral"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Promoter" ADD CONSTRAINT "Promoter_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "Company"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PromoterBankDetail" ADD CONSTRAINT "PromoterBankDetail_promoterId_fkey" FOREIGN KEY ("promoterId") REFERENCES "Promoter"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Referral" ADD CONSTRAINT "Referral_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "Company"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Referral" ADD CONSTRAINT "Referral_promoterId_fkey" FOREIGN KEY ("promoterId") REFERENCES "Promoter"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "ReferralProduct" ADD CONSTRAINT "ReferralProduct_productId_fkey" FOREIGN KEY ("productId") REFERENCES "Product"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "ReferralProduct" ADD CONSTRAINT "ReferralProduct_referralId_fkey" FOREIGN KEY ("referralId") REFERENCES "Referral"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "EmployeeDailySummary" ADD CONSTRAINT "EmployeeDailySummary_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "Company"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "EmployeeDailySummary" ADD CONSTRAINT "EmployeeDailySummary_managerId_fkey" FOREIGN KEY ("managerId") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Summary" ADD CONSTRAINT "Summary_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "Company"("id") ON DELETE CASCADE ON UPDATE CASCADE;
