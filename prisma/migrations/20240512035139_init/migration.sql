/*
  Warnings:

  - You are about to alter the column `blnc_amt` on the `AccountDetail` table. The data in that column could be lost. The data in that column will be cast from `DoublePrecision` to `Decimal(20,2)`.
  - You are about to alter the column `loan_amt` on the `AccountDetail` table. The data in that column could be lost. The data in that column will be cast from `DoublePrecision` to `Decimal(20,2)`.
  - You are about to alter the column `min_loan_pymnt` on the `AccountDetail` table. The data in that column could be lost. The data in that column will be cast from `DoublePrecision` to `Decimal(20,2)`.
  - You are about to alter the column `amt` on the `TransactionDetail` table. The data in that column could be lost. The data in that column will be cast from `DoublePrecision` to `Decimal(20,2)`.
  - A unique constraint covering the columns `[prin_pan]` on the table `AccountDetail` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[email]` on the table `User` will be added. If there are existing duplicate values, this will fail.

*/
-- AlterTable
ALTER TABLE "AccountDetail" ALTER COLUMN "blnc_amt" SET DATA TYPE DECIMAL(20,2),
ALTER COLUMN "loan_amt" SET DATA TYPE DECIMAL(20,2),
ALTER COLUMN "min_loan_pymnt" SET DATA TYPE DECIMAL(20,2);

-- AlterTable
ALTER TABLE "TransactionDetail" ALTER COLUMN "amt" SET DATA TYPE DECIMAL(20,2);

-- CreateIndex
CREATE UNIQUE INDEX "AccountDetail_prin_pan_key" ON "AccountDetail"("prin_pan");

-- CreateIndex
CREATE UNIQUE INDEX "User_email_key" ON "User"("email");
