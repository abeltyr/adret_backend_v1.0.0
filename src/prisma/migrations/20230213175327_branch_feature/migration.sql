/*
  Warnings:

  - You are about to drop the column `latitude` on the `Company` table. All the data in the column will be lost.
  - You are about to drop the column `longitude` on the `Company` table. All the data in the column will be lost.
  - You are about to drop the `Customer` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `OnlineOrderDetail` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `OnlineOrderPayment` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Promoter` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `PromoterBankDetail` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Referral` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `ReferralProduct` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Customer" DROP CONSTRAINT "Customer_companyId_fkey";

-- DropForeignKey
ALTER TABLE "OnlineOrderDetail" DROP CONSTRAINT "OnlineOrderDetail_customerID_fkey";

-- DropForeignKey
ALTER TABLE "OnlineOrderDetail" DROP CONSTRAINT "OnlineOrderDetail_id_fkey";

-- DropForeignKey
ALTER TABLE "OnlineOrderDetail" DROP CONSTRAINT "OnlineOrderDetail_referralID_fkey";

-- DropForeignKey
ALTER TABLE "OnlineOrderPayment" DROP CONSTRAINT "OnlineOrderPayment_id_fkey";

-- DropForeignKey
ALTER TABLE "Promoter" DROP CONSTRAINT "Promoter_companyId_fkey";

-- DropForeignKey
ALTER TABLE "PromoterBankDetail" DROP CONSTRAINT "PromoterBankDetail_promoterId_fkey";

-- DropForeignKey
ALTER TABLE "Referral" DROP CONSTRAINT "Referral_companyId_fkey";

-- DropForeignKey
ALTER TABLE "Referral" DROP CONSTRAINT "Referral_promoterId_fkey";

-- DropForeignKey
ALTER TABLE "ReferralProduct" DROP CONSTRAINT "ReferralProduct_productId_fkey";

-- DropForeignKey
ALTER TABLE "ReferralProduct" DROP CONSTRAINT "ReferralProduct_referralId_fkey";

-- AlterTable
ALTER TABLE "Company" DROP COLUMN "latitude",
DROP COLUMN "longitude";

-- AlterTable
ALTER TABLE "CompanyBranch" ADD COLUMN     "latitude" TEXT,
ADD COLUMN     "longitude" TEXT;

-- AlterTable
ALTER TABLE "User" ADD COLUMN     "branchId" TEXT;

-- DropTable
DROP TABLE "Customer";

-- DropTable
DROP TABLE "OnlineOrderDetail";

-- DropTable
DROP TABLE "OnlineOrderPayment";

-- DropTable
DROP TABLE "Promoter";

-- DropTable
DROP TABLE "PromoterBankDetail";

-- DropTable
DROP TABLE "Referral";

-- DropTable
DROP TABLE "ReferralProduct";

-- DropEnum
DROP TYPE "OrderStatus";

-- AddForeignKey
ALTER TABLE "User" ADD CONSTRAINT "User_branchId_fkey" FOREIGN KEY ("branchId") REFERENCES "CompanyBranch"("id") ON DELETE SET NULL ON UPDATE CASCADE;
