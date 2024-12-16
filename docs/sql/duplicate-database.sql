CREATE DATABASE "dipl-test"
WITH
  TEMPLATE "diplomski-rad-2024-demo-final" OWNER "postgres";

SELECT
  PG_TERMINATE_BACKEND(pg_stat_activity.pid)
FROM
  pg_stat_activity
WHERE
  pg_stat_activity.datname = "diplomski-rad-2024-demo-final"
  AND pid <> PG_BACKEND_PID();