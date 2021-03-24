CREATE TABLE analysis (
   Id  SERIAL,
   Dna TEXT NOT NULL,
   IsMutant BOOLEAN NOT NULL,
   PRIMARY KEY (Id, IsMutant)
);
