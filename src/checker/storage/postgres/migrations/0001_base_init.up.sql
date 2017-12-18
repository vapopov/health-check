
-- logs storage
CREATE TABLE checker_log (
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  url TEXT NOT NULL,
  successful BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX created_at_url_idx ON checker_log (created_at, url);


-- test view
CREATE VIEW _health_stats AS
  SELECT url, coalesce(active_time/all_activity_time * 100, 0) AS active
  FROM
    (SELECT url,
       sum(activity_seconds) AS all_activity_time,
       sum(CASE WHEN successful THEN activity_seconds END) AS active_time
    FROM
       (SELECT created_at,
          successful,
          url,
          next_timestamp,
          extract('epoch' from next_timestamp - created_at) AS activity_seconds
       FROM
          (SELECT a.created_at,
             a.url,
             a.successful,
             min(b.created_at) AS next_timestamp
          FROM checker_log AS a, checker_log AS b
          WHERE a.created_at < b.created_at AND a.url = b.url
          GROUP BY 1, 2, 3) AS am) AS base
    GROUP BY url) AS percentage_base;
