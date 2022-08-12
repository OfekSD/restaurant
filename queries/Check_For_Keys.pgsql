CREATE OR REPLACE FUNCTION check_foreign_key_array(data anyarray, ref_schema text, ref_table text, ref_column text)
    RETURNS BOOL
    RETURNS NULL ON NULL INPUT
    LANGUAGE plpgsql
AS
$body$
DECLARE
    fake_id text;
    sql text default format($$
            select id::text
            from unnest($1) as x(id)
            where id is not null
              and id not in (select %3$I
                             from %1$I.%2$I
                             where %3$I = any($1))
            limit 1;
        $$, ref_schema, ref_table, ref_column);
BEGIN
    EXECUTE sql USING data INTO fake_id;

    IF (fake_id IS NOT NULL) THEN
        RAISE NOTICE 'Array element value % does not exist in column %.%.%', fake_id, ref_schema, ref_table, ref_column;
        RETURN false;
    END IF;

    RETURN true;
END
$body$;

