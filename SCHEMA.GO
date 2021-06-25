insert into comment(link, parent_id, commentor_name, comment_text) values('localhost:1323', 1, 'jawad', 'this is my 4 comment ');

UPDATE comment SET commentor_name = 'hamed', column2 = value WHERE id > 1 and id < 4;

ALTER TABLE comment ALTER comment_id SET DEFAULT NOT NULL;
