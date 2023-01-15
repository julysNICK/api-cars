
CREATE TABLE "cars"(
    "id" serial NOT NULL,
    "model" varchar(255) NOT NULL,
    "make" varchar(255) NOT NULL,
    "image" varchar(255) NOT NULL,
    "reason_to_sell" varchar(255) NOT NULL,
    "description_car" varchar(255) NOT NULL,
    "year" varchar(255) NOT NULL,
    "color" varchar(255) NOT NULL,
    "price" int NOT NULL,
    "is_sold" boolean NOT NULL,
    "user_id" int NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
)