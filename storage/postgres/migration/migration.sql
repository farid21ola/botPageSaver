CREATE TABLE "Page"
(
    "id"       serial                   NOT NULL,
    "URL"      character varying UNIQUE NOT NULL,
    "UserName" character varying        NOT NULL,
    CONSTRAINT "Page_pk" PRIMARY KEY ("id")
) WITH (
      OIDS= FALSE
    );