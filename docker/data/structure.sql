CREATE TABLE p_character (
    id BIGINT NOT NULL AUTO_INCREMENT, 
    name VARCHAR(50),
    class VARCHAR (255),
    discord_user_id VARCHAR(255), 
    PRIMARY KEY (id)
);

-- @toDo Make discord_user_id as primary key if not possible to have several PCs
