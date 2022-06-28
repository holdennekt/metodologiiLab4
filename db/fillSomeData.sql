\c todo;

insert into tasks (title, details, deadline) values
('finish 4 lab', 'yes, 4th wall was crushed', '2022-06-29'),
('get beraly good at css', 'learn css a bit, even level "not being disgusted while formatting small pet project" will be fine', '2022-08-31');

insert into tasks (title) values ('watch Neon Genesis Evangelion');

insert into tasks (title, details, deadline, completed, completed_at) values
('make some tea', 'black tea with 2 spoons of sugar', '2022-06-28', true, '2022-06-28'),
('turn 19', 'yea, that`s it', '2022-04-24', true, '2022-04-23');

insert into tasks (title, details, deadline, expired) values
('read some books', 'science fiction like aisek azimov is fine, as well as rey breadberry', '2022-05-23', true),
('find a gift to the friend', 'maybe paper model constructor, maybe some arduino stuff', '2022-06-28', true);
