CREATE OR REPLACE FUNCTION search_menu(
    query_embedding vector(3072),
    similarity_threshold float,
    match_count int,
    restaurant_id bigint
)
RETURNS TABLE (
    id bigint,
    item_name text,
    price int,
    description text,
    likes int,
    similarity float
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT
        menus.id,
        menus.item_name,
        menus.price,
        menus.description,
        menus.likes,
        menus.embedding <#> query_embedding AS similarity
    FROM
        menus
    WHERE
        menus.restaurant_id = restaurant_id
        AND menus.embedding <#> query_embedding < similarity_threshold
    ORDER BY
        similarity
    LIMIT
        match_count;
END;
$$;