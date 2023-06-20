/*
  Warnings:

  - A unique constraint covering the columns `[table,order,value]` on the table `FileRelation` will be added. If there are existing duplicate values, this will fail.

*/
-- DropIndex
DROP INDEX "FileRelation_table_fileId_order_value_key";

-- AlterTable
ALTER TABLE "Inventory" ADD COLUMN     "order" INTEGER NOT NULL DEFAULT 0;

-- AlterTable
ALTER TABLE "Product" ADD COLUMN     "branchId" TEXT;

-- CreateTable
CREATE TABLE "CompanyBranch" (
    "id" TEXT NOT NULL,
    "branchName" TEXT NOT NULL,
    "companyId" TEXT NOT NULL,

    CONSTRAINT "CompanyBranch_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "CompanyBranch_branchName_companyId_key" ON "CompanyBranch"("branchName", "companyId");

-- CreateIndex
CREATE UNIQUE INDEX "FileRelation_table_order_value_key" ON "FileRelation"("table", "order", "value");

-- AddForeignKey
ALTER TABLE "CompanyBranch" ADD CONSTRAINT "CompanyBranch_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "Company"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Product" ADD CONSTRAINT "Product_branchId_fkey" FOREIGN KEY ("branchId") REFERENCES "CompanyBranch"("id") ON DELETE SET NULL ON UPDATE CASCADE;
