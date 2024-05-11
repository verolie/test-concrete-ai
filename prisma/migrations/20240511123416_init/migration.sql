/*
  Warnings:

  - Added the required column `receiver_pan` to the `TransactionDetail` table without a default value. This is not possible if the table is not empty.
  - Added the required column `sender_pan` to the `TransactionDetail` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "TransactionDetail" ADD COLUMN     "receiver_pan" VARCHAR(50) NOT NULL,
ADD COLUMN     "sender_pan" VARCHAR(50) NOT NULL;
