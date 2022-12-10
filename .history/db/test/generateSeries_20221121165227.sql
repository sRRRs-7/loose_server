SELECT * FROM generate_series(1, 5);

INSERT INTO
  sales (sold_on, amount)
SELECT
  '2000-01-01'::DATE + (RANDOM() * 31)::INTEGER AS sold_on,
  (RANDOM() * 1000)::INTEGER + 1000 AS amount
FROM
  -- 100 件生成するように変更
  GENERATE_SERIES(1, 100)
ORDER BY
  sold_on
;