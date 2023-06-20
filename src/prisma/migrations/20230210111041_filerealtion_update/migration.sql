/*
  Warnings:

  - A unique constraint covering the columns `[table,fileId,order,value]` on the table `FileRelation` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "FileRelation_table_fileId_order_value_key" ON "FileRelation"("table", "fileId", "order", "value");
