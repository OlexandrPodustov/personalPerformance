WITH
    t1 AS (
      SELECT
        SETTINGS.KEY,
        SETTINGS.VALUE,
        1
      FROM
        DASHBOARD_USER_SETTINGS AS SETTINGS
      WHERE
        SETTINGS.KEY = 'Widget1'
        AND SETTINGS.USER_ID = 'userid3'
  ),
    t2 AS (SELECT
             SETTINGS.KEY,
             SETTINGS.VALUE,
             2 AS PR
           FROM
             DASHBOARD_SETTINGS AS SETTINGS
           WHERE
             SETTINGS.KEY = 'Widget1'
             AND SETTINGS.USER_ID = 'userid3'
  ),
    bt AS (SELECT *
           FROM t2
           UNION
           SELECT *
           FROM
             t1)


SELECT
  key,
  value,
  priority,
  CASE
  WHEN priority > 0
    THEN field2 / field1
  ELSE 0
  END
    AS f4
FROM bt
