-- AlterTable
ALTER TABLE "PriceHistory" ALTER COLUMN "initialPrice" DROP NOT NULL,
ALTER COLUMN "minSellingPriceEstimation" DROP NOT NULL,
ALTER COLUMN "maxSellingPriceEstimation" DROP NOT NULL;
