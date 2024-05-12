/*
  Warnings:

  - You are about to alter the column `blnc_amt` on the `AccountDetail` table. The data in that column could be lost. The data in that column will be cast from `Decimal(20,2)` to `DoublePrecision`.
  - You are about to alter the column `loan_amt` on the `AccountDetail` table. The data in that column could be lost. The data in that column will be cast from `Decimal(20,2)` to `DoublePrecision`.
  - You are about to alter the column `min_loan_pymnt` on the `AccountDetail` table. The data in that column could be lost. The data in that column will be cast from `Decimal(20,2)` to `DoublePrecision`.
  - You are about to alter the column `amt` on the `TransactionDetail` table. The data in that column could be lost. The data in that column will be cast from `Decimal(20,2)` to `DoublePrecision`.
  - You are about to alter the column `password` on the `User` table. The data in that column could be lost. The data in that column will be cast from `VarChar(200)` to `VarChar(20)`.

*/
-- AlterTable
ALTER TABLE "AccountDetail" ALTER COLUMN "blnc_amt" SET DATA TYPE DOUBLE PRECISION,
ALTER COLUMN "loan_amt" SET DATA TYPE DOUBLE PRECISION,
ALTER COLUMN "min_loan_pymnt" SET DATA TYPE DOUBLE PRECISION;

-- AlterTable
ALTER TABLE "TransactionDetail" ALTER COLUMN "amt" SET DATA TYPE DOUBLE PRECISION;

-- AlterTable
ALTER TABLE "User" ALTER COLUMN "password" SET DATA TYPE VARCHAR(20);
