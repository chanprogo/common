/*
 Navicat PostgreSQL Data Transfer

 Source Server         : 192.168.56.101-postgres
 Source Server Type    : PostgreSQL
 Source Server Version : 100018
 Source Host           : 192.168.56.101:5432
 Source Catalog        : sammy
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100018
 File Encoding         : 65001

 Date: 12/09/2021 10:17:43
*/


-- ----------------------------
-- Sequence structure for playground_equip_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."playground_equip_id_seq";
CREATE SEQUENCE "public"."playground_equip_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for playground
-- ----------------------------
DROP TABLE IF EXISTS "public"."playground";
CREATE TABLE "public"."playground" (
  "equip_id" int4 NOT NULL DEFAULT nextval('playground_equip_id_seq'::regclass),
  "type" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "color" varchar(25) COLLATE "pg_catalog"."default" NOT NULL,
  "location" varchar(25) COLLATE "pg_catalog"."default",
  "install_date" date
)
;

-- ----------------------------
-- Records of playground
-- ----------------------------

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."playground_equip_id_seq"
OWNED BY "public"."playground"."equip_id";
SELECT setval('"public"."playground_equip_id_seq"', 2, false);

-- ----------------------------
-- Checks structure for table playground
-- ----------------------------
ALTER TABLE "public"."playground" ADD CONSTRAINT "playground_location_check" CHECK (((location)::text = ANY ((ARRAY['north'::character varying, 'south'::character varying, 'west'::character varying, 'east'::character varying, 'northeast'::character varying, 'southeast'::character varying, 'southwest'::character varying, 'northwest'::character varying])::text[])));

-- ----------------------------
-- Primary Key structure for table playground
-- ----------------------------
ALTER TABLE "public"."playground" ADD CONSTRAINT "playground_pkey" PRIMARY KEY ("equip_id");
