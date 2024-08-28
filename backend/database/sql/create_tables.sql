-- SQLite
PRAGMA foreign_keys = ON;
CREATE TABLE if not exists follows (
  follow_id INTEGER PRIMARY KEY AUTOINCREMENT,
  following_user_id uuid,
  followed_user_id uuid,
  created_at varchar,
  FOREIGN KEY (following_user_id) REFERENCES users (user_id)
);


CREATE TABLE if not exists users (
  user_id uuid PRIMARY KEY NOT NULL,
  firstname varchar(255) NOT NULL,
  lastname varchar(255) NOT NULL,
  username varchar(255) NOT NULL UNIQUE,
  avatar_url varchar(255) DEFAULT "",
  email varchar(255) NOT NULL UNIQUE,
  password varchar(255) NOT NULL,
  status Text check(status in ('ONLINE', 'OFFLINE')) DEFAULT "OFFLINE",
  blocked bool,
  email_confirmed bool DEFAULT false,
  role TEXT check(
    role in ("USER", "MODERATOR", "ADMINISTRATOR", "GUEST")
  ) DEFAULT "USER",
  created_at varchar,
  updated_at varchar
);


CREATE TABLE if not exists comments (
  comment_id uuid PRIMARY KEY,
  usr_id uuid,
  pst_id uuid,
  body text NOT NULL,
  username varchar,
  created_at varchar,
  FOREIGN KEY (usr_id) REFERENCES users (user_id),
  FOREIGN key (pst_id) REFERENCES posts (post_id)
);

CREATE TABLE if not exists reactions (
  react_id uuid PRIMARY KEY,
  pst_id uuid,
  comment_id uuid,
  usr_id uuid,
  reactions Text check(reactions in ('LIKE', 'DISLIKE')),
  react_type Text check(react_type in ('COMMENT', 'POST')),
  created_at varchar,
  updated_at varchar,
   FOREIGN KEY (pst_id) REFERENCES posts (post_id),
  FOREIGN KEY (usr_id) REFERENCES users (user_id)
);


CREATE TABLE if not exists posts (
  post_id uuid PRIMARY KEY,
  title varchar(255) NOT NULL,
  body text COMMENT 'Content of the post',
  username varchar(255),
  user_id uuid,
  status varchar(255),
  created_at varchar,
  updated_at varchar,
  FOREIGN KEY (user_id) REFERENCES users (user_id)
);


CREATE TABLE if not exists sessions (
  sess_id uuid PRIMARY KEY,
  user_id uuid,
  expire_at varchar,
  token varchar(255),
  created_at varchar varchar,
  remote_addr varchar(255),
  FOREIGN KEY (user_id) REFERENCES users (user_id)
);


CREATE TABLE if not exists categories (
  category_id INTEGER PRIMARY KEY,
  name varchar(255) UNIQUE NOT NULL,
  color varchar
);


CREATE TABLE if not exists cats_posts (
  cat_id INTEGER,
  pst_id uuid,
  FOREIGN KEY (pst_id) REFERENCES posts (post_id),
  FOREIGN KEY (cat_id) REFERENCES categories (category_id)
);


CREATE TABLE if not exists messages (
  message_id uuid PRIMARY KEY,
  content varchar,
  cht_id uuid,
  sender_id varchar,
  read boolean,
  created_at varchar,

  FOREIGN KEY (sender_id) REFERENCES users (username),
  FOREIGN KEY (cht_id) REFERENCES chats (chat_id)
  
);

CREATE TABLE if not exists chats (
  chat_id uuid PRIMARY KEY,
  requester_id uuid,
  recipient_id uuid,
  last_message_time varchar DEFAULT "",
  created_at varchar,

  FOREIGN KEY (requester_id) REFERENCES users (username),
  FOREIGN KEY (recipient_id) REFERENCES users (username)
);
