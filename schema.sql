CREATE TABLE comments.posts (postid INT NOT NULL AUTO_INCREMENT, title TINYTEXT, post TEXT, timestamp TIMESTAMP, primary key (postid));
insert into posts(title, post) values('Title of best article in the world', 'this is my an awesome article you read in your life. if you do not know then go to hell ok ? thins your welcome');



insert into comment(link, parent_id, commentor_name, comment_text) values('localhost:1323', 1, 'jawad', 'this is my 4 comment ');

UPDATE comment SET commentor_name = 'hamed', column2 = value WHERE id > 1 and id < 4;

ALTER TABLE comment ALTER comment_id SET DEFAULT NOT NULL;

ALTER table posts ADD COLUMN ownerid int unsigned;

Schema::create('comments', function (Blueprint $table) {
       $table->increments('id');
       $table->integer('user_id')->unsigned();
       $table->integer('parent_id')->unsigned();
       $table->text('comment');
       $table->integer('commentable_id')->unsigned();
       $table->string('commentable_type');
       $table->timestamps();
});

