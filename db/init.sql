\c todo;

drop table if exists tasks;

create table tasks (
  id bigserial not null primary key,
  title varchar(50) not null,
  details varchar(200),
  deadline date,
  expired boolean not null default false,
  completed boolean not null default false,
  completed_at date
);
