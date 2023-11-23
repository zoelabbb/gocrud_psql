CREATE TABLE buku
(
 id serial NOT NULL,
 judul_buku character varying NOT NULL,
 penulis character varying,
 tgl_publikasi date,
 CONSTRAINT pk_buku PRIMARY KEY (id )
)
WITH (
 OIDS=FALSE
);
ALTER TABLE buku
 OWNER TO postgres;