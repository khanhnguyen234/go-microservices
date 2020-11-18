package trigger

import (
	"github.com/khanhnguyen234/go-microservices/_postgres"
)

func ProductCreate() {
	db := _postgres.GetPostgres()
	db.Exec(`
		CREATE EXTENSION IF NOT EXISTS "unaccent";
		
		CREATE OR REPLACE FUNCTION slugify("value" TEXT)
		RETURNS TEXT AS $$
			-- removes accents (diacritic signs) from a given string --
			WITH "unaccented" AS (
			SELECT unaccent("value") AS "value"
			),

			-- lowercases the string
			"lowercase" AS (
			SELECT lower("value") AS "value"
			FROM "unaccented"
			),

			-- remove single and double quotes
			"removed_quotes" AS (
			SELECT regexp_replace("value", '[''"]+', '', 'gi') AS "value"
			FROM "lowercase"
			),

			-- replaces anything that's not a letter, number, hyphen('-'), or underscore('_') with a hyphen('-')
			"hyphenated" AS (
			SELECT regexp_replace("value", '[^a-z0-9\\-_]+', '-', 'gi') AS "value"
			FROM "removed_quotes"
			),

			-- trims hyphens('-') if they exist on the head or tail of the string
			"trimmed" AS (
			SELECT regexp_replace(regexp_replace("value", '\-+$', ''), '^\-', '') AS "value"
			FROM "hyphenated"
			),

			"suffix" AS (
				SELECT EXTRACT(EPOCH FROM NOW())::int As "value"
			),
			"result" AS (
				SELECT concat((SELECT "value" FROM "trimmed"), '-', (SELECT "value" FROM "suffix")) as "value"
			)
		  	SELECT "value" FROM "result";
		$$ LANGUAGE SQL STRICT IMMUTABLE;
	`)

	db.Exec(`
		CREATE FUNCTION public.set_slug_from_name() RETURNS TRIGGER
			LANGUAGE plpgsql
		AS $$
		BEGIN
		  NEW.slug := slugify(NEW.name);
		  RETURN NEW;
		END
		$$;
	`)

	db.Exec(`
	CREATE TRIGGER "trigger_insert_product" 
	BEFORE INSERT ON "product_models" FOR EACH ROW 
	WHEN (NEW.name IS NOT NULL AND NEW.slug IS NULL)
	EXECUTE PROCEDURE set_slug_from_name();
	`)
}
