DROP TABLE IF EXISTS docker.followers;
DROP TABLE IF EXISTS docker.users;

CREATE TABLE docker.users (
   id INT AUTO_INCREMENT PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   nick VARCHAR(50) NOT NULL UNIQUE,
   email VARCHAR(50) NOT NULL UNIQUE,
   password VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
);

CREATE TABLE docker.followers(
	user_id int not null,
	follower_id int not null,
	
	PRIMARY KEY (user_id, follower_id),
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE
);