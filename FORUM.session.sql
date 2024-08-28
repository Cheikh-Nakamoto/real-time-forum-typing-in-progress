
-- delete from users where username="JDJDSDSDK";
-- SELECT * FROM posts ORDER BY created_at DESC LIMIT 3, 6;
-- SELECT row_num
-- FROM (
--   SELECT
--     username,
--     ROW_NUMBER() OVER (ORDER BY username) AS row_num
--   FROM users
-- ) 
-- WHERE username = 'coulou789';

-- SELECT * FROM posts JOIN cats_posts ON pst_id = post_id JOIN categories ON cat_id = category_id WHERE category_id=1;

-- delete from sessions where sess_id = "3475e583-4846-4bf2-99bb-0fbe8c035d97"

-- select * from posts where title = "Test Post";
-- select * from users where email = "beck@mail.com";

-- SELECT count(*) from posts;
-- SELECT count(*) from posts;

-- delete from posts where updated_at is NULL or created_at is NULL or status is NULL;

-- UPDATE categories SET color="#b8dbd9" WHERE category_id=2;

-- Insert INTO categories (name,color) values ("Music","#ffd166");
-- Insert INTO categories (name,color) values ("Technology","#b8dbd9");
-- Insert INTO categories (name,color) values ("Science","#a7c957");
-- Insert INTO categories (name,color) values ("Health","#92B4A7");
-- Insert INTO categories (name,color) values ("Travel","#ffb3c1");
-- Insert INTO categories (name,color) values ("Fashion","#f28482");
-- Insert INTO categories (name,color) values ("Food","#E7F59E");
-- Insert INTO categories (name,color) values ("Education","#83c5be");
-- Insert INTO categories (name,color) values ("Entertainment","#8c8a93");

-- SELECT posts.* FROM posts LEFT JOIN comments ON pst_id = post_id WHERE usr_id="b4acf1bd-c24a-41d4-97e2-b3fbea3bdb23"  ;

-- UPDATE reactions SET updated_at=`2023-10-05 12:03:52`, usr_id=`7ffccf9c-a320-482b-be40-2f191a54b4e2`, react_id=`bd3ec121-0093-4e2a-b701-e83a1a491bd6`, reactions=`DISLIKE`, pst_id=`54bf250d-259b-46ae-bb57-fd5d6d2fb0c7`, comment_id=`00000000-0000-0000-0000-000000000000`, react_type=`POST` WHERE react_id="bd3ec121-0093-4e2a-b701-e83a1a491bd6" ;

-- DROP TABLE chats;
-- DROP TABLE messages;

-- SELECT * FROM chats WHERE (requester_id="coulou789"  OR recipient_id="coulou789")   AND (recipient_id="jasonwatkins"  OR requester_id="jasonwatkins" )  ;

-- SELECT COUNT(*) FROM messages WHERE (cht_id="766240ae-32a7-4857-86d9-60e0b900af98" ) AND (sender_id!="walllucas" ) AND (read=false ) ;

-- SELECT * FROM posts ORDER BY created_at DESC LIMIT 0, 3 

